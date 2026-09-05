package structs

import (
	"encoding/base64"
	"errors"
	"fmt"
)

type sendGiftV2 struct {
	UID      int64
	UserName string
	GiftID   int64
	GiftName string
	Num      int64
}

// parseSendGiftV2 decodes the compact protobuf payload carried in data.pb.
// Bilibili introduced SEND_GIFT_V2 without exposing the gift fields in JSON.
func parseSendGiftV2(encoded string) ([]sendGiftV2, error) {
	payload, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("decode data.pb: %w", err)
	}

	var uid int64
	var username string
	var giftPayloads [][]byte
	if err := walkProto(payload, func(field int, wireType byte, value uint64, data []byte) error {
		switch {
		case field == 1 && wireType == 0:
			uid = int64(value)
		case field == 2 && wireType == 2:
			username = string(data)
		case field == 10 && wireType == 2:
			giftPayloads = append(giftPayloads, data)
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("decode gift protobuf: %w", err)
	}
	if uid <= 0 || username == "" || len(giftPayloads) == 0 {
		return nil, errors.New("gift protobuf is missing required fields")
	}

	gifts := make([]sendGiftV2, 0, len(giftPayloads))
	for _, data := range giftPayloads {
		gift := sendGiftV2{UID: uid, UserName: username}
		if err := walkProto(data, func(field int, wireType byte, value uint64, data []byte) error {
			switch {
			case field == 1 && wireType == 0:
				gift.GiftID = int64(value)
			case field == 2 && wireType == 2:
				gift.GiftName = string(data)
			case field == 3 && wireType == 0:
				gift.Num = int64(value)
			}
			return nil
		}); err != nil {
			return nil, fmt.Errorf("decode gift protobuf: %w", err)
		}
		if gift.GiftID <= 0 || gift.GiftName == "" || gift.Num <= 0 {
			return nil, errors.New("gift protobuf is missing required fields")
		}
		gifts = append(gifts, gift)
	}
	return gifts, nil
}

func walkProto(data []byte, visit func(field int, wireType byte, value uint64, data []byte) error) error {
	for len(data) > 0 {
		tag, n := consumeUvarint(data)
		if n <= 0 {
			return errors.New("invalid protobuf tag")
		}
		data = data[n:]
		field, wireType := int(tag>>3), byte(tag&7)
		if field <= 0 {
			return errors.New("invalid protobuf field number")
		}

		var value uint64
		var fieldData []byte
		switch wireType {
		case 0:
			value, n = consumeUvarint(data)
			if n <= 0 {
				return errors.New("invalid protobuf varint")
			}
			data = data[n:]
		case 1:
			if len(data) < 8 {
				return errors.New("truncated protobuf fixed64")
			}
			fieldData, data = data[:8], data[8:]
		case 2:
			length, lengthBytes := consumeUvarint(data)
			if lengthBytes <= 0 {
				return errors.New("invalid protobuf length")
			}
			data = data[lengthBytes:]
			if length > uint64(len(data)) {
				return errors.New("truncated protobuf bytes")
			}
			fieldData, data = data[:int(length)], data[int(length):]
		case 5:
			if len(data) < 4 {
				return errors.New("truncated protobuf fixed32")
			}
			fieldData, data = data[:4], data[4:]
		default:
			return fmt.Errorf("unsupported protobuf wire type %d", wireType)
		}
		if err := visit(field, wireType, value, fieldData); err != nil {
			return err
		}
	}
	return nil
}

func consumeUvarint(data []byte) (uint64, int) {
	var value uint64
	for i, current := range data {
		if i == 10 || (i == 9 && current > 1) {
			return 0, -1
		}
		value |= uint64(current&0x7f) << (7 * i)
		if current < 0x80 {
			return value, i + 1
		}
	}
	return 0, 0
}
