package spider

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	qrcode "github.com/skip2/go-qrcode"
)

const refreshPublicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDLgd2OAkcGVtoE3ThUREbio0Eg
Uc/prcajMKXvkCKFCWhJYJcLkcM2DKKcSeFpD/j6Boy538YXnR6VhcuUJOhH2x71
nzPjfdTcqMz7djHum0qSZA0AyCBDABUqCrfNgCiJ00Ra7GmRj+YCK1NJEuewlb40
JNrRuoEUXpabUzGB8QIDAQAB
-----END PUBLIC KEY-----`

type AuthConfig struct {
	Cookie       string
	RefreshToken string
	StateFile    string
	ExpectedUID  int64
	OnQRCode     func(context.Context, []byte) error
}

type AuthStatus struct {
	LoggedIn     bool
	NeedsRefresh bool
	Refreshed    bool
}

func (s AuthStatus) Label() string {
	if s.Refreshed {
		return "refreshed"
	}
	if s.LoggedIn {
		return "valid"
	}
	return "expired"
}

type authState struct {
	Cookie       string    `json:"cookie"`
	RefreshToken string    `json:"refresh_token"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AuthManager struct {
	mu          sync.RWMutex
	state       authState
	stateFile   string
	expectedUID int64
	onQRCode    func(context.Context, []byte) error
	http        *http.Client
	qrMu        sync.RWMutex
	qr          qrLoginState
}

type qrLoginState struct {
	PNG       []byte
	URL       string
	Status    string
	Message   string
	CreatedAt time.Time
}

func NewAuthManager(config AuthConfig) *AuthManager {
	return &AuthManager{
		state:       authState{Cookie: strings.TrimSpace(config.Cookie), RefreshToken: strings.TrimSpace(config.RefreshToken)},
		stateFile:   config.StateFile,
		expectedUID: config.ExpectedUID,
		onQRCode:    config.OnQRCode,
		http:        &http.Client{Timeout: 15 * time.Second},
	}
}

func (a *AuthManager) Cookie() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.state.Cookie
}

func (a *AuthManager) UID() int64 {
	uid, _ := strconv.ParseInt(cookieValue(a.Cookie(), "DedeUserID"), 10, 64)
	return uid
}

func (a *AuthManager) Load() error {
	if a.stateFile == "" {
		return nil
	}
	data, err := os.ReadFile(a.stateFile)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	var state authState
	if err := json.Unmarshal(data, &state); err != nil {
		return err
	}
	if state.Cookie == "" {
		return errors.New("auth state has empty cookie")
	}
	a.mu.Lock()
	a.state = state
	a.mu.Unlock()
	return nil
}

func (a *AuthManager) EnsureValid(ctx context.Context) (AuthStatus, error) {
	a.mu.RLock()
	current := a.state
	a.mu.RUnlock()
	if current.Cookie == "" {
		return AuthStatus{}, nil
	}
	info, err := a.cookieInfo(ctx, current.Cookie)
	if err != nil {
		return AuthStatus{}, err
	}
	if !info.LoggedIn {
		return AuthStatus{}, nil
	}
	status := AuthStatus{LoggedIn: true, NeedsRefresh: info.Refresh}
	if !info.Refresh {
		return status, nil
	}
	if current.RefreshToken == "" {
		return status, failure(FailureAuth, "refresh login", errors.New("cookie needs refresh but bili.RefreshToken is not configured"))
	}
	updated, err := a.refresh(ctx, current, info.Timestamp)
	if err != nil {
		return status, err
	}
	a.mu.Lock()
	a.state = updated
	a.mu.Unlock()
	if err := a.save(updated); err != nil {
		return status, failure(FailureConfig, "persist refreshed login", err)
	}
	status.Refreshed = true
	return status, nil
}

