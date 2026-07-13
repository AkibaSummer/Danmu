package notify

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"sync"
	"time"
)

type LarkConfig struct {
	AppID     string
	AppSecret string
	ChatID    string
	BaseURL   string
}

type LarkNotifier struct {
	config    LarkConfig
	http      *http.Client
	mu        sync.Mutex
	token     string
	expiresAt time.Time
}

func NewLarkNotifier(config LarkConfig) (*LarkNotifier, error) {
	if strings.TrimSpace(config.AppID) == "" || strings.TrimSpace(config.AppSecret) == "" || strings.TrimSpace(config.ChatID) == "" {
		return nil, errors.New("lark app_id, app_secret and chat_id are required")
	}
	if config.BaseURL == "" {
		config.BaseURL = "https://open.feishu.cn"
	}
	return &LarkNotifier{config: config, http: &http.Client{Timeout: 15 * time.Second}}, nil
}

func (n *LarkNotifier) NotifyQRCode(ctx context.Context, instance string, uid int64, png []byte) error {
	if len(png) == 0 {
		return errors.New("QR image is empty")
	}
	message := fmt.Sprintf("【Danmu 登录过期提醒】\n实例：%s\n登录账号 UID：%d\n录制状态：该实例已停止录制\n操作：请使用此 UID 对应的哔哩哔哩账号扫描下一条二维码并在手机确认。\n时间：%s",
		instance, uid, time.Now().Format("2006-01-02 15:04:05 MST"))
	if err := n.sendMessage(ctx, "text", map[string]string{"text": message}); err != nil {
		return fmt.Errorf("send lark alert text: %w", err)
	}
	imageKey, err := n.uploadImage(ctx, png)
	if err != nil {
		return fmt.Errorf("upload lark QR image: %w", err)
	}
	if err := n.sendMessage(ctx, "image", map[string]string{"image_key": imageKey}); err != nil {
		return fmt.Errorf("send lark QR image: %w", err)
	}
	return nil
}

func (n *LarkNotifier) sendMessage(ctx context.Context, messageType string, contentValue any) error {
	content, err := json.Marshal(contentValue)
	if err != nil {
		return err
	}
	body := map[string]any{"receive_id": n.config.ChatID, "msg_type": messageType, "content": string(content)}
	return n.authorizedJSON(ctx, n.config.BaseURL+"/open-apis/im/v1/messages?receive_id_type=chat_id", body)
}

func (n *LarkNotifier) uploadImage(ctx context.Context, png []byte) (string, error) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	if err := writer.WriteField("image_type", "message"); err != nil {
		return "", err
	}
	part, err := writer.CreateFormFile("image", "bilibili-login-qr.png")
	if err != nil {
		return "", err
	}
	if _, err := part.Write(png); err != nil {
		return "", err
	}
	if err := writer.Close(); err != nil {
		return "", err
	}

	token, err := n.accessToken(ctx)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, n.config.BaseURL+"/open-apis/im/v1/images", &body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := n.http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var result struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			ImageKey string `json:"image_key"`
		} `json:"data"`
	}
	if err := decodeResponse(resp, &result); err != nil {
		return "", err
	}
	if result.Code != 0 || result.Data.ImageKey == "" {
		return "", fmt.Errorf("code=%d msg=%s", result.Code, result.Msg)
	}
	return result.Data.ImageKey, nil
}

func (n *LarkNotifier) authorizedJSON(ctx context.Context, endpoint string, body any) error {
	token, err := n.accessToken(ctx)
	if err != nil {
		return err
	}
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := n.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var result struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	if err := decodeResponse(resp, &result); err != nil {
		return err
	}
	if result.Code != 0 {
		return fmt.Errorf("code=%d msg=%s", result.Code, result.Msg)
	}
	return nil
}

func (n *LarkNotifier) accessToken(ctx context.Context) (string, error) {
	n.mu.Lock()
	defer n.mu.Unlock()
	if n.token != "" && time.Until(n.expiresAt) > 5*time.Minute {
		return n.token, nil
	}
	body, err := json.Marshal(map[string]string{"app_id": n.config.AppID, "app_secret": n.config.AppSecret})
	if err != nil {
		return "", err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, n.config.BaseURL+"/open-apis/auth/v3/tenant_access_token/internal", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := n.http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var result struct {
		Code   int    `json:"code"`
		Msg    string `json:"msg"`
		Token  string `json:"tenant_access_token"`
		Expire int    `json:"expire"`
	}
	if err := decodeResponse(resp, &result); err != nil {
		return "", err
	}
	if result.Code != 0 || result.Token == "" {
		return "", fmt.Errorf("code=%d msg=%s", result.Code, result.Msg)
	}
	n.token = result.Token
	n.expiresAt = time.Now().Add(time.Duration(result.Expire) * time.Second)
	return n.token, nil
}

func decodeResponse(resp *http.Response, dst any) error {
	data, err := io.ReadAll(io.LimitReader(resp.Body, 2<<20))
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("HTTP %s: %s", resp.Status, strings.TrimSpace(string(data)))
	}
	return json.Unmarshal(data, dst)
}
