package utils

import (
	"reflect"
	"strconv"
)

type Decoder struct {
}

var defaultDecoder *Decoder

func DefaultDecoder() *Decoder {
	if defaultDecoder == nil {
		defaultDecoder = &Decoder{}
	}
	return defaultDecoder
}

func (d *Decoder) Decode(v any) []byte {
	rType := reflect.TypeOf(v).Elem()
	for i := 0; i < rType.NumField(); i++ {
		sLen := rType.Field(i).Tag.Get("len")
		len, err := strconv.Atoi(sLen)
		if err != nil {
			panic(err)
		}
		_ = len
	}
	return nil
}