// LoginWithQRCode blocks until a QR login succeeds or the context is canceled.
// Expired QR codes are regenerated automatically while recording remains stopped.
func (a *AuthManager) LoginWithQRCode(ctx context.Context) error {
	for ctx.Err() == nil {
		var generated struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    struct {
				URL       string `json:"url"`
				QRCodeKey string `json:"qrcode_key"`
			} `json:"data"`
		}
		endpoint := "https://passport.bilibili.com/x/passport-login/web/qrcode/generate?source=main-fe-header"
		if err := a.requestJSON(ctx, http.MethodGet, endpoint, "", nil, &generated, nil); err != nil {
			a.setQR(nil, "error", err.Error(), "")
			return err
		}
		if generated.Code != 0 || generated.Data.URL == "" || generated.Data.QRCodeKey == "" {
			err := failure(FailureAuth, "generate login QR", fmt.Errorf("code=%d message=%s", generated.Code, generated.Message))
			a.setQR(nil, "error", err.Error(), "")
			return err
		}
		png, err := qrcode.Encode(generated.Data.URL, qrcode.Medium, 360)
		if err != nil {
			return failure(FailureProtocol, "encode login QR", err)
		}
		a.setQR(png, "waiting_scan", "请使用哔哩哔哩 App 扫码并确认", generated.Data.URL)
		if a.onQRCode != nil {
			notifyCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
			_ = a.onQRCode(notifyCtx, append([]byte(nil), png...))
			cancel()
		}

		deadline := time.NewTimer(185 * time.Second)
		ticker := time.NewTicker(2 * time.Second)
		regenerate := false
		for !regenerate {
			select {
			case <-ctx.Done():
				ticker.Stop()
				deadline.Stop()
				return ctx.Err()
			case <-deadline.C:
				ticker.Stop()
				regenerate = true
			case <-ticker.C:
				result, cookies, pollErr := a.pollQRCode(ctx, generated.Data.QRCodeKey)
				if pollErr != nil {
					ticker.Stop()
					deadline.Stop()
					a.setQR(png, "error", pollErr.Error(), generated.Data.URL)
					return pollErr
				}
				switch result.Code {
				case 0:
					ticker.Stop()
					deadline.Stop()
					return a.finishQRCodeLogin(ctx, result, cookies)
				case 86101:
					a.setQR(png, "waiting_scan", "等待扫码", generated.Data.URL)
				case 86090:
					a.setQR(png, "waiting_confirm", "已扫码，请在手机上确认", generated.Data.URL)
				case 86038, 86083:
					ticker.Stop()
					deadline.Stop()
					regenerate = true
					a.setQR(png, "expired", "二维码已失效，正在重新生成", generated.Data.URL)
				default:
					ticker.Stop()
					deadline.Stop()
					err := failure(FailureAuth, "poll login QR", fmt.Errorf("code=%d message=%s", result.Code, result.Message))
					a.setQR(png, "error", err.Error(), generated.Data.URL)
					return err
				}
			}
		}
	}
	return ctx.Err()
}

type qrPollResult struct {
	URL          string `json:"url"`
	RefreshToken string `json:"refresh_token"`
	Timestamp    int64  `json:"timestamp"`
	Code         int    `json:"code"`
	Message      string `json:"message"`
}

func (a *AuthManager) pollQRCode(ctx context.Context, key string) (qrPollResult, []*http.Cookie, error) {
	endpoint := "https://passport.bilibili.com/x/passport-login/web/qrcode/poll?source=main-fe-header&qrcode_key=" + url.QueryEscape(key)
	var response struct {
		Code    int          `json:"code"`
		Message string       `json:"message"`
		Data    qrPollResult `json:"data"`
	}
	var cookies []*http.Cookie
	if err := a.requestJSON(ctx, http.MethodGet, endpoint, "", nil, &response, &cookies); err != nil {
		return qrPollResult{}, nil, err
	}
	if response.Code != 0 {
		return qrPollResult{}, nil, failure(FailureAuth, "poll login QR", fmt.Errorf("code=%d message=%s", response.Code, response.Message))
	}
	return response.Data, cookies, nil
}

func (a *AuthManager) finishQRCodeLogin(ctx context.Context, result qrPollResult, cookies []*http.Cookie) error {
	newCookie := mergeCookies(a.Cookie(), cookies)
	if redirect, err := url.Parse(result.URL); err == nil {
		for _, name := range []string{"SESSDATA", "bili_jct", "DedeUserID", "DedeUserID__ckMd5", "sid"} {
			if value := redirect.Query().Get(name); value != "" {
				newCookie = mergeCookies(newCookie, []*http.Cookie{{Name: name, Value: value}})
			}
		}
	}
	info, err := a.cookieInfo(ctx, newCookie)
	if err != nil {
		return err
	}
	if !info.LoggedIn {
		return failure(FailureAuth, "verify QR login", errors.New("Bilibili did not accept the new cookie"))
	}
	if scannedUID, _ := strconv.ParseInt(cookieValue(newCookie, "DedeUserID"), 10, 64); a.expectedUID > 0 && scannedUID != a.expectedUID {
		return failure(FailureAuth, "verify QR login", fmt.Errorf("scanned account UID %d does not match configured UID %d", scannedUID, a.expectedUID))
	}
	state := authState{Cookie: newCookie, RefreshToken: result.RefreshToken, UpdatedAt: time.Now()}
	if err := a.save(state); err != nil {
		return failure(FailureConfig, "persist QR login", err)
	}
	a.mu.Lock()
	a.state = state
	a.mu.Unlock()
	a.setQR(nil, "success", "登录成功，录制即将恢复", "")
	return nil
}

