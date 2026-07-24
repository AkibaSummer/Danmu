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
func parseSendGiftV2(encoded string) (sendGiftV2, error) {
	payload, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return sendGiftV2{}, fmt.Errorf("decode data.pb: %w", err)
	}

	var gift sendGiftV2
	err = walkProto(payload, func(field int, wireType byte, value uint64, data []byte) error {
		switch {
		case field == 1 && wireType == 0:
			gift.UID = int64(value)
		case field == 2 && wireType == 2:
			gift.UserName = string(data)
		case field == 10 && wireType == 2:
			return walkProto(data, func(giftField int, giftWireType byte, giftValue uint64, giftData []byte) error {
				switch {
				case giftField == 1 && giftWireType == 0:
					gift.GiftID = int64(giftValue)
				case giftField == 2 && giftWireType == 2:
					gift.GiftName = string(giftData)
				case giftField == 3 && giftWireType == 0:
					gift.Num = int64(giftValue)
				}
				return nil
			})
		}
		return nil
	})
	if err != nil {
		return sendGiftV2{}, fmt.Errorf("decode gift protobuf: %w", err)
	}
	if gift.UID <= 0 || gift.UserName == "" || gift.GiftID <= 0 || gift.GiftName == "" || gift.Num <= 0 {
		return sendGiftV2{}, errors.New("gift protobuf is missing required fields")
	}
	return gift, nil
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
