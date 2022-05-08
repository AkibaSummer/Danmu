package utils

import "encoding/hex"

func PanicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}

func Hex(bytes []byte) string {
	return hex.EncodeToString(bytes)
}
