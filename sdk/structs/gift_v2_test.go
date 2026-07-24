package structs

import (
	"encoding/base64"
	"testing"
)

func TestParseSendGiftV2(t *testing.T) {
	giftPayload := appendProtoVarint(nil, 1, 31036)
	giftPayload = appendProtoBytes(giftPayload, 2, []byte("小花花"))
	giftPayload = appendProtoVarint(giftPayload, 3, 1)
	giftPayload = appendProtoFixed32(giftPayload, 16, []byte{0, 0, 128, 63})

	payload := appendProtoVarint(nil, 1, 123456)
	payload = appendProtoBytes(payload, 2, []byte("测试用户"))
	payload = appendProtoBytes(payload, 10, giftPayload)

	got, err := parseSendGiftV2(base64.StdEncoding.EncodeToString(payload))
	if err != nil {
		t.Fatal(err)
	}
	if got.UID != 123456 || got.UserName != "测试用户" || got.GiftID != 31036 || got.GiftName != "小花花" || got.Num != 1 {
		t.Fatalf("unexpected gift: %#v", got)
	}
}

func TestParseSendGiftV2RejectsMalformedPayload(t *testing.T) {
	if _, err := parseSendGiftV2(base64.StdEncoding.EncodeToString([]byte{0x52, 0x7f})); err == nil {
		t.Fatal("expected malformed protobuf to fail")
	}
}

func appendProtoVarint(dst []byte, field int, value uint64) []byte {
	dst = appendUvarint(dst, uint64(field<<3))
	return appendUvarint(dst, value)
}

func appendProtoBytes(dst []byte, field int, value []byte) []byte {
	dst = appendUvarint(dst, uint64(field<<3|2))
	dst = appendUvarint(dst, uint64(len(value)))
	return append(dst, value...)
}

func appendProtoFixed32(dst []byte, field int, value []byte) []byte {
	dst = appendUvarint(dst, uint64(field<<3|5))
	return append(dst, value...)
}

func appendUvarint(dst []byte, value uint64) []byte {
	for value >= 0x80 {
		dst = append(dst, byte(value)|0x80)
		value >>= 7
	}
	return append(dst, byte(value))
}
