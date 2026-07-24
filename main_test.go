package main

import (
	"errors"
	"testing"

	"github.com/AkibaSummer/Danmu/sdk/spider"
)

func TestAuthFailureRouting(t *testing.T) {
	tests := []struct {
		name       string
		status     spider.AuthStatus
		err        error
		wantQRCode bool
	}{
		{name: "valid login", status: spider.AuthStatus{LoggedIn: true}},
		{name: "expired login", wantQRCode: true},
		{
			name:       "refresh rejected",
			status:     spider.AuthStatus{LoggedIn: true, NeedsRefresh: true},
			err:        &spider.Failure{Kind: spider.FailureAuth, Err: errors.New("refresh rejected")},
			wantQRCode: true,
		},
		{
			name:   "temporary network failure",
			err:    &spider.Failure{Kind: spider.FailureNetwork, Err: errors.New("timeout")},
			status: spider.AuthStatus{LoggedIn: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shouldStartQRCode(tt.status, tt.err)
			if got != tt.wantQRCode {
				t.Fatalf("shouldStartQRCode() = %t, want %t", got, tt.wantQRCode)
			}
		})
	}
}
