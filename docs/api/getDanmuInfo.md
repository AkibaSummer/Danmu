# getDanmuInfo

## URL

> https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo?id=23075997

## Params

| 字段名 | 字段描述     | 字段样例 |
| ------ | ------------ | -------- |
| id     | 直播间房间号 | 23075997 |

## Example Resp

```json
    "code": 0,
    "message": "0",
    "ttl": 1,
    "data": {
        "group": "live",
        "business_id": 0,
        "refresh_row_factor": 0.125,
        "refresh_rate": 100,
        "max_delay": 5000,
        "token": "zUNKvLpv4-Led5AqjxG5YqPfF_VF6aXsQrgiXTYRB1Jths4rxXJ5etJgGGot36tY1hJP1FsmFZ4efV9DmhUNXBUyCewJILYzvC3Lmd-6NETyX-iktEtceymRShxsGUpr807Wc2lMuBg5tbqdFn8Kslf5I7yDyr4WEy4d9btltyJBcve5H4cHEdQ5",
        "host_list": [
            {
                "host": "zj-cn-live-comet.chat.bilibili.com",
                "port": 2243,
                "wss_port": 2245,
                "ws_port": 2244
            },
            {
                "host": "zj-cn-live-comet.chat.bilibili.com",
                "port": 2243,
                "wss_port": 2245,
                "ws_port": 2244
            },
            {
                "host": "ali-bj-live-comet-12.chat.bilibili.com",
                "port": 2243,
                "wss_port": 2245,
                "ws_port": 2244
            },
            {
                "host": "ali-sh-live-comet-10.chat.bilibili.com",
                "port": 2243,
                "wss_port": 2245,
                "ws_port": 2244
            },
            {
                "host": "broadcastlv.chat.bilibili.com",
                "port": 2243,
                "wss_port": 2245,
                "ws_port": 2244
            }
        ]
    }
}
```
