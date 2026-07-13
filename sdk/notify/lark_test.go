package notify

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

func TestNotifyQRCode(t *testing.T) {
	var mu sync.Mutex
	var messageTypes []string
	var textContent string
	var uploaded []byte

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/open-apis/auth/v3/tenant_access_token/internal":
			body, _ := io.ReadAll(r.Body)
			if !strings.Contains(string(body), `"app_secret":"secret"`) {
				t.Errorf("token request did not contain credentials")
			}
			_, _ = io.WriteString(w, `{"code":0,"tenant_access_token":"token","expire":7200}`)
		case "/open-apis/im/v1/images":
			if r.Header.Get("Authorization") != "Bearer token" {
				t.Errorf("missing image authorization")
			}
			if err := r.ParseMultipartForm(1 << 20); err != nil {
				t.Fatalf("parse multipart: %v", err)
			}
			file, _, err := r.FormFile("image")
			if err != nil {
				t.Fatalf("get image: %v", err)
			}
			uploaded, _ = io.ReadAll(file)
			_ = file.Close()
			_, _ = io.WriteString(w, `{"code":0,"data":{"image_key":"img_key"}}`)
		case "/open-apis/im/v1/messages":
			var request struct {
				ReceiveID string `json:"receive_id"`
				MsgType   string `json:"msg_type"`
				Content   string `json:"content"`
			}
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				t.Fatalf("decode message: %v", err)
			}
			if request.ReceiveID != "chat" || r.URL.Query().Get("receive_id_type") != "chat_id" {
				t.Errorf("wrong recipient: %q", request.ReceiveID)
			}
			mu.Lock()
			messageTypes = append(messageTypes, request.MsgType)
			if request.MsgType == "text" {
				textContent = request.Content
			}
			mu.Unlock()
			_, _ = io.WriteString(w, `{"code":0}`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	notifier, err := NewLarkNotifier(LarkConfig{AppID: "app", AppSecret: "secret", ChatID: "chat", BaseURL: server.URL})
	if err != nil {
		t.Fatal(err)
	}
	if err := notifier.NotifyQRCode(context.Background(), "DanmuBackup", 404040887, []byte("png-data")); err != nil {
		t.Fatal(err)
	}
	if len(messageTypes) != 2 || messageTypes[0] != "text" || messageTypes[1] != "image" {
		t.Fatalf("unexpected messages: %#v", messageTypes)
	}
	if !strings.Contains(textContent, "DanmuBackup") || !strings.Contains(textContent, "404040887") {
		t.Fatalf("text does not identify instance and UID: %s", textContent)
	}
	if string(uploaded) != "png-data" {
		t.Fatalf("unexpected uploaded image: %q", uploaded)
	}
}