func (a *AuthManager) setQR(png []byte, status, message, loginURL string) {
	a.qrMu.Lock()
	defer a.qrMu.Unlock()
	a.qr = qrLoginState{PNG: png, URL: loginURL, Status: status, Message: message, CreatedAt: time.Now()}
}

func (a *AuthManager) QRPage(w http.ResponseWriter, _ *http.Request) {
	a.qrMu.RLock()
	state := a.qr
	a.qrMu.RUnlock()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	_, _ = fmt.Fprintf(w, `<!doctype html><html lang="zh-CN"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width"><meta http-equiv="refresh" content="3"><title>Danmu 扫码登录</title><style>body{font-family:sans-serif;text-align:center;margin:2rem;background:#f6f7f8;color:#18191c}main{display:inline-block;background:white;padding:2rem;border-radius:16px;box-shadow:0 4px 24px #0002}img{width:360px;max-width:90vw}code{display:block;margin-top:1rem}</style></head><body><main><h1>哔哩哔哩扫码登录</h1><p>%s</p><p>状态：<code>%s</code></p>%s</main></body></html>`, html.EscapeString(state.Message), html.EscapeString(state.Status), qrImageTag(state))
}

func qrImageTag(state qrLoginState) string {
	if len(state.PNG) == 0 {
		return ""
	}
	return fmt.Sprintf(`<img src="/login/qr.png?t=%d" alt="登录二维码"><p>请使用哔哩哔哩 App 扫码并在手机上确认</p>`, state.CreatedAt.UnixNano())
}

func (a *AuthManager) QRPNG(w http.ResponseWriter, _ *http.Request) {
	a.qrMu.RLock()
	png := append([]byte(nil), a.qr.PNG...)
	a.qrMu.RUnlock()
	if len(png) == 0 {
		http.Error(w, "QR code is not available", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "no-store")
	_, _ = w.Write(png)
}

type cookieInfoResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Refresh   bool  `json:"refresh"`
		Timestamp int64 `json:"timestamp"`
	} `json:"data"`
}

type cookieInfo struct {
	LoggedIn, Refresh bool
	Timestamp         int64
}

func (a *AuthManager) cookieInfo(ctx context.Context, cookie string) (cookieInfo, error) {
	csrf := cookieValue(cookie, "bili_jct")
	endpoint := "https://passport.bilibili.com/x/passport-login/web/cookie/info?csrf=" + url.QueryEscape(csrf)
	var response cookieInfoResponse
	if err := a.requestJSON(ctx, http.MethodGet, endpoint, cookie, nil, &response, nil); err != nil {
		return cookieInfo{}, err
	}
	if response.Code == -101 {
		return cookieInfo{}, nil
	}
	if response.Code != 0 {
		return cookieInfo{}, failure(FailureAuth, "check login", fmt.Errorf("code=%d message=%s", response.Code, response.Message))
	}
	return cookieInfo{LoggedIn: true, Refresh: response.Data.Refresh, Timestamp: response.Data.Timestamp}, nil
}

