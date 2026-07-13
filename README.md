# Danmu

Bilibili 直播弹幕录制机器人，适合 systemd 长期运行。

- 网络异常、节点关闭和读写超时会自动重连，不再触发 `panic`。
- 健康状态区分 `auth`、`network`、`api`、`protocol` 和 `config` 故障。
- 登录态过期或无法确认时采用 fail-closed：立即停止该实例录制，不匿名降级。
- 有 `refresh_token` 时自动续期；无法续期时生成扫码二维码。
- 可将登录账号 UID 和二维码发送到飞书群。
- Bilibili 节点接口遭遇 `-352` 时可复用最近成功的节点缓存。

## 配置

```bash
cp conf/conf.example.yaml conf/conf.yaml
```

填写直播间、账号 UID 和完整 Cookie。程序会把扫码或续期得到的新 Cookie、`refresh_token` 原子写入权限为 `0600` 的 `conf/auth_state.json`。`conf/conf.yaml` 和状态文件均被 Git 忽略。

飞书凭据建议单独存放，不要写进仓库：

```bash
install -d -m 700 /home/ecs-user/.config/danmu
install -m 600 /dev/null /home/ecs-user/.config/danmu/lark.env
```

`lark.env` 使用以下变量：

```text
LARK_APPID=cli_xxx
LARK_APPSECRET=xxx
LARK_CHATID=oc_xxx
```

登录过期时，该实例先停录，再向群内发送实例名、配置 UID 和原生二维码图片。通知发送失败不会阻塞本机扫码页和后续登录轮询。

## 运行与诊断

```bash
go build -o Danmu .
./Danmu
curl -i http://127.0.0.1:10216/healthz
```

登录过期时，可通过 SSH 隧道访问本机扫码页：

```bash
ssh -L 10216:127.0.0.1:10216 ecs-user@47.100.18.57
```

然后打开 `http://127.0.0.1:10216/login/qr`。扫码 UID 与配置 UID 不一致时，程序会拒绝保存登录态。

健康接口判定：

- HTTP 200、`status: ok`：进程和 Bilibili WebSocket 均正常。
- HTTP 503、`failure_kind: auth`：登录过期、续期失败或无法确认，录制已停止。
- HTTP 503、`failure_kind: network`：服务器在线，但到 Bilibili 的网络或 WebSocket 异常。
- 健康接口不可达但 SSH 正常：检查 systemd 和日志。
- 健康接口与 SSH 均不可达：检查 ECS、网络和安全组。

## systemd 部署

```bash
sudo cp deploy/danmu@.service /etc/systemd/system/danmu@.service
sudo systemctl daemon-reload
sudo systemctl enable --now danmu@Danmu danmu@DanmuBackup
systemctl status danmu@Danmu danmu@DanmuBackup
journalctl -u 'danmu@*' -f
```

两个实例正常情况下同时录制，健康端口分别设为 `10216` 和 `10217`。

版本升级采用备份实例先行：

```bash
chmod +x deploy/rolling-update.sh
deploy/rolling-update.sh /path/to/new/Danmu
```

脚本先升级 `DanmuBackup` 并等待它恢复 `ok/connected`，成功后才升级 `Danmu`。任一实例验证失败会回滚该实例并停止继续发布。
