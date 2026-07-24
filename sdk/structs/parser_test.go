package structs

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/AkibaSummer/Danmu/sdk/utils/logger"
)

func TestReadableParserGiftFormatsAreEquivalent(t *testing.T) {
	const expected = "测试用户(123456) 赠送了 10 个 小花花"

	legacy := `{"cmd":"SEND_GIFT","data":{"sender_uinfo":{"uid":123456,"base":{"name":"测试用户"}},"num":10,"giftName":"小花花"}}`

	giftPayload := appendProtoVarint(nil, 1, 31036)
	giftPayload = appendProtoBytes(giftPayload, 2, []byte("小花花"))
	giftPayload = appendProtoVarint(giftPayload, 3, 10)
	payload := appendProtoVarint(nil, 1, 123456)
	payload = appendProtoBytes(payload, 2, []byte("测试用户"))
	payload = appendProtoBytes(payload, 10, giftPayload)
	current := `{"cmd":"SEND_GIFT_V2","data":{"pb":"` + base64.StdEncoding.EncodeToString(payload) + `"}}`

	for _, test := range []struct {
		name string
		raw  string
	}{
		{name: "SEND_GIFT", raw: legacy},
		{name: "SEND_GIFT_V2", raw: current},
	} {
		t.Run(test.name, func(t *testing.T) {
			result := ReadableParser(logger.NewMsgInternalLoggerChannelMessage(test.raw))
			if result.Skip {
				t.Fatal("gift event was skipped")
			}
			if !strings.HasSuffix(result.Message, expected) {
				t.Fatalf("message = %q, want suffix %q", result.Message, expected)
			}
		})
	}
}

func TestHistoricalGiftSamples(t *testing.T) {
	samplePath := os.Getenv("DANMU_HISTORY_SAMPLES")
	if samplePath == "" {
		t.Skip("set DANMU_HISTORY_SAMPLES to an NDJSON sample file")
	}

	sampleFile, err := os.Open(samplePath)
	if err != nil {
		t.Fatal(err)
	}
	defer sampleFile.Close()

	var count int
	scanner := bufio.NewScanner(sampleFile)
	scanner.Buffer(make([]byte, 64*1024), 4*1024*1024)
	for scanner.Scan() {
		var sample struct {
			Raw      string `json:"raw"`
			Expected string `json:"expected"`
		}
		if err := json.Unmarshal(scanner.Bytes(), &sample); err != nil {
			t.Fatalf("decode sample %d: %v", count+1, err)
		}
		result := ReadableParser(logger.NewMsgInternalLoggerChannelMessage(sample.Raw))
		if result.Skip {
			t.Fatalf("sample %d was skipped", count+1)
		}
		if !strings.HasSuffix(result.Message, sample.Expected) {
			t.Fatalf("sample %d message = %q, want suffix %q", count+1, result.Message, sample.Expected)
		}
		count++
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
	if count == 0 {
		t.Fatal("no historical samples received")
	}
	t.Logf("validated %d historical gift events", count)
}
