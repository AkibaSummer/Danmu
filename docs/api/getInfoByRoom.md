# getInfoByRoom

## URL

> https://api.live.bilibili.com/xlive/web-room/v1/index/getInfoByRoom?room_id=23075997

## Params

| 字段名  | 字段描述     | 字段样例 |
| ------- | ------------ | -------- |
| room_id | 直播间房间号 | 23075997 |

## Example Resp

- **直播间状态：LIVE**

  ```json
  {
      "code": 0,
      "message": "0",
      "ttl": 1,u
      "data": {
          "room_info": {
              "uid": 1355412269,
              "room_id": 23075997,
              "short_id": 0,
              "title": "该唠唠了",
              "cover": "http://i0.hdslb.com/bfs/live/new_room_cover/fb038f94f5f8a24d0bff4c17de0409445a3d8069.jpg",
              "tags": "",
              "background": "http://i0.hdslb.com/bfs/live/room_bg/d62474f3335ebe3c6543977d4c8aba5078cfdb41.jpg",
              "description": "✿能想起来的部分歌单：http://suiii.minamini.cn/",
              "live_status": 1,
              "live_start_time": 1718281912,
              "live_screen_type": 0,
              "lock_status": 0,
              "lock_time": 0,
              "hidden_status": 0,
              "hidden_time": 0,
              "area_id": 744,
              "area_name": "虚拟Singer",
              "parent_area_id": 9,
              "parent_area_name": "虚拟主播",
              "keyframe": "http://i0.hdslb.com/bfs/live-key-frame/keyframe06132201000023075997vc633n.jpg",
              "special_type": 0,
              "up_session": "504307921683750045",
              "pk_status": 0,
              "is_studio": false,
              "pendants": {
                  "frame": {
                      "name": "百人舰队主播头像",
                      "value": "https://i0.hdslb.com/bfs/vc/071eb10548fe9bc482ff69331983d94192ce9507.png",
                      "desc": ""
                  }
              },
              "on_voice_join": 0,
              "online": 71153,
              "room_type": {
                  "3-21": 0,
                  "3-50": 1,
                  "3-51": 0
              },
              "sub_session_key": "504307921683750045sub_time:1718281912",
              "live_id": 504307921683750045,
              "live_id_str": "504307921683750045",
              "official_room_id": 0,
              "official_room_info": null,
              "voice_background": ""
          },
          "anchor_info": {
              "base_info": {
                  "uname": "随一Suiii",
                  "face": "https://i2.hdslb.com/bfs/face/131b8eb3bf0c347df5bb078a49a35f6f669cb93f.jpg",
                  "gender": "保密",
                  "official_info": {
                      "role": 0,
                      "title": "bilibili直播高能主播",
                      "desc": "",
                      "is_nft": 0,
                      "nft_dmark": "https://i0.hdslb.com/bfs/live/9f176ff49d28c50e9c53ec1c3297bd1ee539b3d6.gif"
                  }
              },
              "live_info": {
                  "level": 32,
                  "level_color": 16746162,
                  "score": 21221461,
                  "upgrade_score": 6092349,
                  "current": [
                      5000000,
                      20613810
                  ],
                  "next": [
                      6700000,
                      27313810
                  ],
                  "rank": "2083"
              },
              "relation_info": {
                  "attention": 28332
              },
              "medal_info": {
                  "medal_name": "随随冰",
                  "medal_id": 448828,
                  "fansclub": 832
              },
              "gift_info": {
                  "price": 0,
                  "price_update_time": 0
              }
          },
          "news_info": {
              "uid": 1355412269,
              "ctime": "2023-02-19 11:59:46",
              "content": "午播恢复中/晚上八点半唱歌杂谈为主，偶尔游戏"
          },
          "rankdb_info": {
              "roomid": 23075997,
              "rank_desc": "",
              "color": "",
              "h5_url": "",
              "web_url": "",
              "timestamp": 1718287290
          },
          "area_rank_info": {
              "areaRank": {
                  "index": 0,
                  "rank": "\u003e1000"
              },
              "liveRank": {
                  "rank": "2083"
              }
          },
          "battle_rank_entry_info": null,
          "tab_info": {
              "list": [
                  {
                      "type": "seven-rank",
                      "desc": "高能用户",
                      "isFirst": 1,
                      "isEvent": 0,
                      "eventType": "",
                      "listType": "",
                      "apiPrefix": "",
                      "rank_name": "room_7day"
                  },
                  {
                      "type": "guard",
                      "desc": "大航海",
                      "isFirst": 0,
                      "isEvent": 0,
                      "eventType": "",
                      "listType": "top-list",
                      "apiPrefix": "",
                      "rank_name": ""
                  }
              ]
          },
          "activity_init_info": {
              "eventList": [],
              "weekInfo": {
                  "bannerInfo": null,
                  "giftName": null
              },
              "giftName": null,
              "lego": {
                  "timestamp": 1718287290,
                  "config": "[{\"name\":\"frame-mng\",\"url\":\"https:\\/\\/live.bilibili.com\\/p\\/html\\/live-web-mng\\/index.html?roomid=#roomid#\u0026arae_id=#area_id#\u0026parent_area_id=#parent_area_id#\u0026ruid=#ruid#\",\"startTime\":1559544736,\"endTime\":1877167950,\"type\":\"frame-mng\"},{\"name\":\"s10-fun\",\"target\":\"sidebar\",\"icon\":\"https:\\/\\/i0.hdslb.com\\/bfs\\/activity-plat\\/static\\/20200908\\/3435f7521efc759ae1f90eae5629a8f0\\/HpxrZ7SOT.png\",\"text\":\"\\u7545\\u73a9s10\",\"url\":\"https:\\/\\/live.bilibili.com\\/s10\\/fun\\/index.html?room_id=#roomid#\u0026width=376\u0026height=600\u0026source=sidebar\",\"color\":\"#2e6fc0\",\"startTime\":1600920000,\"endTime\":1604721600,\"parentAreaId\":2,\"areaId\":86},{\"name\":\"genshin-avatar\",\"target\":\"sidebar\",\"icon\":\"https:\\/\\/i0.hdslb.com\\/bfs\\/activity-plat\\/static\\/20210721\\/fa538c98e9e32dc98919db4f2527ad02\\/qWxN1d0ACu.jpg\",\"text\":\"\\u539f\\u77f3\\u798f\\u5229\",\"url\":\"https:\\/\\/live.bilibili.com\\/activity\\/live-activity-full\\/genshin_avatar\\/mobile.html?no-jump=1\u0026room_id=#roomid#\u0026width=376\u0026height=550#\\/\",\"color\":\"#2e6fc0\",\"frameAllowNoBg\":\"1\",\"frameAllowDrag\":\"1\",\"startTime\":1627012800,\"endTime\":1630425540,\"parentAreaId\":3,\"areaId\":321}]"
              }
          },
          "voice_join_info": {
              "status": {
                  "open": 0,
                  "anchor_open": 0,
                  "status": 0,
                  "uid": 0,
                  "user_name": "",
                  "head_pic": "",
                  "guard": 0,
                  "start_at": 0,
                  "current_time": 1718287290
              },
              "icons": {
                  "icon_close": "https://i0.hdslb.com/bfs/live/a176d879dffe8de1586a5eb54c2a08a0c7d31392.png",
                  "icon_open": "https://i0.hdslb.com/bfs/live/70f0844c9a12d29db1e586485954290144534be9.png",
                  "icon_wait": "https://i0.hdslb.com/bfs/live/1049bb88f1e7afd839cc1de80e13228ccd5807e8.png",
                  "icon_starting": "https://i0.hdslb.com/bfs/live/948433d1647a0704f8216f017c406224f9fff518.gif"
              },
              "web_share_link": "https://live.bilibili.com/h5/23075997"
          },
          "ad_banner_info": {
              "data": null
          },
          "skin_info": {
              "id": 65,
              "skin_name": "百人舰队皮肤",
              "skin_config": "{\"zip\":\"http://i0.hdslb.com/bfs/live/fe21bb238c7ca435dc09e78a4b838d78b3435dbd.zip\",\"md5\":\"2AEB56D860055A3F897FDB4BC481DFEE\",\"platform\":\"web\",\"version\":\"1\",\"headInfoBgPic\":\"http://i0.hdslb.com/bfs/live/7ba5a32cda0f985aa02bb05f453eac1f03cb976d.png\",\"giftControlBgPic\":\"http://i0.hdslb.com/bfs/live/f864809e6be7b3fe834de6d37b5b7f42b2cdff2a.png\",\"rankListBgPic\":\"http://i0.hdslb.com/bfs/live/00d1718591af1b5117c588f7bac1efb6c1e97fde.png\",\"mainText\":\"#FFffd423\",\"normalText\":\"#FFffffff\",\"highlightContent\":\"#FFffd432\",\"border\":\"#FF999999\",\"infoCardBgPic\":\"\"}",
              "show_text": "保持直播间大航海人数100-999人",
              "skin_url": "https://i0.hdslb.com/bfs/live/10143ee1233a4b88dc902a4dc92c9f30a4ca0241.png",
              "start_time": 1715260781,
              "end_time": 2145888000,
              "current_time": 1718287290
          },
          "web_banner_info": {
              "id": 0,
              "title": "",
              "left": "",
              "right": "",
              "jump_url": "",
              "bg_color": "",
              "hover_color": "",
              "text_bg_color": "",
              "text_hover_color": "",
              "link_text": "",
              "link_color": "",
              "input_color": "",
              "input_text_color": "",
              "input_hover_color": "",
              "input_border_color": "",
              "input_search_color": ""
          },
          "lol_info": null,
          "pk_info": null,
          "battle_info": null,
          "silent_room_info": {
              "type": "",
              "level": 0,
              "second": 0,
              "expire_time": 0
          },
          "switch_info": {
              "close_guard": false,
              "close_gift": false,
              "close_online": false,
              "close_danmaku": false
          },
          "record_switch_info": null,
          "room_config_info": {
              "dm_text": "发个弹幕呗~"
          },
          "gift_memory_info": {
              "list": null
          },
          "new_switch_info": {
              "room-socket": 1,
              "room-prop-send": 1,
              "room-sailing": 1,
              "room-info-popularity": 1,
              "room-danmaku-editor": 1,
              "room-effect": 1,
              "room-fans_medal": 1,
              "room-report": 1,
              "room-feedback": 1,
              "room-player-watermark": 1,
              "room-recommend-live_off": 1,
              "room-activity": 1,
              "room-web_banner": 1,
              "room-silver_seeds-box": 1,
              "room-wishing_bottle": 1,
              "room-board": 1,
              "room-supplication": 1,
              "room-hour_rank": 1,
              "room-week_rank": 1,
              "room-anchor_rank": 1,
              "room-info-integral": 1,
              "room-super-chat": 1,
              "room-tab": 1,
              "room-hot-rank": 1,
              "fans-medal-progress": 1,
              "gift-bay-screen": 1,
              "room-enter": 1,
              "room-my-idol": 1,
              "room-topic": 1,
              "fans-club": 1,
              "room-popular-rank": 1,
              "mic_user_gift": 1,
              "new-room-area-rank": 1,
              "wealth_medal": 1,
              "bubble": 1,
              "title": 1,
              "room_rank_rearrange": 1,
              "web-gift-batter-bar": 1,
              "popular_rank_anchor_ab": 1
          },
          "super_chat_info": {
              "status": 1,
              "jump_url": "https://live.bilibili.com/p/html/live-app-superchat2/index.html?is_live_half_webview=1\u0026hybrid_half_ui=1,3,100p,70p,ffffff,0,30,100,12,0;2,2,375,100p,ffffff,0,30,100,0,0;3,3,100p,70p,ffffff,0,30,100,12,0;4,2,375,100p,ffffff,0,30,100,0,0;5,3,100p,60p,ffffff,0,30,100,12,0;6,3,100p,60p,ffffff,0,30,100,12,0;7,3,100p,60p,ffffff,0,30,100,12,0",
              "icon": "https://i0.hdslb.com/bfs/live/0a9ebd72c76e9cbede9547386dd453475d4af6fe.png",
              "ranked_mark": 0,
              "message_list": []
          },
          "online_gold_rank_info_v2": {
              "list": [
                  {
                      "uid": 405771212,
                      "face": "https://i0.hdslb.com/bfs/face/member/noface.jpg",
                      "uname": "---愤怒的小鸟---",
                      "score": "477",
                      "rank": 1,
                      "guard_level": 3,
                      "wealth_level": 27,
                      "is_mystery": false,
                      "uinfo": {
                          "uid": 405771212,
                          "base": {
                              "name": "---愤怒的小鸟---",
                              "face": "https://i0.hdslb.com/bfs/face/member/noface.jpg",
                              "name_color": 0,
                              "is_mystery": false,
                              "risk_ctrl_info": {
                                  "name": "---愤怒的小鸟---",
                                  "face": "https://i0.hdslb.com/bfs/face/member/noface.jpg"
                              },
                              "origin_info": {
                                  "name": "---愤怒的小鸟---",
                                  "face": "https://i0.hdslb.com/bfs/face/member/noface.jpg"
                              },
                              "official_info": {
                                  "role": 0,
                                  "title": "",
                                  "desc": "",
                                  "type": -1
                              },
                              "name_color_str": ""
                          },
                          "medal": null,
                          "wealth": null,
                          "title": null,
                          "guard": null,
                          "uhead_frame": null,
                          "guard_leader": null
                      }
                  },
                  {
                      "uid": 287837821,
                      "face": "https://i1.hdslb.com/bfs/face/bb26f0ed92a1ca9f7836de95f74c4a4945a272a3.jpg",
                      "uname": "境花与蝶",
                      "score": "405",
                      "rank": 2,
                      "guard_level": 3,
                      "wealth_level": 41,
                      "is_mystery": false,
                      "uinfo": {
                          "uid": 287837821,
                          "base": {
                              "name": "境花与蝶",
                              "face": "https://i1.hdslb.com/bfs/face/bb26f0ed92a1ca9f7836de95f74c4a4945a272a3.jpg",
                              "name_color": 0,
                              "is_mystery": false,
                              "risk_ctrl_info": {
                                  "name": "境花与蝶",
                                  "face": "https://i1.hdslb.com/bfs/face/bb26f0ed92a1ca9f7836de95f74c4a4945a272a3.jpg"
                              },
                              "origin_info": {
                                  "name": "境花与蝶",
                                  "face": "https://i1.hdslb.com/bfs/face/bb26f0ed92a1ca9f7836de95f74c4a4945a272a3.jpg"
                              },
                              "official_info": {
                                  "role": 0,
                                  "title": "",
                                  "desc": "",
                                  "type": -1
                              },
                              "name_color_str": ""
                          },
                          "medal": null,
                          "wealth": null,
                          "title": null,
                          "guard": null,
                          "uhead_frame": null,
                          "guard_leader": null
                      }
                  },
                  {
                      "uid": 156077349,
                      "face": "https://i0.hdslb.com/bfs/face/32fdee8a66a61d38b6e65d3310363c6af46480a2.jpg",
                      "uname": "雨齐青",
                      "score": "282",
                      "rank": 3,
                      "guard_level": 2,
                      "wealth_level": 42,
                      "is_mystery": false,
                      "uinfo": {
                          "uid": 156077349,
                          "base": {
                              "name": "雨齐青",
                              "face": "https://i0.hdslb.com/bfs/face/32fdee8a66a61d38b6e65d3310363c6af46480a2.jpg",
                              "name_color": 0,
                              "is_mystery": false,
                              "risk_ctrl_info": {
                                  "name": "雨齐青",
                                  "face": "https://i0.hdslb.com/bfs/face/32fdee8a66a61d38b6e65d3310363c6af46480a2.jpg"
                              },
                              "origin_info": {
                                  "name": "雨齐青",
                                  "face": "https://i0.hdslb.com/bfs/face/32fdee8a66a61d38b6e65d3310363c6af46480a2.jpg"
                              },
                              "official_info": {
                                  "role": 0,
                                  "title": "",
                                  "desc": "",
                                  "type": -1
                              },
                              "name_color_str": ""
                          },
                          "medal": null,
                          "wealth": null,
                          "title": null,
                          "guard": null,
                          "uhead_frame": null,
                          "guard_leader": null
                      }
                  },
                  {
                      "uid": 21048761,
                      "face": "https://i2.hdslb.com/bfs/face/82d1561e3df7cfc73176d06a5de750bdcc66e89b.jpg",
                      "uname": "随喵的令和",
                      "score": "181",
                      "rank": 4,
                      "guard_level": 3,
                      "wealth_level": 38,
                      "is_mystery": false,
                      "uinfo": {
                          "uid": 21048761,
                          "base": {
                              "name": "随喵的令和",
                              "face": "https://i2.hdslb.com/bfs/face/82d1561e3df7cfc73176d06a5de750bdcc66e89b.jpg",
                              "name_color": 0,
                              "is_mystery": false,
                              "risk_ctrl_info": {
                                  "name": "随喵的令和",
                                  "face": "https://i2.hdslb.com/bfs/face/82d1561e3df7cfc73176d06a5de750bdcc66e89b.jpg"
                              },
                              "origin_info": {
                                  "name": "随喵的令和",
                                  "face": "https://i2.hdslb.com/bfs/face/82d1561e3df7cfc73176d06a5de750bdcc66e89b.jpg"
                              },
                              "official_info": {
                                  "role": 0,
                                  "title": "",
                                  "desc": "",
                                  "type": -1
                              },
                              "name_color_str": ""
                          },
                          "medal": null,
                          "wealth": null,
                          "title": null,
                          "guard": null,
                          "uhead_frame": null,
                          "guard_leader": null
                      }
                  },
                  {
                      "uid": 548191,
                      "face": "https://i1.hdslb.com/bfs/face/081a8d1767bb7aad26472f2686dc6a7401789a68.jpg",
                      "uname": "随一宝宝家的Holan",
                      "score": "179",
                      "rank": 5,
                      "guard_level": 3,
                      "wealth_level": 33,
                      "is_mystery": false,
                      "uinfo": {
                          "uid": 548191,
                          "base": {
                              "name": "随一宝宝家的Holan",
                              "face": "https://i1.hdslb.com/bfs/face/081a8d1767bb7aad26472f2686dc6a7401789a68.jpg",
                              "name_color": 0,
                              "is_mystery": false,
                              "risk_ctrl_info": {
                                  "name": "随一宝宝家的Holan",
                                  "face": "https://i1.hdslb.com/bfs/face/081a8d1767bb7aad26472f2686dc6a7401789a68.jpg"
                              },
                              "origin_info": {
                                  "name": "随一宝宝家的Holan",
                                  "face": "https://i1.hdslb.com/bfs/face/081a8d1767bb7aad26472f2686dc6a7401789a68.jpg"
                              },
                              "official_info": {
                                  "role": 0,
                                  "title": "",
                                  "desc": "",
                                  "type": -1
                              },
                              "name_color_str": ""
                          },
                          "medal": null,
                          "wealth": null,
                          "title": null,
                          "guard": null,
                          "uhead_frame": null,
                          "guard_leader": null
                      }
                  },
                  {
                      "uid": 3756806,
                      "face": "https://i2.hdslb.com/bfs/face/7bee89f83468a5d26c8fbd3bf4071ae5aa0ef412.jpg",
                      "uname": "半句话半个冷笑话",
                      "score": "69",
                      "rank": 6,
                      "guard_level": 3,
                      "wealth_level": 25,
                      "is_mystery": false,
                      "uinfo": {
                          "uid": 3756806,
                          "base": {
                              "name": "半句话半个冷笑话",
                              "face": "https://i2.hdslb.com/bfs/face/7bee89f83468a5d26c8fbd3bf4071ae5aa0ef412.jpg",
                              "name_color": 0,
                              "is_mystery": false,
                              "risk_ctrl_info": {
                                  "name": "半句话半个冷笑话",
                                  "face": "https://i2.hdslb.com/bfs/face/7bee89f83468a5d26c8fbd3bf4071ae5aa0ef412.jpg"
                              },
                              "origin_info": {
                                  "name": "半句话半个冷笑话",
                                  "face": "https://i2.hdslb.com/bfs/face/7bee89f83468a5d26c8fbd3bf4071ae5aa0ef412.jpg"
                              },
                              "official_info": {
                                  "role": 0,
                                  "title": "",
                                  "desc": "",
                                  "type": -1
                              },
                              "name_color_str": ""
                          },
                          "medal": null,
                          "wealth": null,
                          "title": null,
                          "guard": null,
                          "uhead_frame": null,
                          "guard_leader": null
                      }
                  },
                  {
                      "uid": 19522017,
                      "face": "https://i0.hdslb.com/bfs/face/d99bfeabe3ca77594faea29fef3e8dd6d0b5da45.jpg",
                      "uname": "Evariste_鲤家的饲主某E",
                      "score": "49",
                      "rank": 7,
                      "guard_level": 0,
                      "wealth_level": 3,
                      "is_mystery": false,
                      "uinfo": {
                          "uid": 19522017,
                          "base": {
                              "name": "Evariste_鲤家的饲主某E",
                              "face": "https://i0.hdslb.com/bfs/face/d99bfeabe3ca77594faea29fef3e8dd6d0b5da45.jpg",
                              "name_color": 0,
                              "is_mystery": false,
                              "risk_ctrl_info": {
                                  "name": "Evariste_鲤家的饲主某E",
                                  "face": "https://i0.hdslb.com/bfs/face/d99bfeabe3ca77594faea29fef3e8dd6d0b5da45.jpg"
                              },
                              "origin_info": {
                                  "name": "Evariste_鲤家的饲主某E",
                                  "face": "https://i0.hdslb.com/bfs/face/d99bfeabe3ca77594faea29fef3e8dd6d0b5da45.jpg"
                              },
                              "official_info": {
                                  "role": 0,
                                  "title": "",
                                  "desc": "",
                                  "type": -1
                              },
                              "name_color_str": ""
                          },
                          "medal": null,
                          "wealth": null,
                          "title": null,
                          "guard": null,
                          "uhead_frame": null,
                          "guard_leader": null
                      }
                  }
              ],
              "count": 199,
              "count_text": "199"
          },
          "dm_brush_info": {
              "min_time": 700,
              "brush_count": 100,
              "slice_count": 2,
              "storage_time": 3000
          },
          "dm_emoticon_info": {
              "is_open_emoticon": 1,
              "is_shield_emoticon": 0
          },
          "dm_tag_info": {
              "dm_tag": 0,
              "platform": null,
              "extra": "",
              "dm_chronos_extra": "",
              "dm_mode": null,
              "dm_setting_switch": 0,
              "material_conf": null
          },
          "topic_info": {
              "topic_id": 0,
              "topic_name": ""
          },
          "game_info": {
              "game_status": 0
          },
          "watched_show": {
              "switch": true,
              "num": 768,
              "text_small": "768",
              "text_large": "768人看过",
              "icon": "",
              "icon_location": 0,
              "icon_web": ""
          },
          "topic_room_info": {
              "interactive_h5_url": "",
              "watermark": 1
          },
          "show_reserve_status": false,
          "second_create_info": null,
          "play_together_info": null,
          "cloud_game_info": {
              "is_gaming": 0
          },
          "like_info_v3": {
              "total_likes": 1094,
              "click_block": false,
              "count_block": false,
              "guild_emo_text": "试试双击点赞 让主播被更多人看到吧～",
              "guild_dm_text": "点赞30次可以帮主播冲刺热门榜哦～",
              "like_dm_text": "谢谢你的赞，每点赞30次有概率为主播增加曝光哦～",
              "hand_icons": [
                  "https://i0.hdslb.com/bfs/live/1aac78a01b4cba5aa3f7c935c49b0a0bd8992ce1.webp",
                  "https://i0.hdslb.com/bfs/live/81c69596db4a94c53bebed567c1e457c8240e91c.webp",
                  "https://i0.hdslb.com/bfs/live/f281abe0bb186253acfe7b82d54d62fd25f5fee3.webp",
                  "https://i0.hdslb.com/bfs/live/e99441bacfb0a516a08455c5eb9266e7d6174a5c.webp",
                  "https://i0.hdslb.com/bfs/live/2e9145b42b5acbd4a13de1540b42e709832f9bae.webp",
                  "https://i0.hdslb.com/bfs/live/09f36c21175ca217268896305f1d248f3f64add0.webp",
                  "https://i0.hdslb.com/bfs/live/8e90f8aedd66237922958654efb36d97ac4a5701.webp",
                  "https://i0.hdslb.com/bfs/live/29b65444d287680263b1024a6a0d479f7ad41ad7.webp",
                  "https://i0.hdslb.com/bfs/live/b0faa7d645c035cd1d61de5011cdc3880ef232d6.webp",
                  "https://i0.hdslb.com/bfs/live/10fe45806cb0c3e13ef5c9a3ca03a0707667934b.webp"
              ],
              "dm_icons": [
                  "https://i0.hdslb.com/bfs/live/ecc2c5a2efc1c40a1125bb0d648d891e0e8cbd3b.png",
                  "https://i0.hdslb.com/bfs/live/391b45ed00617391b863f6441d5d040165c4d1b4.png",
                  "https://i0.hdslb.com/bfs/live/ad7c7adc9ee0c778cea161dea6c5a96c4ff7f845.png",
                  "https://i0.hdslb.com/bfs/live/c6a4781558d84838cc4963fcf94792e54beabe15.png",
                  "https://i0.hdslb.com/bfs/live/6920410aca21e3397f23ee7ae5ddbf8f1c625943.png",
                  "https://i0.hdslb.com/bfs/live/7e671e053c365f91a0a0f4837ec92ac9581977ec.png"
              ],
              "eggshells_icon": "https://i0.hdslb.com/bfs/live/ea6454847e9b9d1b6be03f04f829620b4f0a5d46.svga",
              "count_show_time": 15,
              "process_icon": "https://i0.hdslb.com/bfs/live/6d96e9de0cc5c80acf3537fe9a23d94da0ab5ff5.png",
              "process_color": "#4DFF6699",
              "report_click_limit": 15,
              "report_time_min": 5,
              "report_time_max": 10,
              "icon": "https://i0.hdslb.com/bfs/live/fe8e848118e8f18eccbd11cad62b29105c51a797.png",
              "cooldown": 0.35,
              "hand_use_face": true,
              "guide_icon_urls": [
                  "https://i0.hdslb.com/bfs/live/2e9145b42b5acbd4a13de1540b42e709832f9bae.webp",
                  "https://i0.hdslb.com/bfs/live/b0faa7d645c035cd1d61de5011cdc3880ef232d6.webp"
              ],
              "guide_icon_ratio": 1.3
          },
          "live_play_info": {
              "show_widget_banner": true,
              "show_left_entry": true,
              "widget_version": 1
          },
          "multi_voice": {
              "switch_status": 2,
              "members": [],
              "mv_role": 0,
              "seat_type": 0,
              "invoking_time": 0,
              "version": 0,
              "pk": null,
              "biz_session_id": "",
              "mode_details": null,
              "hat_list": null,
              "battle_info": null
          },
          "popular_rank_info": {
              "rank": 9,
              "countdown": 3511,
              "timestamp": 1718287290,
              "url": "https://live.bilibili.com/p/html/live-app-hotrank/index.html?is_live_half_webview=1\u0026hybrid_rotate_d=1\u0026hybrid_half_ui=1,3,100p,70p,0,0,30,100,12;2,2,375,100p,0,0,30,100,0;3,3,100p,70p,0,0,30,100,12;4,2,375,100p,0,0,30,100,0;5,3,100p,70p,0,0,30,100,0;6,3,100p,70p,0,0,30,100,0;7,3,100p,70p,0,0,30,100,0;8,3,100p,70p,0,0,30,100,0\u0026pc_ui=338,465,f4eefa,0\u0026redirect=v2\u0026anchorId=1355412269",
              "on_rank_name": "人气",
              "rank_name": "人气榜"
          },
          "new_area_rank_info": {
              "items": [
                  {
                      "conf_id": 11,
                      "rank_name": "虚拟航海",
                      "uid": 1355412269,
                      "rank": 13,
                      "icon_url_blue": "https://i0.hdslb.com/bfs/live/18e2990a546d33368200f9058f3d9dbc4038eb5c.png",
                      "icon_url_pink": "https://i0.hdslb.com/bfs/live/a6c490c36e88c7b191a04883a5ec15aed187a8f7.png",
                      "icon_url_grey": "https://i0.hdslb.com/bfs/live/cb7444b1faf1d785df6265bfdc1fcfc993419b76.png",
                      "jump_url_link": "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=3\u0026ruid=1355412269\u0026conf_id=11\u0026is_live_half_webview=1\u0026hybrid_rotate_d=1\u0026is_cling_player=1\u0026hybrid_half_ui=1,3,100p,70p,f4eefa,0,30,100,0,0;2,2,375,100p,f4eefa,0,30,100,0,0;3,3,100p,70p,f4eefa,0,30,100,0,0;4,2,375,100p,f4eefa,0,30,100,0,0;5,3,100p,70p,f4eefa,0,30,100,0,0;6,3,100p,70p,f4eefa,0,30,100,0,0;7,3,100p,70p,f4eefa,0,30,100,0,0;8,3,100p,70p,f4eefa,0,30,100,0,0#/area-rank",
                      "jump_url_pc": "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=4\u0026ruid=1355412269\u0026conf_id=11\u0026pc_ui=338,465,f4eefa,0#/area-rank",
                      "jump_url_pink": "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=1\u0026ruid=1355412269\u0026conf_id=11\u0026is_live_half_webview=1\u0026hybrid_rotate_d=1\u0026hybrid_half_ui=1,3,100p,70p,ffffff,0,30,100,12,0;2,2,375,100p,ffffff,0,30,100,0,0;3,3,100p,70p,ffffff,0,30,100,12,0;4,2,375,100p,ffffff,0,30,100,0,0;5,3,100p,70p,ffffff,0,30,100,0,0;6,3,100p,70p,ffffff,0,30,100,0,0;7,3,100p,70p,ffffff,0,30,100,0,0;8,3,100p,70p,ffffff,0,30,100,0,0#/area-rank",
                      "jump_url_web": "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=2\u0026ruid=1355412269\u0026conf_id=11#/area-rank"
                  }
              ],
              "rotation_cycle_time_web": 10000
          },
          "gift_star": {
              "show": true,
              "display_widget_ab_group": 0
          },
          "progress_for_widget": {
              "gift_star_process": {
                  "task_info": {
                      "start_date": 20240610,
                      "process_list": [
                          {
                              "gift_id": 31049,
                              "gift_img": "https://s1.hdslb.com/bfs/live/96ec38f351a4e190c4a525bc5e11ff09d2874064.png",
                              "gift_name": "礼物星球",
                              "completed_num": 1,
                              "target_num": 10
                          },
                          {
                              "gift_id": 31588,
                              "gift_img": "https://s1.hdslb.com/bfs/live/311930350df3b8e467d13b992e62344ca1e3664f.png",
                              "gift_name": "礼物星球",
                              "completed_num": 0,
                              "target_num": 5
                          },
                          {
                              "gift_id": 31053,
                              "gift_img": "https://s1.hdslb.com/bfs/live/0205d703553d0625ba24aefd421eb5e350ce3c20.png",
                              "gift_name": "礼物星球",
                              "completed_num": 0,
                              "target_num": 5
                          }
                      ],
                      "finished": false,
                      "ddl_timestamp": 1718553600,
                      "version": 1718284261687,
                      "reward_gift": 0,
                      "reward_gift_img": "https://i0.hdslb.com/bfs/live/52edb4ab7377ece34ac15b21154d13d188874b01.png",
                      "reward_gift_name": "礼物星球",
                      "level_info": {
                          "star_name": "礼物星球",
                          "level_tip": "暂未达成",
                          "level_img": "https://i0.hdslb.com/bfs/live/91292553d02aa10a133cdf6d8e579d8588e2067a.png",
                          "level_id": 0
                      }
                  },
                  "preload_timestamp": 1718552858,
                  "preload": false,
                  "preload_task_info": null,
                  "widget_bg": "",
                  "jump_schema": "https://live.bilibili.com/p/html/live-app-gift-planet/index.html?roomId=23075997\u0026ruid=1355412269\u0026app_common=open\u0026plt=\u0026pc_ui=375,580,2E345A,0\u0026is_live_half_webview=1\u0026hybrid_half_ui=1,3,100p,70p,0,0,30,100,12,0;2,2,375,100p,0,0,30,100,0,0;3,3,100p,70p,0,0,30,100,12,0;4,2,375,100p,0,0,30,100,0,0;5,3,100p,70p,0,0,30,100,12,0;6,3,100p,70p,0,0,30,100,12,0;7,3,100p,70p,0,0,30,100,12,0;8,2,320,480,0,0,30,100,0,0\u0026display_widget_ab_group=0",
                  "ab_group": 0
              },
              "wish_process": null,
              "star_knight": null,
              "collection_praise_process": {
                  "id": 0,
                  "uid": 0,
                  "target_praise": 0,
                  "current_praise": 0,
                  "start_time": 0,
                  "end_time": 0,
                  "benefit": "",
                  "isSuccess": false,
                  "exist": false,
                  "audit_status": 0,
                  "jump_url": "",
                  "current_praise_text": "",
                  "icon_url": "https://i0.hdslb.com/bfs/live/fb44b6cc32793f15eb588ca93d2d3fc522c31dfd.png",
                  "live_id": ""
              }
          },
          "revenue_demotion": {
              "global_gift_config_demotion": false
          },
          "revenue_material_md5": null,
          "block_info": {
              "block": false
          },
          "danmu_extra": {
              "screen_switch_off": false
          },
          "video_connection_info": null,
          "player_throttle_info": {
              "status": 0,
              "normal_sleep_time": 0,
              "fullscreen_sleep_time": 0,
              "tab_sleep_time": 0,
              "prompt_time": 0
          },
          "guard_info": {
              "count": 242,
              "anchor_guard_achieve_level": 100
          },
          "hot_rank_info": null,
          "room_rank_info": {
              "anchor_rank_entry": null,
              "user_rank_entry": {
                  "user_contribution_rank_entry": {
                      "item": [
                          {
                              "uid": 405771212,
                              "name": "---愤怒的小鸟---",
                              "face": "https://i0.hdslb.com/bfs/face/member/noface.jpg",
                              "rank": 1,
                              "score": 477,
                              "medal_info": {
                                  "guard_level": 3,
                                  "medal_color_start": 398668,
                                  "medal_color_end": 6850801,
                                  "medal_color_border": 6809855,
                                  "medal_name": "随随冰",
                                  "level": 26,
                                  "target_id": 1355412269,
                                  "is_light": 1
                              },
                              "guard_level": 0,
                              "wealth_level": 27,
                              "is_mystery": false,
                              "uinfo": {
                                  "uid": 405771212,
                                  "base": {
                                      "name": "---愤怒的小鸟---",
                                      "face": "https://i0.hdslb.com/bfs/face/member/noface.jpg",
                                      "name_color": 0,
                                      "is_mystery": false,
                                      "risk_ctrl_info": {
                                          "name": "---愤怒的小鸟---",
                                          "face": "https://i0.hdslb.com/bfs/face/member/noface.jpg"
                                      },
                                      "origin_info": {
                                          "name": "---愤怒的小鸟---",
                                          "face": "https://i0.hdslb.com/bfs/face/member/noface.jpg"
                                      },
                                      "official_info": {
                                          "role": 0,
                                          "title": "",
                                          "desc": "",
                                          "type": -1
                                      },
                                      "name_color_str": ""
                                  },
                                  "medal": {
                                      "name": "随随冰",
                                      "level": 26,
                                      "color_start": 398668,
                                      "color_end": 6850801,
                                      "color_border": 6809855,
                                      "color": 398668,
                                      "id": 0,
                                      "typ": 0,
                                      "is_light": 1,
                                      "ruid": 1355412269,
                                      "guard_level": 3,
                                      "score": 50041942,
                                      "guard_icon": "https://i0.hdslb.com/bfs/live/143f5ec3003b4080d1b5f817a9efdca46d631945.png",
                                      "honor_icon": ""
                                  },
                                  "wealth": {
                                      "level": 27,
                                      "dm_icon_key": ""
                                  },
                                  "title": {
                                      "old_title_css_id": "",
                                      "title_css_id": ""
                                  },
                                  "guard": null,
                                  "uhead_frame": null,
                                  "guard_leader": null
                              }
                          },
                          {
                              "uid": 287837821,
                              "name": "境花与蝶",
                              "face": "https://i1.hdslb.com/bfs/face/bb26f0ed92a1ca9f7836de95f74c4a4945a272a3.jpg",
                              "rank": 2,
                              "score": 405,
                              "medal_info": null,
                              "guard_level": 0,
                              "wealth_level": 41,
                              "is_mystery": false,
                              "uinfo": {
                                  "uid": 287837821,
                                  "base": {
                                      "name": "境花与蝶",
                                      "face": "https://i1.hdslb.com/bfs/face/bb26f0ed92a1ca9f7836de95f74c4a4945a272a3.jpg",
                                      "name_color": 0,
                                      "is_mystery": false,
                                      "risk_ctrl_info": {
                                          "name": "境花与蝶",
                                          "face": "https://i1.hdslb.com/bfs/face/bb26f0ed92a1ca9f7836de95f74c4a4945a272a3.jpg"
                                      },
                                      "origin_info": {
                                          "name": "境花与蝶",
                                          "face": "https://i1.hdslb.com/bfs/face/bb26f0ed92a1ca9f7836de95f74c4a4945a272a3.jpg"
                                      },
                                      "official_info": {
                                          "role": 0,
                                          "title": "",
                                          "desc": "",
                                          "type": -1
                                      },
                                      "name_color_str": ""
                                  },
                                  "medal": null,
                                  "wealth": {
                                      "level": 41,
                                      "dm_icon_key": "ChronosWealth_4.png"
                                  },
                                  "title": {
                                      "old_title_css_id": "",
                                      "title_css_id": ""
                                  },
                                  "guard": null,
                                  "uhead_frame": null,
                                  "guard_leader": null
                              }
                          },
                          {
                              "uid": 156077349,
                              "name": "雨齐青",
                              "face": "https://i0.hdslb.com/bfs/face/32fdee8a66a61d38b6e65d3310363c6af46480a2.jpg",
                              "rank": 3,
                              "score": 282,
                              "medal_info": {
                                  "guard_level": 2,
                                  "medal_color_start": 2951253,
                                  "medal_color_end": 10329087,
                                  "medal_color_border": 16771156,
                                  "medal_name": "随随冰",
                                  "level": 30,
                                  "target_id": 1355412269,
                                  "is_light": 1
                              },
                              "guard_level": 0,
                              "wealth_level": 42,
                              "is_mystery": false,
                              "uinfo": {
                                  "uid": 156077349,
                                  "base": {
                                      "name": "雨齐青",
                                      "face": "https://i0.hdslb.com/bfs/face/32fdee8a66a61d38b6e65d3310363c6af46480a2.jpg",
                                      "name_color": 0,
                                      "is_mystery": false,
                                      "risk_ctrl_info": {
                                          "name": "雨齐青",
                                          "face": "https://i0.hdslb.com/bfs/face/32fdee8a66a61d38b6e65d3310363c6af46480a2.jpg"
                                      },
                                      "origin_info": {
                                          "name": "雨齐青",
                                          "face": "https://i0.hdslb.com/bfs/face/32fdee8a66a61d38b6e65d3310363c6af46480a2.jpg"
                                      },
                                      "official_info": {
                                          "role": 0,
                                          "title": "",
                                          "desc": "",
                                          "type": -1
                                      },
                                      "name_color_str": ""
                                  },
                                  "medal": {
                                      "name": "随随冰",
                                      "level": 30,
                                      "color_start": 2951253,
                                      "color_end": 10329087,
                                      "color_border": 16771156,
                                      "color": 2951253,
                                      "id": 0,
                                      "typ": 0,
                                      "is_light": 1,
                                      "ruid": 1355412269,
                                      "guard_level": 2,
                                      "score": 51014254,
                                      "guard_icon": "https://i0.hdslb.com/bfs/live/98a201c14a64e860a758f089144dcf3f42e7038c.png",
                                      "honor_icon": ""
                                  },
                                  "wealth": {
                                      "level": 42,
                                      "dm_icon_key": "ChronosWealth_4.png"
                                  },
                                  "title": {
                                      "old_title_css_id": "",
                                      "title_css_id": ""
                                  },
                                  "guard": null,
                                  "uhead_frame": null,
                                  "guard_leader": null
                              }
                          },
                          {
                              "uid": 21048761,
                              "name": "随喵的令和",
                              "face": "https://i2.hdslb.com/bfs/face/82d1561e3df7cfc73176d06a5de750bdcc66e89b.jpg",
                              "rank": 4,
                              "score": 181,
                              "medal_info": null,
                              "guard_level": 0,
                              "wealth_level": 38,
                              "is_mystery": false,
                              "uinfo": {
                                  "uid": 21048761,
                                  "base": {
                                      "name": "随喵的令和",
                                      "face": "https://i2.hdslb.com/bfs/face/82d1561e3df7cfc73176d06a5de750bdcc66e89b.jpg",
                                      "name_color": 0,
                                      "is_mystery": false,
                                      "risk_ctrl_info": {
                                          "name": "随喵的令和",
                                          "face": "https://i2.hdslb.com/bfs/face/82d1561e3df7cfc73176d06a5de750bdcc66e89b.jpg"
                                      },
                                      "origin_info": {
                                          "name": "随喵的令和",
                                          "face": "https://i2.hdslb.com/bfs/face/82d1561e3df7cfc73176d06a5de750bdcc66e89b.jpg"
                                      },
                                      "official_info": {
                                          "role": 0,
                                          "title": "",
                                          "desc": "",
                                          "type": -1
                                      },
                                      "name_color_str": ""
                                  },
                                  "medal": null,
                                  "wealth": {
                                      "level": 38,
                                      "dm_icon_key": ""
                                  },
                                  "title": {
                                      "old_title_css_id": "",
                                      "title_css_id": ""
                                  },
                                  "guard": null,
                                  "uhead_frame": null,
                                  "guard_leader": null
                              }
                          },
                          {
                              "uid": 548191,
                              "name": "随一宝宝家的Holan",
                              "face": "https://i1.hdslb.com/bfs/face/081a8d1767bb7aad26472f2686dc6a7401789a68.jpg",
                              "rank": 5,
                              "score": 179,
                              "medal_info": {
                                  "guard_level": 3,
                                  "medal_color_start": 398668,
                                  "medal_color_end": 6850801,
                                  "medal_color_border": 6809855,
                                  "medal_name": "随随冰",
                                  "level": 27,
                                  "target_id": 1355412269,
                                  "is_light": 1
                              },
                              "guard_level": 0,
                              "wealth_level": 33,
                              "is_mystery": false,
                              "uinfo": {
                                  "uid": 548191,
                                  "base": {
                                      "name": "随一宝宝家的Holan",
                                      "face": "https://i1.hdslb.com/bfs/face/081a8d1767bb7aad26472f2686dc6a7401789a68.jpg",
                                      "name_color": 0,
                                      "is_mystery": false,
                                      "risk_ctrl_info": {
                                          "name": "随一宝宝家的Holan",
                                          "face": "https://i1.hdslb.com/bfs/face/081a8d1767bb7aad26472f2686dc6a7401789a68.jpg"
                                      },
                                      "origin_info": {
                                          "name": "随一宝宝家的Holan",
                                          "face": "https://i1.hdslb.com/bfs/face/081a8d1767bb7aad26472f2686dc6a7401789a68.jpg"
                                      },
                                      "official_info": {
                                          "role": 0,
                                          "title": "",
                                          "desc": "",
                                          "type": -1
                                      },
                                      "name_color_str": ""
                                  },
                                  "medal": {
                                      "name": "随随冰",
                                      "level": 27,
                                      "color_start": 398668,
                                      "color_end": 6850801,
                                      "color_border": 6809855,
                                      "color": 398668,
                                      "id": 0,
                                      "typ": 0,
                                      "is_light": 1,
                                      "ruid": 1355412269,
                                      "guard_level": 3,
                                      "score": 50128470,
                                      "guard_icon": "https://i0.hdslb.com/bfs/live/143f5ec3003b4080d1b5f817a9efdca46d631945.png",
                                      "honor_icon": ""
                                  },
                                  "wealth": {
                                      "level": 33,
                                      "dm_icon_key": ""
                                  },
                                  "title": {
                                      "old_title_css_id": "",
                                      "title_css_id": ""
                                  },
                                  "guard": null,
                                  "uhead_frame": null,
                                  "guard_leader": null
                              }
                          },
                          {
                              "uid": 3756806,
                              "name": "半句话半个冷笑话",
                              "face": "https://i2.hdslb.com/bfs/face/7bee89f83468a5d26c8fbd3bf4071ae5aa0ef412.jpg",
                              "rank": 6,
                              "score": 69,
                              "medal_info": {
                                  "guard_level": 3,
                                  "medal_color_start": 1725515,
                                  "medal_color_end": 5414290,
                                  "medal_color_border": 6809855,
                                  "medal_name": "随随冰",
                                  "level": 21,
                                  "target_id": 1355412269,
                                  "is_light": 1
                              },
                              "guard_level": 0,
                              "wealth_level": 25,
                              "is_mystery": false,
                              "uinfo": {
                                  "uid": 3756806,
                                  "base": {
                                      "name": "半句话半个冷笑话",
                                      "face": "https://i2.hdslb.com/bfs/face/7bee89f83468a5d26c8fbd3bf4071ae5aa0ef412.jpg",
                                      "name_color": 0,
                                      "is_mystery": false,
                                      "risk_ctrl_info": {
                                          "name": "半句话半个冷笑话",
                                          "face": "https://i2.hdslb.com/bfs/face/7bee89f83468a5d26c8fbd3bf4071ae5aa0ef412.jpg"
                                      },
                                      "origin_info": {
                                          "name": "半句话半个冷笑话",
                                          "face": "https://i2.hdslb.com/bfs/face/7bee89f83468a5d26c8fbd3bf4071ae5aa0ef412.jpg"
                                      },
                                      "official_info": {
                                          "role": 0,
                                          "title": "",
                                          "desc": "",
                                          "type": -1
                                      },
                                      "name_color_str": ""
                                  },
                                  "medal": {
                                      "name": "随随冰",
                                      "level": 21,
                                      "color_start": 1725515,
                                      "color_end": 5414290,
                                      "color_border": 6809855,
                                      "color": 1725515,
                                      "id": 0,
                                      "typ": 0,
                                      "is_light": 1,
                                      "ruid": 1355412269,
                                      "guard_level": 3,
                                      "score": 50001604,
                                      "guard_icon": "https://i0.hdslb.com/bfs/live/143f5ec3003b4080d1b5f817a9efdca46d631945.png",
                                      "honor_icon": ""
                                  },
                                  "wealth": {
                                      "level": 25,
                                      "dm_icon_key": ""
                                  },
                                  "title": {
                                      "old_title_css_id": "",
                                      "title_css_id": ""
                                  },
                                  "guard": null,
                                  "uhead_frame": null,
                                  "guard_leader": null
                              }
                          },
                          {
                              "uid": 19522017,
                              "name": "Evariste_鲤家的饲主某E",
                              "face": "https://i0.hdslb.com/bfs/face/d99bfeabe3ca77594faea29fef3e8dd6d0b5da45.jpg",
                              "rank": 7,
                              "score": 49,
                              "medal_info": {
                                  "guard_level": 0,
                                  "medal_color_start": 9272486,
                                  "medal_color_end": 9272486,
                                  "medal_color_border": 9272486,
                                  "medal_name": "随随冰",
                                  "level": 11,
                                  "target_id": 1355412269,
                                  "is_light": 1
                              },
                              "guard_level": 0,
                              "wealth_level": 3,
                              "is_mystery": false,
                              "uinfo": {
                                  "uid": 19522017,
                                  "base": {
                                      "name": "Evariste_鲤家的饲主某E",
                                      "face": "https://i0.hdslb.com/bfs/face/d99bfeabe3ca77594faea29fef3e8dd6d0b5da45.jpg",
                                      "name_color": 0,
                                      "is_mystery": false,
                                      "risk_ctrl_info": {
                                          "name": "Evariste_鲤家的饲主某E",
                                          "face": "https://i0.hdslb.com/bfs/face/d99bfeabe3ca77594faea29fef3e8dd6d0b5da45.jpg"
                                      },
                                      "origin_info": {
                                          "name": "Evariste_鲤家的饲主某E",
                                          "face": "https://i0.hdslb.com/bfs/face/d99bfeabe3ca77594faea29fef3e8dd6d0b5da45.jpg"
                                      },
                                      "official_info": {
                                          "role": 0,
                                          "title": "",
                                          "desc": "",
                                          "type": -1
                                      },
                                      "name_color_str": ""
                                  },
                                  "medal": {
                                      "name": "随随冰",
                                      "level": 11,
                                      "color_start": 9272486,
                                      "color_end": 9272486,
                                      "color_border": 9272486,
                                      "color": 9272486,
                                      "id": 0,
                                      "typ": 0,
                                      "is_light": 1,
                                      "ruid": 1355412269,
                                      "guard_level": 0,
                                      "score": 17550,
                                      "guard_icon": "",
                                      "honor_icon": ""
                                  },
                                  "wealth": {
                                      "level": 3,
                                      "dm_icon_key": ""
                                  },
                                  "title": {
                                      "old_title_css_id": "",
                                      "title_css_id": ""
                                  },
                                  "guard": null,
                                  "uhead_frame": null,
                                  "guard_leader": null
                              }
                          }
                      ],
                      "count": 256,
                      "show_max": 7,
                      "count_text": "256"
                  }
              },
              "user_rank_tab_list": {
                  "tab": [
                      {
                          "type": "contribution_tab",
                          "title": "高能用户",
                          "status": 1,
                          "default": 1,
                          "comment": "",
                          "desc_url": "",
                          "switch": null,
                          "sub_tab": [
                              {
                                  "type": "online_rank",
                                  "title": "在线榜",
                                  "status": 1,
                                  "default": 1,
                                  "comment": "",
                                  "desc_url": "https://live.bilibili.com/p/contribute-rank-description-h5/index.html?is_live_half_webview=1\u0026hybrid_half_ui=1,3,100p,70p,0,0,30,100,12,0;2,2,375,100p,0,0,30,100,0,0;3,3,100p,70p,0,0,30,100,12,0;4,2,375,100p,0,0,30,100,0,0;5,3,100p,70p,0,0,30,100,12,0;6,3,100p,70p,0,0,30,100,12,0;7,3,100p,70p,0,0,30,100,12,0;8,3,100p,70p,0,0,30,100,0,0\u0026rank_type=1\u0026test_type=b",
                                  "switch": [
                                      {
                                          "text": "贡献值",
                                          "switch": "contribution_rank",
                                          "ui_type": {
                                              "op_button_text": 1,
                                              "rank_prefix": 1,
                                              "refresh_entry": 1,
                                              "show_score": 1
                                          },
                                          "comment": "投喂、发弹幕均可获得贡献值"
                                      },
                                      {
                                          "text": "进房时间",
                                          "switch": "entry_time_rank",
                                          "ui_type": {
                                              "op_button_text": 2,
                                              "rank_prefix": 0,
                                              "show_score": 0
                                          },
                                          "comment": "用户进房顺序排序"
                                      }
                                  ],
                                  "sub_tab": null
                              },
                              {
                                  "type": "daily_rank",
                                  "title": "日榜",
                                  "status": 1,
                                  "default": 0,
                                  "comment": "",
                                  "desc_url": "https://live.bilibili.com/p/contribute-rank-description-h5/index.html?is_live_half_webview=1\u0026hybrid_half_ui=1,3,100p,70p,0,0,30,100,12,0;2,2,375,100p,0,0,30,100,0,0;3,3,100p,70p,0,0,30,100,12,0;4,2,375,100p,0,0,30,100,0,0;5,3,100p,70p,0,0,30,100,12,0;6,3,100p,70p,0,0,30,100,12,0;7,3,100p,70p,0,0,30,100,12,0;8,3,100p,70p,0,0,30,100,0,0\u0026rank_type=2\u0026test_type=b",
                                  "switch": [
                                      {
                                          "text": "当日",
                                          "switch": "today_rank",
                                          "ui_type": {
                                              "op_button_text": 1,
                                              "rank_prefix": 0,
                                              "show_score": 0
                                          },
                                          "comment": "用户当日贡献值排序"
                                      },
                                      {
                                          "text": "昨日",
                                          "switch": "yesterday_rank",
                                          "ui_type": {
                                              "op_button_text": 0,
                                              "rank_prefix": 0,
                                              "show_score": 0
                                          },
                                          "comment": "昨日贡献值前100排名"
                                      }
                                  ],
                                  "sub_tab": null
                              },
                              {
                                  "type": "weekly_rank",
                                  "title": "周榜",
                                  "status": 1,
                                  "default": 0,
                                  "comment": "",
                                  "desc_url": "https://live.bilibili.com/p/contribute-rank-description-h5/index.html?is_live_half_webview=1\u0026hybrid_half_ui=1,3,100p,70p,0,0,30,100,12,0;2,2,375,100p,0,0,30,100,0,0;3,3,100p,70p,0,0,30,100,12,0;4,2,375,100p,0,0,30,100,0,0;5,3,100p,70p,0,0,30,100,12,0;6,3,100p,70p,0,0,30,100,12,0;7,3,100p,70p,0,0,30,100,12,0;8,3,100p,70p,0,0,30,100,0,0\u0026rank_type=2\u0026test_type=b",
                                  "switch": [
                                      {
                                          "text": "本周",
                                          "switch": "current_week_rank",
                                          "ui_type": {
                                              "op_button_text": 1,
                                              "rank_prefix": 0,
                                              "show_score": 0
                                          },
                                          "comment": "用户当周贡献值排序"
                                      },
                                      {
                                          "text": "上周",
                                          "switch": "last_week_rank",
                                          "ui_type": {
                                              "op_button_text": 0,
                                              "rank_prefix": 0,
                                              "show_score": 0
                                          },
                                          "comment": "上周贡献值前100排名"
                                      }
                                  ],
                                  "sub_tab": null
                              },
                              {
                                  "type": "monthly_rank",
                                  "title": "月榜",
                                  "status": 1,
                                  "default": 0,
                                  "comment": "",
                                  "desc_url": "https://live.bilibili.com/p/contribute-rank-description-h5/index.html?is_live_half_webview=1\u0026hybrid_half_ui=1,3,100p,70p,0,0,30,100,12,0;2,2,375,100p,0,0,30,100,0,0;3,3,100p,70p,0,0,30,100,12,0;4,2,375,100p,0,0,30,100,0,0;5,3,100p,70p,0,0,30,100,12,0;6,3,100p,70p,0,0,30,100,12,0;7,3,100p,70p,0,0,30,100,12,0;8,3,100p,70p,0,0,30,100,0,0\u0026rank_type=2\u0026test_type=b",
                                  "switch": [
                                      {
                                          "text": "本月",
                                          "switch": "current_month_rank",
                                          "ui_type": {
                                              "op_button_text": 1,
                                              "rank_prefix": 0,
                                              "show_score": 0
                                          },
                                          "comment": "用户当月贡献值排序"
                                      },
                                      {
                                          "text": "上月",
                                          "switch": "last_month_rank",
                                          "ui_type": {
                                              "op_button_text": 0,
                                              "rank_prefix": 0,
                                              "show_score": 0
                                          },
                                          "comment": "上月贡献值前100排名"
                                      }
                                  ],
                                  "sub_tab": null
                              }
                          ]
                      },
                      {
                          "type": "guard_tab",
                          "title": "大航海",
                          "status": 1,
                          "default": 0,
                          "comment": "头号粉丝大航海，上船后可上榜",
                          "desc_url": "",
                          "switch": null,
                          "sub_tab": null
                      }
                  ]
              }
          },
          "dm_reply": {
              "show_reply": true
          },
          "dm_combo": null,
          "dm_vote": null,
          "location": null,
          "interactive_game_tag": {
              "action": 0,
              "game_id": "",
              "game_name": ""
          },
          "video_enhancement": {
              "title": "画质增强",
              "desc": "仅针对超清及以下生效",
              "default_switch_status": 2,
              "highest_quality": 250,
              "is_enabled": true
          },
          "guard_leader": {
              "uid": 0,
              "name": "",
              "face": "",
              "jump_url": "https://live.bilibili.com/p/html/live-app-guard-pilot/index.html?is_live_half_webview=1\u0026hybrid_half_ui=1,3,100p,73p,0,0,30,0,12,0;2,2,375,100p,0,0,30,0,0,0;3,3,100p,73p,0,0,30,0,12,0;4,2,375,100p,0,0,30,0,0,0;5,3,100p,73p,0,0,30,0,12,0;6,3,100p,73p,0,0,30,0,12,0;7,3,100p,73p,0,0,30,0,12,0;8,2,320,480,0,0,30,0,0,0\u0026anchorId=1355412269\u0026roomId=23075997",
              "text": "成为舰队指挥官，解锁房间冠名权益",
              "rank_top_icon1": "https://i0.hdslb.com/bfs/live/64b22e65979b32f7e4e8bec1edb38c697fb320fc.png",
              "rank_top_icon2": "https://i0.hdslb.com/bfs/live/7b9d773c6018ffac9f0eadd3c92f0090e09055f2.png",
              "rank_top_background_url1": "https://i0.hdslb.com/bfs/live/b2832de9cca6a0b3b4872c8d96c05ae713bc51d2.png",
              "rank_top_background_url2": "https://i0.hdslb.com/bfs/live/71397554da8a7bd2ac14905c69891df54ad62ede.png",
              "background_url": "",
              "anchor_background_url": "",
              "input_background_url": "",
              "newly": 0,
              "entry_effect_id": 0,
              "show": 1,
              "rank_top_background_light_url1": "",
              "rank_top_background_light_url2": "https://i0.hdslb.com/bfs/live/c0fbd28b6ddf170b8db2e2c7163eb9d66f66fd8b.png",
              "display_src": "https://i0.hdslb.com/bfs/live/7aed22c78a2a41a5c1b964f1f2a3220c52c1663e.png",
              "avatar_src": "https://i0.hdslb.com/bfs/live/4d1f0d9a39e368c4b9b4128f58f945099a295c39.png",
              "icon_src": "https://i0.hdslb.com/bfs/live/81e9d659b251a74277fe9b1a6f7876c4d00a4950.png"
          },
          "room_anonymous": {
              "open_anonymous": false
          },
          "tab_switches": {
              "subtitle": 0,
              "realtime_data": null
          },
          "universal_interact_info": null,
          "pk_info_v2": null,
          "area_mask_info": {
              "area_masks": {
                  "horizontal_masks": null,
                  "vertical_masks": null
              }
          },
          "xtemplate_config": {
              "dm_brush_info": {
                  "landScape": null,
                  "verticalscreen": {
                      "min_time": 500,
                      "brush_count": 50,
                      "slice_count": 2,
                      "storage_time": 500,
                      "is_hide_anti_brush": 0
                  }
              },
              "dm_speed_info": {
                  "landScape": null,
                  "verticalscreen": {
                      "valley": {
                          "consumetime": 500,
                          "consumecount": 2,
                          "animationtime": 0
                      },
                      "peak": {
                          "consumetime": 600,
                          "consumecount": 3,
                          "animationtime": 0
                      },
                      "proportion": 30,
                      "interval": 10000
                  }
              },
              "dm_pool_info": {
                  "landScape": null,
                  "verticalscreen": {
                      "master_ceiling": 100,
                      "master_count": 1,
                      "guest_config": [
                          {
                              "score_floor": 80,
                              "score_ceiling": 999999,
                              "dm_max": 150,
                              "consume": 80
                          },
                          {
                              "score_floor": 32,
                              "score_ceiling": 79,
                              "dm_max": 100,
                              "consume": 10
                          },
                          {
                              "score_floor": 0,
                              "score_ceiling": 31,
                              "dm_max": 100,
                              "consume": 10
                          }
                      ],
                      "timeout": 300000,
                      "unusual_score": 50
                  }
              }
          },
          "dm_activity": {
              "activity_list": null,
              "ts": 1718287290
          },
          "dm_interaction_ab": {
              "102": 0,
              "103": 0,
              "104": 0,
              "105": 0,
              "106": 0
          },
          "guard_intimacy_rank_status": {
              "guard_rank_new_ab": 0,
              "guard_rank_new_total_status": 0,
              "guard_rank_new_month_status": 0,
              "guard_rank_new_week_status": 0
          }
      }
  }
  ```
