package spider

import (
	"net/http"
	"strings"
	"testing"
)

func TestCookieValueAndMerge(t *testing.T) {
	original := "SESSDATA=old; bili_jct=csrf; buvid3=device"
	if got := cookieValue(original, "bili_jct"); got != "csrf" {
		t.Fatalf("cookieValue() = %q", got)
	}
	merged := mergeCookies(original, []*http.Cookie{{Name: "SESSDATA", Value: "new"}, {Name: "sid", Value: "1"}})
	for _, expected := range []string{"SESSDATA=new", "bili_jct=csrf", "buvid3=device", "sid=1"} {
		if !strings.Contains(merged, expected) {
			t.Errorf("mergeCookies() missing %q: %s", expected, merged)
		}
	}
}

func TestCorrespondPath(t *testing.T) {
	path, err := correspondPath(1684466082562)
	if err != nil {
		t.Fatal(err)
	}
	if len(path) != 256 {
		t.Fatalf("correspondPath length = %d, want 256", len(path))
	}
}

func TestUnpackMessagesRejectsMalformedPacket(t *testing.T) {
	if _, err := unpackMessages([]byte{0, 0, 0, 16}); err == nil {
		t.Fatal("unpackMessages accepted a truncated packet")
	}
}
