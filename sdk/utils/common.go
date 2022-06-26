package utils

import "encoding/hex"
import "github.com/apache/thrift/lib/go/thrift"

func PanicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}

func Hex(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

func GetInt64(i interface{}) *int64 {
	switch v := i.(type) {
	case int:
		return thrift.Int64Ptr(int64(v))
	case int8:
		return thrift.Int64Ptr(int64(v))
	case int16:
		return thrift.Int64Ptr(int64(v))
	case int32:
		return thrift.Int64Ptr(int64(v))
	case int64:
		return thrift.Int64Ptr(int64(v))
	case uint:
		return thrift.Int64Ptr(int64(v))
	case uint8:
		return thrift.Int64Ptr(int64(v))
	case uint16:
		return thrift.Int64Ptr(int64(v))
	case uint32:
		return thrift.Int64Ptr(int64(v))
	case uint64:
		return thrift.Int64Ptr(int64(v))
	case float32:
		return thrift.Int64Ptr(int64(v))
	case float64:
		return thrift.Int64Ptr(int64(v))
	default:
		return nil
	}
}

func GetString(i interface{}) *string {
	switch v := i.(type) {
	case string:
		return &v
	default:
		return nil
	}
}
