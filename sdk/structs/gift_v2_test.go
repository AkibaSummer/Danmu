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

	gifts, err := parseSendGiftV2(base64.StdEncoding.EncodeToString(payload))
	if err != nil {
		t.Fatal(err)
	}
	if len(gifts) != 1 {
		t.Fatalf("got %d gifts, want 1", len(gifts))
	}
	got := gifts[0]
	if got.UID != 123456 || got.UserName != "测试用户" || got.GiftID != 31036 || got.GiftName != "小花花" || got.Num != 1 {
		t.Fatalf("unexpected gift: %#v", got)
	}
}

func TestParseSendGiftV2ReturnsEveryPackedGift(t *testing.T) {
	first := appendProtoVarint(nil, 1, 32126)
	first = appendProtoBytes(first, 2, []byte("棉花糖"))
	first = appendProtoVarint(first, 3, 7)
	second := appendProtoVarint(nil, 1, 32128)
	second = appendProtoBytes(second, 2, []byte("爱心抱枕"))
	second = appendProtoVarint(second, 3, 3)
	payload := appendProtoVarint(nil, 1, 629184597)
	payload = appendProtoBytes(payload, 2, []byte("柠汪汪"))
	payload = appendProtoBytes(payload, 10, first)
	payload = appendProtoBytes(payload, 10, second)

	gifts, err := parseSendGiftV2(base64.StdEncoding.EncodeToString(payload))
	if err != nil {
		t.Fatal(err)
	}
	if len(gifts) != 2 || gifts[0].GiftName != "棉花糖" || gifts[0].Num != 7 ||
		gifts[1].GiftName != "爱心抱枕" || gifts[1].Num != 3 {
		t.Fatalf("unexpected gifts: %#v", gifts)
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
