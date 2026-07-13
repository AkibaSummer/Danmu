package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/AkibaSummer/Danmu/sdk/spider"
)

func main() {
	room := flag.Int("room", 0, "live room short ID")
	uid := flag.Int64("account-uid", 0, "Bilibili UID")
	buvid := flag.String("account-buvid", "", "buvid3 value")
	cookieFile := flag.String("cookie-file", "", "file containing a Cookie header")
	out := flag.String("out", "conf/danmu_endpoint.json", "output cache file")
	flag.Parse()
	if *room <= 0 {
		fatal("-room is required")
	}
	var cookie string
	if *cookieFile != "" {
		data, err := os.ReadFile(*cookieFile)
		if err != nil {
			fatal(err.Error())
		}
		cookie = string(data)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	state, err := spider.FetchEndpoint(ctx, *room, *uid, *buvid, cookie)
	if err != nil {
		fatal(err.Error())
	}
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		fatal(err.Error())
	}
	if err := os.WriteFile(*out, data, 0600); err != nil {
		fatal(err.Error())
	}
	fmt.Printf("endpoint cache written: room=%d hosts=%d\n", state.RoomID, len(state.Hosts))
}

func fatal(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}
