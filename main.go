package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/AkibaSummer/Danmu/sdk/notify"
	"github.com/AkibaSummer/Danmu/sdk/spider"
	"github.com/spf13/viper"
)

type healthState struct {
	mu             sync.RWMutex
	Status         string    `json:"status"`
	FailureKind    string    `json:"failure_kind,omitempty"`
	Detail         string    `json:"detail,omitempty"`
	AuthStatus     string    `json:"auth_status"`
	Connected      bool      `json:"connected"`
	LastConnected  time.Time `json:"last_connected,omitempty"`
	LastTransition time.Time `json:"last_transition"`
	RestartCount   int       `json:"restart_count"`
}

func (h *healthState) update(fn func(*healthState)) {
	h.mu.Lock()
	defer h.mu.Unlock()
	fn(h)
	h.LastTransition = time.Now()
}

func (h *healthState) serveHTTP(w http.ResponseWriter, _ *http.Request) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	if !h.Connected {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	_ = json.NewEncoder(w).Encode(h)
}

func main() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetDefault("bili.QRLogin", true)
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("read conf/conf.yaml: ", err)
	}

	shortID := viper.GetInt("bili.ShortID")
	if shortID <= 0 {
		log.Fatal("bili.ShortID must be greater than zero")
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	expectedUID := viper.GetInt64("bili.UID")
	instance := valueOr(os.Getenv("DANMU_INSTANCE"), "Danmu")
	var onQRCode func(context.Context, []byte) error
	larkAppID := viper.GetString("lark.AppID")
	larkAppSecret := viper.GetString("lark.AppSecret")
	larkChatID := viper.GetString("lark.ChatID")
	if larkAppID != "" && larkAppSecret != "" && larkChatID != "" {
		larkNotifier, err := notify.NewLarkNotifier(notify.LarkConfig{
			AppID: larkAppID, AppSecret: larkAppSecret, ChatID: larkChatID,
		})
		if err != nil {
			log.Printf("LARK_NOTIFY_DISABLED: %v", err)
		} else {
			onQRCode = func(ctx context.Context, png []byte) error {
				if err := larkNotifier.NotifyQRCode(ctx, instance, expectedUID, png); err != nil {
					log.Printf("LARK_QR_NOTIFY_FAILED instance=%s uid=%d: %v", instance, expectedUID, err)
					return err
				}
				log.Printf("LARK_QR_NOTIFY_SENT instance=%s uid=%d", instance, expectedUID)
				return nil
			}
		}
	} else if larkAppID != "" || larkAppSecret != "" || larkChatID != "" {
		log.Printf("LARK_NOTIFY_DISABLED: incomplete lark configuration")
	}

	cookie := buildCookie(viper.GetString("bili.Cookie"), viper.GetString("bili.SESSDATA"), viper.GetString("bili.BUVID"))
	auth := spider.NewAuthManager(spider.AuthConfig{
		Cookie:       cookie,
		RefreshToken: viper.GetString("bili.RefreshToken"),
		StateFile:    valueOr(viper.GetString("bili.AuthStateFile"), "conf/auth_state.json"),
		ExpectedUID:  expectedUID,
		OnQRCode:     onQRCode,
	})
	if err := auth.Load(); err != nil {
		log.Printf("AUTH_STATE_LOAD_FAILED: %v", err)
	}

	state := &healthState{Status: "starting", AuthStatus: "unknown", LastTransition: time.Now()}
	healthAddr := valueOr(viper.GetString("health.addr"), "127.0.0.1:10216")
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", state.serveHTTP)
	mux.HandleFunc("/login/qr", auth.QRPage)
	mux.HandleFunc("/login/qr.png", auth.QRPNG)
	server := &http.Server{Addr: healthAddr, Handler: mux, ReadHeaderTimeout: 3 * time.Second}
	go func() {
		log.Printf("health endpoint listening on http://%s", healthAddr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("HEALTH_SERVER_FAILED: %v", err)
		}
	}()
	defer server.Shutdown(context.Background())

	spider.Init()
	backoff := time.Second
	authCheckInterval := viper.GetDuration("bili.AuthCheckInterval")
	if authCheckInterval <= 0 {
		authCheckInterval = 15 * time.Minute
	}
	authRetryInterval := viper.GetDuration("bili.AuthRetryInterval")
	if authRetryInterval <= 0 {
		authRetryInterval = time.Minute
	}
	for ctx.Err() == nil {
		authStatus, err := auth.EnsureValid(ctx)
		if err != nil && spider.KindOf(err) != spider.FailureAuth {
			kind := spider.KindOf(err)
			state.update(func(h *healthState) {
				h.Status, h.Connected, h.AuthStatus = "stopped", false, "check_failed"
				h.FailureKind, h.Detail = string(kind), err.Error()
			})
			log.Printf("AUTH_CHECK_FAILED kind=%s; recording is stopped (fail-closed): %v", kind, err)
			if !wait(ctx, authRetryInterval) {
				break
			}
			continue
		}
		if shouldStartQRCode(authStatus, err) {
			detail := "login expired; recording stopped; scan QR to resume"
			authLabel := authStatus.Label()
			if err != nil {
				detail = err.Error()
				authLabel = "refresh_failed"
				log.Printf("AUTH_CHECK_FAILED kind=%s; switching to QR login: %v", spider.KindOf(err), err)
			}
			state.update(func(h *healthState) {
				h.Status, h.Connected, h.AuthStatus = "awaiting_qr_login", false, authLabel
				h.FailureKind, h.Detail = string(spider.FailureAuth), detail
			})
			if !viper.GetBool("bili.QRLogin") {
				log.Printf("AUTH_EXPIRED: recording is stopped; QR login is disabled")
				if !wait(ctx, authRetryInterval) {
					break
				}
				continue
			}
			log.Printf("AUTH_EXPIRED: recording stopped; scan QR at http://%s/login/qr", healthAddr)
			if err := auth.LoginWithQRCode(ctx); err != nil {
				if ctx.Err() != nil {
					break
				}
				state.update(func(h *healthState) { h.Status, h.Detail = "qr_login_failed", err.Error() })
				log.Printf("QR_LOGIN_FAILED: %v", err)
				if !wait(ctx, authRetryInterval) {
					break
				}
				continue
			}
			log.Printf("QR_LOGIN_SUCCESS: login state saved; recording will resume")
			continue
		}
		state.update(func(h *healthState) { h.AuthStatus = authStatus.Label() })

		uid := auth.UID()
		if uid == 0 {
			uid = viper.GetInt64("bili.UID")
		}
		client := spider.NewDanmuSpider(shortID, uid, viper.GetString("bili.BUVID"), auth.Cookie())
		client.EndpointStateFile = valueOr(viper.GetString("bili.EndpointStateFile"), "conf/danmu_endpoint.json")
		state.update(func(h *healthState) {
			h.Status, h.Detail, h.FailureKind = "connecting", "", ""
			h.RestartCount++
		})
		runCtx, cancelRun := context.WithTimeout(ctx, authCheckInterval)
		err = client.Run(runCtx, func() {
			backoff = time.Second
			state.update(func(h *healthState) {
				h.Status, h.Connected, h.LastConnected = "ok", true, time.Now()
			})
		})
		periodicAuthCheck := errors.Is(runCtx.Err(), context.DeadlineExceeded)
		cancelRun()
		if ctx.Err() != nil {
			break
		}
		if periodicAuthCheck {
			log.Printf("AUTH_RECHECK: pausing connection for periodic login validation")
			continue
		}
		kind := spider.KindOf(err)
		state.update(func(h *healthState) {
			h.Status, h.Connected, h.FailureKind, h.Detail = "degraded", false, string(kind), err.Error()
		})
		log.Printf("DANMU_DISCONNECTED kind=%s retry_in=%s error=%v", kind, backoff, err)
		jitter := time.Duration(rand.Int63n(int64(backoff/4 + 1)))
		select {
		case <-ctx.Done():
		case <-time.After(backoff + jitter):
		}
		if backoff < time.Minute {
			backoff *= 2
			if backoff > time.Minute {
				backoff = time.Minute
			}
		}
	}
	log.Println("shutdown complete")
}

func shouldStartQRCode(status spider.AuthStatus, err error) bool {
	return !status.LoggedIn || (err != nil && spider.KindOf(err) == spider.FailureAuth)
}

func wait(ctx context.Context, duration time.Duration) bool {
	timer := time.NewTimer(duration)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return false
	case <-timer.C:
		return true
	}
}

func buildCookie(full, sessdata, buvid string) string {
	if strings.TrimSpace(full) != "" {
		return strings.TrimSpace(full)
	}
	parts := make([]string, 0, 2)
	if strings.TrimSpace(sessdata) != "" {
		// Backwards compatible: older deployments put "SESSDATA; bili_jct=..." here.
		parts = append(parts, "SESSDATA="+strings.TrimSpace(sessdata))
	}
	if strings.TrimSpace(buvid) != "" && !strings.Contains(strings.Join(parts, ";"), "buvid3=") {
		parts = append(parts, "buvid3="+strings.TrimSpace(buvid))
	}
	return strings.Join(parts, "; ")
}

func valueOr(value, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	return value
}