- **直播间状态：PREPARE**

  ```json
  {
      "code": 0,
      "message": "0",
      "ttl": 1,
      "data": {
          "room_info": {
              "uid": 1355412269,
              "room_id": 23075997,
              "short_id": 0,
              "title": "该杂交了",
              "cover": "http://i0.hdslb.com/bfs/live/new_room_cover/fb038f94f5f8a24d0bff4c17de0409445a3d8069.jpg",
              "tags": "",
              "background": "http://i0.hdslb.com/bfs/live/room_bg/d62474f3335ebe3c6543977d4c8aba5078cfdb41.jpg",
              "description": "✿能想起来的部分歌单：http://suiii.minamini.cn/",
              "live_status": 2,
              "live_start_time": 0,
              "live_screen_type": 0,
              "lock_status": 0,
              "lock_time": 0,
              "hidden_status": 0,
              "hidden_time": 0,
              "area_id": 745,
              "area_name": "虚拟Gamer",
              "parent_area_id": 9,
              "parent_area_name": "虚拟主播",
              "keyframe": "http://i0.hdslb.com/bfs/live-key-frame/keyframe06122300000023075997imgpgp.jpg",
              "special_type": 0,
              "up_session": "0",
              "pk_status": 0,
              "is_studio": false,
              "pendants": {
                  "frame": {
                      "name": "星辉同轨",
                      "value": "https://i0.hdslb.com/bfs/live/a34b4c869014318859346312c5c6a6ddcb55d1f3.png",
                      "desc": ""
                  }
              },
              "on_voice_join": 0,
              "online": 0,
              "room_type": {
                  "3-21": 0,
                  "3-50": 1
              },
              "sub_session_key": "",
              "live_id": 0,
              "live_id_str": "0",
              "official_room_id": 0,
              "official_room_info": null,
              "voice_background": ""
          },
          "anchor_info": {
              "base_info": {
                  "uname": "随一Suiii",
                  "face": "https://i2.hdslb.com/bfs/face/131b8eb3bf0c347df5bb078a49a35f6f669cb93f.jpg",
                  "gender": "保密",
                  "official_info": {
                      "role": 0,
                      "title": "bilibili直播高能主播",
                      "desc": "",
                      "is_nft": 0,
                      "nft_dmark": "https://i0.hdslb.com/bfs/live/9f176ff49d28c50e9c53ec1c3297bd1ee539b3d6.gif"
                  }
              },
              "live_info": {
                  "level": 32,
                  "level_color": 16746162,
                  "score": 21219165,
                  "upgrade_score": 6094645,
                  "current": [
                      5000000,
                      20613810
                  ],
                  "next": [
                      6700000,
                      27313810
                  ],
                  "rank": "2082"
              },
              "relation_info": {
                  "attention": 28317
              },
              "medal_info": {
                  "medal_name": "随随冰",
                  "medal_id": 448828,
                  "fansclub": 869
              },
              "gift_info": {
                  "price": 0,
                  "price_update_time": 0
              }
          },
          "news_info": {
              "uid": 1355412269,
              "ctime": "2023-02-19 11:59:46",
              "content": "午播恢复中/晚上八点半唱歌杂谈为主，偶尔游戏"
          },
          "rankdb_info": {
              "roomid": 23075997,
              "rank_desc": "",
              "color": "",
              "h5_url": "",
              "web_url": "",
              "timestamp": 1718211564
          },
          "area_rank_info": {
              "areaRank": {
                  "index": 0,
                  "rank": "\u003e1000"
              },
              "liveRank": {
                  "rank": "2082"
              }
          },
          "battle_rank_entry_info": null,
          "tab_info": {
              "list": [
                  {
                      "type": "seven-rank",
                      "desc": "高能用户",
                      "isFirst": 1,
                      "isEvent": 0,
                      "eventType": "",
                      "listType": "",
                      "apiPrefix": "",
                      "rank_name": "room_7day"
                  },
                  {
                      "type": "guard",
                      "desc": "大航海",
                      "isFirst": 0,
                      "isEvent": 0,
                      "eventType": "",
                      "listType": "top-list",
                      "apiPrefix": "",
                      "rank_name": ""
                  }
              ]
          },
          "activity_init_info": {
              "eventList": [],
              "weekInfo": {
                  "bannerInfo": null,
                  "giftName": null
              },
              "giftName": null,
              "lego": {
                  "timestamp": 1718211564,
                  "config": "[{\"name\":\"frame-mng\",\"url\":\"https:\\/\\/live.bilibili.com\\/p\\/html\\/live-web-mng\\/index.html?roomid=#roomid#\u0026arae_id=#area_id#\u0026parent_area_id=#parent_area_id#\u0026ruid=#ruid#\",\"startTime\":1559544736,\"endTime\":1877167950,\"type\":\"frame-mng\"},{\"name\":\"s10-fun\",\"target\":\"sidebar\",\"icon\":\"https:\\/\\/i0.hdslb.com\\/bfs\\/activity-plat\\/static\\/20200908\\/3435f7521efc759ae1f90eae5629a8f0\\/HpxrZ7SOT.png\",\"text\":\"\\u7545\\u73a9s10\",\"url\":\"https:\\/\\/live.bilibili.com\\/s10\\/fun\\/index.html?room_id=#roomid#\u0026width=376\u0026height=600\u0026source=sidebar\",\"color\":\"#2e6fc0\",\"startTime\":1600920000,\"endTime\":1604721600,\"parentAreaId\":2,\"areaId\":86},{\"name\":\"genshin-avatar\",\"target\":\"sidebar\",\"icon\":\"https:\\/\\/i0.hdslb.com\\/bfs\\/activity-plat\\/static\\/20210721\\/fa538c98e9e32dc98919db4f2527ad02\\/qWxN1d0ACu.jpg\",\"text\":\"\\u539f\\u77f3\\u798f\\u5229\",\"url\":\"https:\\/\\/live.bilibili.com\\/activity\\/live-activity-full\\/genshin_avatar\\/mobile.html?no-jump=1\u0026room_id=#roomid#\u0026width=376\u0026height=550#\\/\",\"color\":\"#2e6fc0\",\"frameAllowNoBg\":\"1\",\"frameAllowDrag\":\"1\",\"startTime\":1627012800,\"endTime\":1630425540,\"parentAreaId\":3,\"areaId\":321}]"
              }
          },
          "voice_join_info": {
              "status": {
                  "open": 0,
                  "anchor_open": 0,
                  "status": 0,
                  "uid": 0,
                  "user_name": "",
                  "head_pic": "",
                  "guard": 0,
                  "start_at": 0,
                  "current_time": 1718211564
              },
              "icons": {
                  "icon_close": "https://i0.hdslb.com/bfs/live/a176d879dffe8de1586a5eb54c2a08a0c7d31392.png",
                  "icon_open": "https://i0.hdslb.com/bfs/live/70f0844c9a12d29db1e586485954290144534be9.png",
                  "icon_wait": "https://i0.hdslb.com/bfs/live/1049bb88f1e7afd839cc1de80e13228ccd5807e8.png",
                  "icon_starting": "https://i0.hdslb.com/bfs/live/948433d1647a0704f8216f017c406224f9fff518.gif"
              },
              "web_share_link": "https://live.bilibili.com/h5/23075997"
          },
          "ad_banner_info": {
              "data": null
          },
          "skin_info": {
              "id": 65,
              "skin_name": "百人舰队皮肤",
              "skin_config": "{\"zip\":\"http://i0.hdslb.com/bfs/live/fe21bb238c7ca435dc09e78a4b838d78b3435dbd.zip\",\"md5\":\"2AEB56D860055A3F897FDB4BC481DFEE\",\"platform\":\"web\",\"version\":\"1\",\"headInfoBgPic\":\"http://i0.hdslb.com/bfs/live/7ba5a32cda0f985aa02bb05f453eac1f03cb976d.png\",\"giftControlBgPic\":\"http://i0.hdslb.com/bfs/live/f864809e6be7b3fe834de6d37b5b7f42b2cdff2a.png\",\"rankListBgPic\":\"http://i0.hdslb.com/bfs/live/00d1718591af1b5117c588f7bac1efb6c1e97fde.png\",\"mainText\":\"#FFffd423\",\"normalText\":\"#FFffffff\",\"highlightContent\":\"#FFffd432\",\"border\":\"#FF999999\",\"infoCardBgPic\":\"\"}",
              "show_text": "保持直播间大航海人数100-999人",
              "skin_url": "https://i0.hdslb.com/bfs/live/10143ee1233a4b88dc902a4dc92c9f30a4ca0241.png",
              "start_time": 1715260781,
              "end_time": 2145888000,
              "current_time": 1718211564
          },
          "web_banner_info": {
              "id": 0,
              "title": "",
              "left": "",
              "right": "",
              "jump_url": "",
              "bg_color": "",
              "hover_color": "",
              "text_bg_color": "",
              "text_hover_color": "",
              "link_text": "",
              "link_color": "",
              "input_color": "",
              "input_text_color": "",
              "input_hover_color": "",
              "input_border_color": "",
              "input_search_color": ""
          },
          "lol_info": null,
          "pk_info": null,
          "battle_info": null,
          "silent_room_info": {
              "type": "",
              "level": 0,
              "second": 0,
              "expire_time": 0
          },
          "switch_info": {
              "close_guard": false,
              "close_gift": false,
              "close_online": false,
              "close_danmaku": false
          },
          "record_switch_info": null,
          "room_config_info": {
              "dm_text": "发个弹幕呗~"
          },
          "gift_memory_info": {
              "list": null
          },
          "new_switch_info": {
              "room-socket": 1,
              "room-prop-send": 1,
              "room-sailing": 1,
              "room-info-popularity": 1,
              "room-danmaku-editor": 1,
              "room-effect": 1,
              "room-fans_medal": 1,
              "room-report": 1,
              "room-feedback": 1,
              "room-player-watermark": 1,
              "room-recommend-live_off": 1,
              "room-activity": 1,
              "room-web_banner": 1,
              "room-silver_seeds-box": 1,
              "room-wishing_bottle": 1,
              "room-board": 1,
              "room-supplication": 1,
              "room-hour_rank": 1,
              "room-week_rank": 1,
              "room-anchor_rank": 1,
              "room-info-integral": 1,
              "room-super-chat": 1,
              "room-tab": 1,
              "room-hot-rank": 1,
              "fans-medal-progress": 1,
              "gift-bay-screen": 1,
              "room-enter": 1,
              "room-my-idol": 1,
              "room-topic": 1,
              "fans-club": 1,
              "room-popular-rank": 1,
              "mic_user_gift": 1,
              "new-room-area-rank": 1,
              "wealth_medal": 1,
              "bubble": 1,
              "title": 1,
              "room_rank_rearrange": 1,
              "web-gift-batter-bar": 1,
              "popular_rank_anchor_ab": 1
          },
          "super_chat_info": {
              "status": 1,
              "jump_url": "https://live.bilibili.com/p/html/live-app-superchat2/index.html?is_live_half_webview=1\u0026hybrid_half_ui=1,3,100p,70p,ffffff,0,30,100,12,0;2,2,375,100p,ffffff,0,30,100,0,0;3,3,100p,70p,ffffff,0,30,100,12,0;4,2,375,100p,ffffff,0,30,100,0,0;5,3,100p,60p,ffffff,0,30,100,12,0;6,3,100p,60p,ffffff,0,30,100,12,0;7,3,100p,60p,ffffff,0,30,100,12,0",
              "icon": "https://i0.hdslb.com/bfs/live/0a9ebd72c76e9cbede9547386dd453475d4af6fe.png",
              "ranked_mark": 0,
              "message_list": []
          },
          "online_gold_rank_info_v2": {
              "list": [
                  {
                      "uid": 95203965,
                      "face": "https://i0.hdslb.com/bfs/face/f4de9aa5ad241f5a5e6c8478fef73f12f17b4f81.jpg",
                      "uname": "特麗珑",
                      "score": "1380",
                      "rank": 1,
                      "guard_level": 3,
                      "wealth_level": 29,
                      "is_mystery": false,
                      "uinfo": {
                          "uid": 95203965,
                          "base": {
                              "name": "特麗珑",
                              "face": "https://i0.hdslb.com/bfs/face/f4de9aa5ad241f5a5e6c8478fef73f12f17b4f81.jpg",
                              "name_color": 0,
                              "is_mystery": false,
                              "risk_ctrl_info": {
                                  "name": "特麗珑",
                                  "face": "https://i0.hdslb.com/bfs/face/f4de9aa5ad241f5a5e6c8478fef73f12f17b4f81.jpg"
                              },
                              "origin_info": {
                                  "name": "特麗珑",
                                  "face": "https://i0.hdslb.com/bfs/face/f4de9aa5ad241f5a5e6c8478fef73f12f17b4f81.jpg"
                              },
                              "official_info": {
                                  "role": 0,
                                  "title": "",
                                  "desc": "",
                                  "type": -1
                              },
                              "name_color_str": ""
                          },
                          "medal": null,
                          "wealth": null,
                          "title": null,
                          "guard": null,
                          "uhead_frame": null,
                          "guard_leader": null
                      }
                  },
                  {
                      "uid": 5881429,
                      "face": "https://i1.hdslb.com/bfs/face/da06cc9ca255964ab9b3111992cea83a6599a28d.jpg",
                      "uname": "小春寒",
                      "score": "1380",
                      "rank": 2,
                      "guard_level": 3,
                      "wealth_level": 28,
                      "is_mystery": false,
                      "uinfo": {
                          "uid": 5881429,
                          "base": {
                              "name": "小春寒",
                              "face": "https://i1.hdslb.com/bfs/face/da06cc9ca255964ab9b3111992cea83a6599a28d.jpg",
                              "name_color": 0,
                              "is_mystery": false,
                              "risk_ctrl_info": {
                                  "name": "小春寒",
                                  "face": "https://i1.hdslb.com/bfs/face/da06cc9ca255964ab9b3111992cea83a6599a28d.jpg"
                              },
                              "origin_info": {
                                  "name": "小春寒",
                                  "face": "https://i1.hdslb.com/bfs/face/da06cc9ca255964ab9b3111992cea83a6599a28d.jpg"
                              },
                              "official_info": {
                                  "role": 0,
                                  "title": "",
                                  "desc": "",
                                  "type": -1
                              },
                              "name_color_str": ""
                          },
                          "medal": null,
                          "wealth": null,
                          "title": null,
                          "guard": null,
                          "uhead_frame": null,
                          "guard_leader": null
                      }
                  },
                  {
                      "uid": 156077349,
                      "face": "https://i0.hdslb.com/bfs/face/32fdee8a66a61d38b6e65d3310363c6af46480a2.jpg",
                      "uname": "雨齐青",
                      "score": "73",
                      "rank": 3,
                      "guard_level": 2,
                      "wealth_level": 42,
                      "is_mystery": false,
                      "uinfo": {
                          "uid": 156077349,
                          "base": {
                              "name": "雨齐青",
                              "face": "https://i0.hdslb.com/bfs/face/32fdee8a66a61d38b6e65d3310363c6af46480a2.jpg",
                              "name_color": 0,
                              "is_mystery": false,
                              "risk_ctrl_info": {
                                  "name": "雨齐青",
                                  "face": "https://i0.hdslb.com/bfs/face/32fdee8a66a61d38b6e65d3310363c6af46480a2.jpg"
                              },
                              "origin_info": {
                                  "name": "雨齐青",
                                  "face": "https://i0.hdslb.com/bfs/face/32fdee8a66a61d38b6e65d3310363c6af46480a2.jpg"
                              },
                              "official_info": {
                                  "role": 0,
                                  "title": "",
                                  "desc": "",
                                  "type": -1
                              },
                              "name_color_str": ""
                          },
                          "medal": null,
                          "wealth": null,
                          "title": null,
                          "guard": null,
                          "uhead_frame": null,
                          "guard_leader": null
                      }
                  },
                  {
                      "uid": 21048761,
                      "face": "https://i2.hdslb.com/bfs/face/82d1561e3df7cfc73176d06a5de750bdcc66e89b.jpg",
                      "uname": "随喵的令和",
                      "score": "20",
                      "rank": 4,
                      "guard_level": 3,
                      "wealth_level": 38,
                      "is_mystery": false,
                      "uinfo": {
                          "uid": 21048761,
                          "base": {
                              "name": "随喵的令和",
                              "face": "https://i2.hdslb.com/bfs/face/82d1561e3df7cfc73176d06a5de750bdcc66e89b.jpg",
                              "name_color": 0,
                              "is_mystery": false,
                              "risk_ctrl_info": {
                                  "name": "随喵的令和",
                                  "face": "https://i2.hdslb.com/bfs/face/82d1561e3df7cfc73176d06a5de750bdcc66e89b.jpg"
                              },
                              "origin_info": {
                                  "name": "随喵的令和",
                                  "face": "https://i2.hdslb.com/bfs/face/82d1561e3df7cfc73176d06a5de750bdcc66e89b.jpg"
                              },
                              "official_info": {
                                  "role": 0,
                                  "title": "",
                                  "desc": "",
                                  "type": -1
                              },
                              "name_color_str": ""
                          },
                          "medal": null,
                          "wealth": null,
                          "title": null,
                          "guard": null,
                          "uhead_frame": null,
                          "guard_leader": null
                      }
                  },
                  {
                      "uid": 398539061,
                      "face": "https://i2.hdslb.com/bfs/face/d2b20943fe445a6be602635507dcba4bb3164322.jpg",
                      "uname": "别逼我跪下求你看MyGO",
                      "score": "20",
                      "rank": 5,
                      "guard_level": 0,
                      "wealth_level": 23,
                      "is_mystery": false,
                      "uinfo": {
                          "uid": 398539061,
                          "base": {
                              "name": "别逼我跪下求你看MyGO",
                              "face": "https://i2.hdslb.com/bfs/face/d2b20943fe445a6be602635507dcba4bb3164322.jpg",
                              "name_color": 0,
                              "is_mystery": false,
                              "risk_ctrl_info": {
                                  "name": "别逼我跪下求你看MyGO",
                                  "face": "https://i2.hdslb.com/bfs/face/d2b20943fe445a6be602635507dcba4bb3164322.jpg"
                              },
                              "origin_info": {
                                  "name": "别逼我跪下求你看MyGO",
                                  "face": "https://i2.hdslb.com/bfs/face/d2b20943fe445a6be602635507dcba4bb3164322.jpg"
                              },
                              "official_info": {
                                  "role": 0,
                                  "title": "",
                                  "desc": "",
                                  "type": -1
                              },
                              "name_color_str": ""
                          },
                          "medal": null,
                          "wealth": null,
                          "title": null,
                          "guard": null,
                          "uhead_frame": null,
                          "guard_leader": null
                      }
                  },
                  {
                      "uid": 36419356,
                      "face": "https://i1.hdslb.com/bfs/face/2a3f9b5b8eed7bb50b5a76b38f78380c29e567af.jpg",
                      "uname": "皮卡丘从小就很萌",
                      "score": "20",
                      "rank": 6,
                      "guard_level": 0,
                      "wealth_level": 26,
                      "is_mystery": false,
                      "uinfo": {
                          "uid": 36419356,
                          "base": {
                              "name": "皮卡丘从小就很萌",
                              "face": "https://i1.hdslb.com/bfs/face/2a3f9b5b8eed7bb50b5a76b38f78380c29e567af.jpg",
                              "name_color": 0,
                              "is_mystery": false,
                              "risk_ctrl_info": {
                                  "name": "皮卡丘从小就很萌",
                                  "face": "https://i1.hdslb.com/bfs/face/2a3f9b5b8eed7bb50b5a76b38f78380c29e567af.jpg"
                              },
                              "origin_info": {
                                  "name": "皮卡丘从小就很萌",
                                  "face": "https://i1.hdslb.com/bfs/face/2a3f9b5b8eed7bb50b5a76b38f78380c29e567af.jpg"
                              },
                              "official_info": {
                                  "role": 0,
                                  "title": "",
                                  "desc": "",
                                  "type": -1
                              },
                              "name_color_str": ""
                          },
                          "medal": null,
                          "wealth": null,
                          "title": null,
                          "guard": null,
                          "uhead_frame": null,
                          "guard_leader": null
                      }
                  },
                  {
                      "uid": 16046433,
                      "face": "https://i0.hdslb.com/bfs/face/39945a0629e5c383c7d335056e4411350e2f7610.jpg",
                      "uname": "梦回灬如初",
                      "score": "20",
                      "rank": 7,
                      "guard_level": 3,
                      "wealth_level": 33,
                      "is_mystery": false,
                      "uinfo": {
                          "uid": 16046433,
                          "base": {
                              "name": "梦回灬如初",
                              "face": "https://i0.hdslb.com/bfs/face/39945a0629e5c383c7d335056e4411350e2f7610.jpg",
                              "name_color": 0,
                              "is_mystery": false,
                              "risk_ctrl_info": {
                                  "name": "梦回灬如初",
                                  "face": "https://i0.hdslb.com/bfs/face/39945a0629e5c383c7d335056e4411350e2f7610.jpg"
                              },
                              "origin_info": {
                                  "name": "梦回灬如初",
                                  "face": "https://i0.hdslb.com/bfs/face/39945a0629e5c383c7d335056e4411350e2f7610.jpg"
                              },
                              "official_info": {
                                  "role": 0,
                                  "title": "",
                                  "desc": "",
                                  "type": -1
                              },
                              "name_color_str": ""
                          },
                          "medal": null,
                          "wealth": null,
                          "title": null,
                          "guard": null,
                          "uhead_frame": null,
                          "guard_leader": null
                      }
                  }
              ],
              "count": 116,
              "count_text": "116"
          },
          "dm_brush_info": {
              "min_time": 700,
              "brush_count": 100,
              "slice_count": 2,
              "storage_time": 3000
          },
          "dm_emoticon_info": {
              "is_open_emoticon": 1,
              "is_shield_emoticon": 0
          },
          "dm_tag_info": {
              "dm_tag": 0,
              "platform": null,
              "extra": "",
              "dm_chronos_extra": "",
              "dm_mode": null,
              "dm_setting_switch": 0,
              "material_conf": null
          },
          "topic_info": {
              "topic_id": 0,
              "topic_name": ""
          },
          "game_info": {
              "game_status": 0
          },
          "watched_show": {
              "switch": true,
              "num": 880,
              "text_small": "880",
              "text_large": "880人看过",
              "icon": "",
              "icon_location": 0,
              "icon_web": ""
          },
          "topic_room_info": {
              "interactive_h5_url": "",
              "watermark": 1
          },
          "show_reserve_status": false,
          "second_create_info": null,
          "play_together_info": null,
          "cloud_game_info": {
              "is_gaming": 0
          },
          "like_info_v3": {
              "total_likes": 0,
              "click_block": false,
              "count_block": false,
              "guild_emo_text": "试试双击点赞 让主播被更多人看到吧～",
              "guild_dm_text": "点赞30次可以帮主播冲刺热门榜哦～",
              "like_dm_text": "谢谢你的赞，每点赞30次有概率为主播增加曝光哦～",
              "hand_icons": [
                  "https://i0.hdslb.com/bfs/live/1aac78a01b4cba5aa3f7c935c49b0a0bd8992ce1.webp",
                  "https://i0.hdslb.com/bfs/live/81c69596db4a94c53bebed567c1e457c8240e91c.webp",
                  "https://i0.hdslb.com/bfs/live/f281abe0bb186253acfe7b82d54d62fd25f5fee3.webp",
                  "https://i0.hdslb.com/bfs/live/e99441bacfb0a516a08455c5eb9266e7d6174a5c.webp",
                  "https://i0.hdslb.com/bfs/live/2e9145b42b5acbd4a13de1540b42e709832f9bae.webp",
                  "https://i0.hdslb.com/bfs/live/09f36c21175ca217268896305f1d248f3f64add0.webp",
                  "https://i0.hdslb.com/bfs/live/8e90f8aedd66237922958654efb36d97ac4a5701.webp",
                  "https://i0.hdslb.com/bfs/live/29b65444d287680263b1024a6a0d479f7ad41ad7.webp",
                  "https://i0.hdslb.com/bfs/live/b0faa7d645c035cd1d61de5011cdc3880ef232d6.webp",
                  "https://i0.hdslb.com/bfs/live/10fe45806cb0c3e13ef5c9a3ca03a0707667934b.webp"
              ],
              "dm_icons": [
                  "https://i0.hdslb.com/bfs/live/ecc2c5a2efc1c40a1125bb0d648d891e0e8cbd3b.png",
                  "https://i0.hdslb.com/bfs/live/391b45ed00617391b863f6441d5d040165c4d1b4.png",
                  "https://i0.hdslb.com/bfs/live/ad7c7adc9ee0c778cea161dea6c5a96c4ff7f845.png",
                  "https://i0.hdslb.com/bfs/live/c6a4781558d84838cc4963fcf94792e54beabe15.png",
                  "https://i0.hdslb.com/bfs/live/6920410aca21e3397f23ee7ae5ddbf8f1c625943.png",
                  "https://i0.hdslb.com/bfs/live/7e671e053c365f91a0a0f4837ec92ac9581977ec.png"
              ],
              "eggshells_icon": "https://i0.hdslb.com/bfs/live/ea6454847e9b9d1b6be03f04f829620b4f0a5d46.svga",
              "count_show_time": 15,
              "process_icon": "https://i0.hdslb.com/bfs/live/6d96e9de0cc5c80acf3537fe9a23d94da0ab5ff5.png",
              "process_color": "#4DFF6699",
              "report_click_limit": 15,
              "report_time_min": 5,
              "report_time_max": 10,
              "icon": "https://i0.hdslb.com/bfs/live/fe8e848118e8f18eccbd11cad62b29105c51a797.png",
              "cooldown": 0.35,
              "hand_use_face": true,
              "guide_icon_urls": [
                  "https://i0.hdslb.com/bfs/live/2e9145b42b5acbd4a13de1540b42e709832f9bae.webp",
                  "https://i0.hdslb.com/bfs/live/b0faa7d645c035cd1d61de5011cdc3880ef232d6.webp"
              ],
              "guide_icon_ratio": 1.3
          },
          "live_play_info": {
              "show_widget_banner": true,
              "show_left_entry": true,
              "widget_version": 1
          },
          "multi_voice": {
              "switch_status": 2,
              "members": [],
              "mv_role": 0,
              "seat_type": 0,
              "invoking_time": 0,
              "version": 0,
              "pk": null,
              "biz_session_id": "",
              "mode_details": null,
              "hat_list": null,
              "battle_info": null
          },
          "popular_rank_info": {
              "rank": 0,
              "countdown": 37,
              "timestamp": 1718211564,
              "url": "https://live.bilibili.com/p/html/live-app-hotrank/index.html?is_live_half_webview=1\u0026hybrid_rotate_d=1\u0026hybrid_half_ui=1,3,100p,70p,0,0,30,100,12;2,2,375,100p,0,0,30,100,0;3,3,100p,70p,0,0,30,100,12;4,2,375,100p,0,0,30,100,0;5,3,100p,70p,0,0,30,100,0;6,3,100p,70p,0,0,30,100,0;7,3,100p,70p,0,0,30,100,0;8,3,100p,70p,0,0,30,100,0\u0026pc_ui=338,465,f4eefa,0\u0026redirect=v2\u0026anchorId=1355412269",
              "on_rank_name": "人气",
              "rank_name": "人气榜"
          },
          "new_area_rank_info": {
              "items": [
                  {
                      "conf_id": 11,
                      "rank_name": "虚拟航海",
                      "uid": 1355412269,
                      "rank": 0,
                      "icon_url_blue": "https://i0.hdslb.com/bfs/live/18e2990a546d33368200f9058f3d9dbc4038eb5c.png",
                      "icon_url_pink": "https://i0.hdslb.com/bfs/live/a6c490c36e88c7b191a04883a5ec15aed187a8f7.png",
                      "icon_url_grey": "https://i0.hdslb.com/bfs/live/cb7444b1faf1d785df6265bfdc1fcfc993419b76.png",
                      "jump_url_link": "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=3\u0026ruid=1355412269\u0026conf_id=11\u0026is_live_half_webview=1\u0026hybrid_rotate_d=1\u0026is_cling_player=1\u0026hybrid_half_ui=1,3,100p,70p,f4eefa,0,30,100,0,0;2,2,375,100p,f4eefa,0,30,100,0,0;3,3,100p,70p,f4eefa,0,30,100,0,0;4,2,375,100p,f4eefa,0,30,100,0,0;5,3,100p,70p,f4eefa,0,30,100,0,0;6,3,100p,70p,f4eefa,0,30,100,0,0;7,3,100p,70p,f4eefa,0,30,100,0,0;8,3,100p,70p,f4eefa,0,30,100,0,0#/area-rank",
                      "jump_url_pc": "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=4\u0026ruid=1355412269\u0026conf_id=11\u0026pc_ui=338,465,f4eefa,0#/area-rank",
                      "jump_url_pink": "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=1\u0026ruid=1355412269\u0026conf_id=11\u0026is_live_half_webview=1\u0026hybrid_rotate_d=1\u0026hybrid_half_ui=1,3,100p,70p,ffffff,0,30,100,12,0;2,2,375,100p,ffffff,0,30,100,0,0;3,3,100p,70p,ffffff,0,30,100,12,0;4,2,375,100p,ffffff,0,30,100,0,0;5,3,100p,70p,ffffff,0,30,100,0,0;6,3,100p,70p,ffffff,0,30,100,0,0;7,3,100p,70p,ffffff,0,30,100,0,0;8,3,100p,70p,ffffff,0,30,100,0,0#/area-rank",
                      "jump_url_web": "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=2\u0026ruid=1355412269\u0026conf_id=11#/area-rank"
                  }
              ],
              "rotation_cycle_time_web": 10000
          },
          "gift_star": {
              "show": true,
              "display_widget_ab_group": 0
          },
          "progress_for_widget": {
              "gift_star_process": {
                  "task_info": {
                      "start_date": 20240610,
                      "process_list": [
                          {
                              "gift_id": 31037,
                              "gift_img": "https://s1.hdslb.com/bfs/live/461be640f60788c1d159ec8d6c5d5cf1ef3d1830.png",
                              "gift_name": "礼物星球",
                              "completed_num": 29,
                              "target_num": 30
                          },
                          {
                              "gift_id": 31049,
                              "gift_img": "https://s1.hdslb.com/bfs/live/96ec38f351a4e190c4a525bc5e11ff09d2874064.png",
                              "gift_name": "礼物星球",
                              "completed_num": 1,
                              "target_num": 10
                          },
                          {
                              "gift_id": 31588,
                              "gift_img": "https://s1.hdslb.com/bfs/live/311930350df3b8e467d13b992e62344ca1e3664f.png",
                              "gift_name": "礼物星球",
                              "completed_num": 0,
                              "target_num": 5
                          }
                      ],
                      "finished": false,
                      "ddl_timestamp": 1718553600,
                      "version": 1718205471608,
                      "reward_gift": 0,
                      "reward_gift_img": "https://i0.hdslb.com/bfs/live/52edb4ab7377ece34ac15b21154d13d188874b01.png",
                      "reward_gift_name": "礼物星球",
                      "level_info": {
                          "star_name": "礼物星球",
                          "level_tip": "暂未达成",
                          "level_img": "https://i0.hdslb.com/bfs/live/91292553d02aa10a133cdf6d8e579d8588e2067a.png",
                          "level_id": 0
                      }
                  },
                  "preload_timestamp": 1718553166,
                  "preload": false,
                  "preload_task_info": null,
                  "widget_bg": "",
                  "jump_schema": "https://live.bilibili.com/p/html/live-app-gift-planet/index.html?roomId=23075997\u0026ruid=1355412269\u0026app_common=open\u0026plt=\u0026pc_ui=375,580,2E345A,0\u0026is_live_half_webview=1\u0026hybrid_half_ui=1,3,100p,70p,0,0,30,100,12,0;2,2,375,100p,0,0,30,100,0,0;3,3,100p,70p,0,0,30,100,12,0;4,2,375,100p,0,0,30,100,0,0;5,3,100p,70p,0,0,30,100,12,0;6,3,100p,70p,0,0,30,100,12,0;7,3,100p,70p,0,0,30,100,12,0;8,2,320,480,0,0,30,100,0,0\u0026display_widget_ab_group=0",
                  "ab_group": 0
              },
              "wish_process": null,
              "star_knight": null,
              "collection_praise_process": {
                  "id": 0,
                  "uid": 0,
                  "target_praise": 0,
                  "current_praise": 0,
                  "start_time": 0,
                  "end_time": 0,
                  "benefit": "",
                  "isSuccess": false,
                  "exist": false,
                  "audit_status": 0,
                  "jump_url": "",
                  "current_praise_text": "",
                  "icon_url": "https://i0.hdslb.com/bfs/live/fb44b6cc32793f15eb588ca93d2d3fc522c31dfd.png",
                  "live_id": ""
              }
          },
          "revenue_demotion": {
              "global_gift_config_demotion": false
          },
          "revenue_material_md5": null,
          "block_info": {
              "block": true
          },
          "danmu_extra": {
              "screen_switch_off": false
          },
          "video_connection_info": null,
          "player_throttle_info": {
              "status": 0,
              "normal_sleep_time": 0,
              "fullscreen_sleep_time": 0,
              "tab_sleep_time": 0,
              "prompt_time": 0
          },
          "guard_info": {
              "count": 242,
              "anchor_guard_achieve_level": 100
          },
          "hot_rank_info": null,
          "room_rank_info": {
              "anchor_rank_entry": null,
              "user_rank_entry": {
                  "user_contribution_rank_entry": {
                      "item": [
                          {
                              "uid": 95203965,
                              "name": "特麗珑",
                              "face": "https://i0.hdslb.com/bfs/face/f4de9aa5ad241f5a5e6c8478fef73f12f17b4f81.jpg",
                              "rank": 1,
                              "score": 1380,
                              "medal_info": null,
                              "guard_level": 0,
                              "wealth_level": 0,
                              "is_mystery": false,
                              "uinfo": {
                                  "uid": 95203965,
                                  "base": {
                                      "name": "特麗珑",
                                      "face": "https://i0.hdslb.com/bfs/face/f4de9aa5ad241f5a5e6c8478fef73f12f17b4f81.jpg",
                                      "name_color": 0,
                                      "is_mystery": false,
                                      "risk_ctrl_info": {
                                          "name": "特麗珑",
                                          "face": "https://i0.hdslb.com/bfs/face/f4de9aa5ad241f5a5e6c8478fef73f12f17b4f81.jpg"
                                      },
                                      "origin_info": {
                                          "name": "特麗珑",
                                          "face": "https://i0.hdslb.com/bfs/face/f4de9aa5ad241f5a5e6c8478fef73f12f17b4f81.jpg"
                                      },
                                      "official_info": {
                                          "role": 0,
                                          "title": "",
                                          "desc": "",
                                          "type": -1
                                      },
                                      "name_color_str": ""
                                  },
                                  "medal": null,
                                  "wealth": null,
                                  "title": {
                                      "old_title_css_id": "",
                                      "title_css_id": ""
                                  },
                                  "guard": null,
                                  "uhead_frame": null,
                                  "guard_leader": null
                              }
                          },
                          {
                              "uid": 5881429,
                              "name": "小春寒",
                              "face": "https://i1.hdslb.com/bfs/face/da06cc9ca255964ab9b3111992cea83a6599a28d.jpg",
                              "rank": 2,
                              "score": 1380,
                              "medal_info": null,
                              "guard_level": 0,
                              "wealth_level": 0,
                              "is_mystery": false,
                              "uinfo": {
                                  "uid": 5881429,
                                  "base": {
                                      "name": "小春寒",
                                      "face": "https://i1.hdslb.com/bfs/face/da06cc9ca255964ab9b3111992cea83a6599a28d.jpg",
                                      "name_color": 0,
                                      "is_mystery": false,
                                      "risk_ctrl_info": {
                                          "name": "小春寒",
                                          "face": "https://i1.hdslb.com/bfs/face/da06cc9ca255964ab9b3111992cea83a6599a28d.jpg"
                                      },
                                      "origin_info": {
                                          "name": "小春寒",
                                          "face": "https://i1.hdslb.com/bfs/face/da06cc9ca255964ab9b3111992cea83a6599a28d.jpg"
                                      },
                                      "official_info": {
                                          "role": 0,
                                          "title": "",
                                          "desc": "",
                                          "type": -1
                                      },
                                      "name_color_str": ""
                                  },
                                  "medal": null,
                                  "wealth": null,
                                  "title": {
                                      "old_title_css_id": "",
                                      "title_css_id": ""
                                  },
                                  "guard": null,
                                  "uhead_frame": null,
                                  "guard_leader": null
                              }
                          },
                          {
                              "uid": 156077349,
                              "name": "雨齐青",
                              "face": "https://i0.hdslb.com/bfs/face/32fdee8a66a61d38b6e65d3310363c6af46480a2.jpg",
                              "rank": 3,
                              "score": 73,
                              "medal_info": null,
                              "guard_level": 0,
                              "wealth_level": 0,
                              "is_mystery": false,
                              "uinfo": {
                                  "uid": 156077349,
                                  "base": {
                                      "name": "雨齐青",
                                      "face": "https://i0.hdslb.com/bfs/face/32fdee8a66a61d38b6e65d3310363c6af46480a2.jpg",
                                      "name_color": 0,
                                      "is_mystery": false,
                                      "risk_ctrl_info": {
                                          "name": "雨齐青",
                                          "face": "https://i0.hdslb.com/bfs/face/32fdee8a66a61d38b6e65d3310363c6af46480a2.jpg"
                                      },
                                      "origin_info": {
                                          "name": "雨齐青",
                                          "face": "https://i0.hdslb.com/bfs/face/32fdee8a66a61d38b6e65d3310363c6af46480a2.jpg"
                                      },
                                      "official_info": {
                                          "role": 0,
                                          "title": "",
                                          "desc": "",
                                          "type": -1
                                      },
                                      "name_color_str": ""
                                  },
                                  "medal": null,
                                  "wealth": null,
                                  "title": {
                                      "old_title_css_id": "",
                                      "title_css_id": ""
                                  },
                                  "guard": null,
                                  "uhead_frame": null,
                                  "guard_leader": null
                              }
                          }
                      ],
                      "count": 17,
                      "show_max": 3,
                      "count_text": "17"
                  }
              },
              "user_rank_tab_list": {
                  "tab": [
                      {
                          "type": "contribution_tab",
                          "title": "高能用户",
                          "status": 1,
                          "default": 1,
                          "comment": "",
                          "desc_url": "",
                          "switch": null,
                          "sub_tab": [
                              {
                                  "type": "online_rank",
                                  "title": "在线榜",
                                  "status": 1,
                                  "default": 1,
                                  "comment": "",
                                  "desc_url": "https://live.bilibili.com/p/contribute-rank-description-h5/index.html?is_live_half_webview=1\u0026hybrid_half_ui=1,3,100p,70p,0,0,30,100,12,0;2,2,375,100p,0,0,30,100,0,0;3,3,100p,70p,0,0,30,100,12,0;4,2,375,100p,0,0,30,100,0,0;5,3,100p,70p,0,0,30,100,12,0;6,3,100p,70p,0,0,30,100,12,0;7,3,100p,70p,0,0,30,100,12,0;8,3,100p,70p,0,0,30,100,0,0\u0026rank_type=1\u0026test_type=b",
                                  "switch": [
                                      {
                                          "text": "贡献值",
                                          "switch": "contribution_rank",
                                          "ui_type": {
                                              "op_button_text": 1,
                                              "rank_prefix": 1,
                                              "refresh_entry": 1,
                                              "show_score": 1
                                          },
                                          "comment": "投喂、发弹幕均可获得贡献值"
                                      },
                                      {
                                          "text": "进房时间",
                                          "switch": "entry_time_rank",
                                          "ui_type": {
                                              "op_button_text": 2,
                                              "rank_prefix": 0,
                                              "show_score": 0
                                          },
                                          "comment": "用户进房顺序排序"
                                      }
                                  ],
                                  "sub_tab": null
                              },
                              {
                                  "type": "daily_rank",
                                  "title": "日榜",
                                  "status": 1,
                                  "default": 0,
                                  "comment": "",
                                  "desc_url": "https://live.bilibili.com/p/contribute-rank-description-h5/index.html?is_live_half_webview=1\u0026hybrid_half_ui=1,3,100p,70p,0,0,30,100,12,0;2,2,375,100p,0,0,30,100,0,0;3,3,100p,70p,0,0,30,100,12,0;4,2,375,100p,0,0,30,100,0,0;5,3,100p,70p,0,0,30,100,12,0;6,3,100p,70p,0,0,30,100,12,0;7,3,100p,70p,0,0,30,100,12,0;8,3,100p,70p,0,0,30,100,0,0\u0026rank_type=2\u0026test_type=b",
                                  "switch": [
                                      {
                                          "text": "当日",
                                          "switch": "today_rank",
                                          "ui_type": {
                                              "op_button_text": 1,
                                              "rank_prefix": 0,
                                              "show_score": 0
                                          },
                                          "comment": "用户当日贡献值排序"
                                      },
                                      {
                                          "text": "昨日",
                                          "switch": "yesterday_rank",
                                          "ui_type": {
                                              "op_button_text": 0,
                                              "rank_prefix": 0,
                                              "show_score": 0
                                          },
                                          "comment": "昨日贡献值前100排名"
                                      }
                                  ],
                                  "sub_tab": null
                              },
                              {
                                  "type": "weekly_rank",
                                  "title": "周榜",
                                  "status": 1,
                                  "default": 0,
                                  "comment": "",
                                  "desc_url": "https://live.bilibili.com/p/contribute-rank-description-h5/index.html?is_live_half_webview=1\u0026hybrid_half_ui=1,3,100p,70p,0,0,30,100,12,0;2,2,375,100p,0,0,30,100,0,0;3,3,100p,70p,0,0,30,100,12,0;4,2,375,100p,0,0,30,100,0,0;5,3,100p,70p,0,0,30,100,12,0;6,3,100p,70p,0,0,30,100,12,0;7,3,100p,70p,0,0,30,100,12,0;8,3,100p,70p,0,0,30,100,0,0\u0026rank_type=2\u0026test_type=b",
                                  "switch": [
                                      {
                                          "text": "本周",
                                          "switch": "current_week_rank",
                                          "ui_type": {
                                              "op_button_text": 1,
                                              "rank_prefix": 0,
                                              "show_score": 0
                                          },
                                          "comment": "用户当周贡献值排序"
                                      },
                                      {
                                          "text": "上周",
                                          "switch": "last_week_rank",
                                          "ui_type": {
                                              "op_button_text": 0,
                                              "rank_prefix": 0,
                                              "show_score": 0
                                          },
                                          "comment": "上周贡献值前100排名"
                                      }
                                  ],
                                  "sub_tab": null
                              },
                              {
                                  "type": "monthly_rank",
                                  "title": "月榜",
                                  "status": 1,
                                  "default": 0,
                                  "comment": "",
                                  "desc_url": "https://live.bilibili.com/p/contribute-rank-description-h5/index.html?is_live_half_webview=1\u0026hybrid_half_ui=1,3,100p,70p,0,0,30,100,12,0;2,2,375,100p,0,0,30,100,0,0;3,3,100p,70p,0,0,30,100,12,0;4,2,375,100p,0,0,30,100,0,0;5,3,100p,70p,0,0,30,100,12,0;6,3,100p,70p,0,0,30,100,12,0;7,3,100p,70p,0,0,30,100,12,0;8,3,100p,70p,0,0,30,100,0,0\u0026rank_type=2\u0026test_type=b",
                                  "switch": [
                                      {
                                          "text": "本月",
                                          "switch": "current_month_rank",
                                          "ui_type": {
                                              "op_button_text": 1,
                                              "rank_prefix": 0,
                                              "show_score": 0
                                          },
                                          "comment": "用户当月贡献值排序"
                                      },
                                      {
                                          "text": "上月",
                                          "switch": "last_month_rank",
                                          "ui_type": {
                                              "op_button_text": 0,
                                              "rank_prefix": 0,
                                              "show_score": 0
                                          },
                                          "comment": "上月贡献值前100排名"
                                      }
                                  ],
                                  "sub_tab": null
                              }
                          ]
                      },
                      {
                          "type": "guard_tab",
                          "title": "大航海",
                          "status": 1,
                          "default": 0,
                          "comment": "头号粉丝大航海，上船后可上榜",
                          "desc_url": "",
                          "switch": null,
                          "sub_tab": null
                      }
                  ]
              }
          },
          "dm_reply": {
              "show_reply": true
          },
          "dm_combo": null,
          "dm_vote": null,
          "location": null,
          "interactive_game_tag": {
              "action": 0,
              "game_id": "",
              "game_name": ""
          },
          "video_enhancement": {
              "title": "画质增强",
              "desc": "仅针对超清及以下生效",
              "default_switch_status": 2,
              "highest_quality": 250,
              "is_enabled": true
          },
          "guard_leader": {
              "uid": 0,
              "name": "",
              "face": "",
              "jump_url": "https://live.bilibili.com/p/html/live-app-guard-pilot/index.html?is_live_half_webview=1\u0026hybrid_half_ui=1,3,100p,73p,0,0,30,0,12,0;2,2,375,100p,0,0,30,0,0,0;3,3,100p,73p,0,0,30,0,12,0;4,2,375,100p,0,0,30,0,0,0;5,3,100p,73p,0,0,30,0,12,0;6,3,100p,73p,0,0,30,0,12,0;7,3,100p,73p,0,0,30,0,12,0;8,2,320,480,0,0,30,0,0,0\u0026anchorId=1355412269\u0026roomId=23075997",
              "text": "成为舰队指挥官，解锁房间冠名权益",
              "rank_top_icon1": "https://i0.hdslb.com/bfs/live/64b22e65979b32f7e4e8bec1edb38c697fb320fc.png",
              "rank_top_icon2": "https://i0.hdslb.com/bfs/live/7b9d773c6018ffac9f0eadd3c92f0090e09055f2.png",
              "rank_top_background_url1": "https://i0.hdslb.com/bfs/live/b2832de9cca6a0b3b4872c8d96c05ae713bc51d2.png",
              "rank_top_background_url2": "https://i0.hdslb.com/bfs/live/71397554da8a7bd2ac14905c69891df54ad62ede.png",
              "background_url": "",
              "anchor_background_url": "",
              "input_background_url": "",
              "newly": 0,
              "entry_effect_id": 0,
              "show": 1,
              "rank_top_background_light_url1": "",
              "rank_top_background_light_url2": "https://i0.hdslb.com/bfs/live/c0fbd28b6ddf170b8db2e2c7163eb9d66f66fd8b.png",
              "display_src": "https://i0.hdslb.com/bfs/live/7aed22c78a2a41a5c1b964f1f2a3220c52c1663e.png",
              "avatar_src": "https://i0.hdslb.com/bfs/live/4d1f0d9a39e368c4b9b4128f58f945099a295c39.png",
              "icon_src": "https://i0.hdslb.com/bfs/live/81e9d659b251a74277fe9b1a6f7876c4d00a4950.png"
          },
          "room_anonymous": {
              "open_anonymous": false
          },
          "tab_switches": {
              "subtitle": 0,
              "realtime_data": null
          },
          "universal_interact_info": null,
          "pk_info_v2": null,
          "area_mask_info": null,
          "xtemplate_config": {
              "dm_brush_info": {
                  "landScape": null,
                  "verticalscreen": {
                      "min_time": 500,
                      "brush_count": 50,
                      "slice_count": 2,
                      "storage_time": 500,
                      "is_hide_anti_brush": 0
                  }
              },
              "dm_speed_info": {
                  "landScape": null,
                  "verticalscreen": {
                      "valley": {
                          "consumetime": 500,
                          "consumecount": 2,
                          "animationtime": 0
                      },
                      "peak": {
                          "consumetime": 600,
                          "consumecount": 3,
                          "animationtime": 0
                      },
                      "proportion": 30,
                      "interval": 10000
                  }
              },
              "dm_pool_info": {
                  "landScape": null,
                  "verticalscreen": {
                      "master_ceiling": 100,
                      "master_count": 1,
                      "guest_config": [
                          {
                              "score_floor": 80,
                              "score_ceiling": 999999,
                              "dm_max": 150,
                              "consume": 80
                          },
                          {
                              "score_floor": 32,
                              "score_ceiling": 79,
                              "dm_max": 100,
                              "consume": 10
                          },
                          {
                              "score_floor": 0,
                              "score_ceiling": 31,
                              "dm_max": 100,
                              "consume": 10
                          }
                      ],
                      "timeout": 300000,
                      "unusual_score": 50
                  }
              }
          },
          "dm_activity": {
              "activity_list": null,
              "ts": 1718211564
          },
          "dm_interaction_ab": {
              "102": 0,
              "103": 0,
              "104": 0,
              "105": 0,
              "106": 0
          },
          "guard_intimacy_rank_status": {
              "guard_rank_new_ab": 0,
              "guard_rank_new_total_status": 0,
              "guard_rank_new_month_status": 0,
              "guard_rank_new_week_status": 0
          }
      }
  }
  ```