func (a *AuthManager) refresh(ctx context.Context, old authState, timestamp int64) (authState, error) {
	path, err := correspondPath(timestamp)
	if err != nil {
		return authState{}, failure(FailureProtocol, "build refresh path", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.bilibili.com/correspond/1/"+path, nil)
	if err != nil {
		return authState{}, err
	}
	req.Header.Set("Cookie", old.Cookie)
	req.Header.Set("User-Agent", userAgent)
	resp, err := a.http.Do(req)
	if err != nil {
		return authState{}, failure(FailureNetwork, "get refresh csrf", err)
	}
	body, readErr := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	resp.Body.Close()
	if readErr != nil {
		return authState{}, failure(FailureNetwork, "read refresh csrf", readErr)
	}
	match := regexp.MustCompile(`<div\s+id=["']1-name["']>([^<]+)</div>`).FindSubmatch(body)
	if resp.StatusCode != http.StatusOK || len(match) != 2 {
		return authState{}, failure(FailureAuth, "get refresh csrf", fmt.Errorf("status=%s token_found=%t", resp.Status, len(match) == 2))
	}

	form := url.Values{"csrf": {cookieValue(old.Cookie, "bili_jct")}, "refresh_csrf": {html.UnescapeString(string(match[1]))}, "source": {"main_web"}, "refresh_token": {old.RefreshToken}}
	var refreshed struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			RefreshToken string `json:"refresh_token"`
		} `json:"data"`
	}
	var setCookies []*http.Cookie
	if err := a.requestJSON(ctx, http.MethodPost, "https://passport.bilibili.com/x/passport-login/web/cookie/refresh", old.Cookie, form, &refreshed, &setCookies); err != nil {
		return authState{}, err
	}
	if refreshed.Code != 0 || refreshed.Data.RefreshToken == "" {
		return authState{}, failure(FailureAuth, "refresh login", fmt.Errorf("code=%d message=%s", refreshed.Code, refreshed.Message))
	}
	newCookie := mergeCookies(old.Cookie, setCookies)
	confirm := url.Values{"csrf": {cookieValue(newCookie, "bili_jct")}, "refresh_token": {old.RefreshToken}}
	var confirmed struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	if err := a.requestJSON(ctx, http.MethodPost, "https://passport.bilibili.com/x/passport-login/web/confirm/refresh", newCookie, confirm, &confirmed, nil); err != nil {
		return authState{}, err
	}
	if confirmed.Code != 0 {
		return authState{}, failure(FailureAuth, "confirm refreshed login", fmt.Errorf("code=%d message=%s", confirmed.Code, confirmed.Message))
	}
	return authState{Cookie: newCookie, RefreshToken: refreshed.Data.RefreshToken, UpdatedAt: time.Now()}, nil
}

func (a *AuthManager) requestJSON(ctx context.Context, method, endpoint, cookie string, form url.Values, dst interface{}, setCookies *[]*http.Cookie) error {
	var body io.Reader
	if form != nil {
		body = strings.NewReader(form.Encode())
	}
	req, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		return failure(FailureConfig, "create auth request", err)
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Cookie", cookie)
	if form != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	resp, err := a.http.Do(req)
	if err != nil {
		return failure(FailureNetwork, "auth HTTP request", err)
	}
	defer resp.Body.Close()
	if setCookies != nil {
		*setCookies = resp.Cookies()
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return failure(FailureAPI, "auth HTTP request", fmt.Errorf("status=%s", resp.Status))
	}
	if err := json.NewDecoder(io.LimitReader(resp.Body, 1<<20)).Decode(dst); err != nil {
		return failure(FailureProtocol, "decode auth response", err)
	}
	return nil
}

func (a *AuthManager) save(state authState) error {
	if a.stateFile == "" {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(a.stateFile), 0700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	tmp := a.stateFile + ".tmp"
	if err := os.WriteFile(tmp, data, 0600); err != nil {
		return err
	}
	return os.Rename(tmp, a.stateFile)
}

func correspondPath(timestamp int64) (string, error) {
	block, _ := pem.Decode([]byte(refreshPublicKey))
	if block == nil {
		return "", errors.New("invalid public key PEM")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	publicKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return "", errors.New("refresh key is not RSA")
	}
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(fmt.Sprintf("refresh_%d", timestamp)), nil)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(ciphertext), nil
}

func cookieValue(header, name string) string {
	for _, part := range strings.Split(header, ";") {
		pair := strings.SplitN(strings.TrimSpace(part), "=", 2)
		if len(pair) == 2 && pair[0] == name {
			return pair[1]
		}
	}
	return ""
}

func mergeCookies(header string, updates []*http.Cookie) string {
	values := map[string]string{}
	for _, part := range strings.Split(header, ";") {
		pair := strings.SplitN(strings.TrimSpace(part), "=", 2)
		if len(pair) == 2 {
			values[pair[0]] = pair[1]
		}
	}
	for _, cookie := range updates {
		if cookie.MaxAge < 0 {
			delete(values, cookie.Name)
		} else if cookie.Name != "" {
			values[cookie.Name] = cookie.Value
		}
	}
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, key := range keys {
		parts = append(parts, key+"="+values[key])
	}
	return strings.Join(parts, "; ")
}
