package helper

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

// Convert integers to string, return empty string if the conversion fails
func Int2Str(val interface{}) string {
	switch val.(type) {
	// Integers
	case int:
		return fmt.Sprintf("%d", val)
	case int8:
		return fmt.Sprintf("%d", val)
	case int16:
		return fmt.Sprintf("%d", val)
	case int32:
		return fmt.Sprintf("%d", val)
	case int64:
		return fmt.Sprintf("%d", val)
	case uint:
		return fmt.Sprintf("%d", val)
	case uint8:
		return fmt.Sprintf("%d", val)
	case uint16:
		return fmt.Sprintf("%d", val)
	case uint32:
		return fmt.Sprintf("%d", val)
	case uint64:
		return fmt.Sprintf("%d", val)
	// Type is not integers, return empty string
	default:
		return ""
	}

}

// Convert floats to string, parameter length specifies the number of digits behind the float point, return empty string if the conversion fails
func Float2Str(val interface{}, length int) string {
	switch val.(type) {
	// Floats
	case float32:
		return fmt.Sprintf("%."+Int2Str(length)+"f", val)
	case float64:
		return fmt.Sprintf("%."+Int2Str(length)+"f", val)
		// Type is not floats, return empty string
	default:
		return ""
	}
}

// Convert bool to string
func Bool2Str(val bool) string {
	if val {
		return "true"
	}
	return "false"
}

func internalStr2Int(val string, bitSize int, noPanic bool) (int64, error) {
	res, err := strconv.ParseInt(val, 0, bitSize)
	if err != nil {
		if noPanic {
			return 0, err
		} else {
			panic(err)
		}
	}
	return res, nil
}

// Simply return the conversion result.
// For conversions that may panic, the code won't come here.
// For conversions that won't panic, res will be zero value and err will come here as well.
func wrapInt64Result(res int64, err error) int64 {
	return res
}

// Convert string to int64, panic when the string is not of a valid int64 format.
func Str2Int64(val string) int64 {
	return int64(wrapInt64Result(internalStr2Int(val, 64, false)))
}

// Convert string to int64, return 0 if the conversion fails
func Str2Int64NoPanic(val string) int64 {
	return int64(wrapInt64Result(internalStr2Int(val, 64, true)))
}

// Convert string to int64, return the default value you set when the conversion fails
func Str2Int64WithDefaultValue(val string, defaultVal int64) int64 {
	if res, err := internalStr2Int(val, 64, true); err != nil {
		return defaultVal
	} else {
		return int64(res)
	}
}

// Convert string to int32, panic when the string is not of a valid int32 format.
func Str2Int32(val string) int32 {
	return int32(wrapInt64Result(internalStr2Int(val, 32, false)))
}

// Convert string to int32, return 0 if the conversion fails
func Str2Int32NoPanic(val string) int32 {
	return int32(wrapInt64Result(internalStr2Int(val, 32, true)))
}

// Convert string to int32, return the default value you set when the conversion fails
func Str2Int32WithDefaultValue(val string, defaultVal int32) int32 {
	if res, err := internalStr2Int(val, 32, true); err != nil {
		return defaultVal
	} else {
		return int32(res)
	}
}

// Convert string to int16, panic when the string is not of a valid int16 format.
func Str2Int16(val string) int16 {
	return int16(wrapInt64Result(internalStr2Int(val, 16, false)))
}

// Convert string to int16, return 0 if the conversion fails
func Str2Int16NoPanic(val string) int16 {
	return int16(wrapInt64Result(internalStr2Int(val, 16, true)))
}

// Convert string to int16, return the default value you set when the conversion fails
func Str2Int16WithDefaultValue(val string, defaultVal int16) int16 {
	if res, err := internalStr2Int(val, 16, true); err != nil {
		return defaultVal
	} else {
		return int16(res)
	}
}

// Convert string to int8, panic when the string is not of a valid int8 format.
func Str2Int8(val string) int8 {
	return int8(wrapInt64Result(internalStr2Int(val, 8, false)))
}

// Convert string to int8, return 0 if the conversion fails
func Str2Int8NoPanic(val string) int8 {
	return int8(wrapInt64Result(internalStr2Int(val, 8, true)))
}

// Convert string to int8, return the default value you set when the conversion fails
func Str2Int8WithDefaultValue(val string, defaultVal int8) int8 {
	if res, err := internalStr2Int(val, 8, true); err != nil {
		return defaultVal
	} else {
		return int8(res)
	}
}

// Convert string to int, panic when the string is not of a valid int format.
func Str2Int(val string) int {
	return int(wrapInt64Result(internalStr2Int(val, 0, false)))
}

// Convert string to int, return 0 if the conversion fails
func Str2IntNoPanic(val string) int {
	return int(wrapInt64Result(internalStr2Int(val, 0, true)))
}

// Convert string to int, return the default value you set when the conversion fails
func Str2IntWithDefaultValue(val string, defaultVal int) int {
	if res, err := internalStr2Int(val, 0, true); err != nil {
		return defaultVal
	} else {
		return int(res)
	}
}

func internalStr2Uint(val string, bitSize int, noPanic bool) (uint64, error) {
	res, err := strconv.ParseUint(val, 0, bitSize)
	if err != nil {
		if noPanic {
			return 0, err
		} else {
			panic(err)
		}
	}
	return res, nil
}

// Simply return the conversion result.
// For conversions that may panic, the code won't come here.
// For conversions that won't panic, res will be zero value and err will come here as well.
func wrapUint64Result(res uint64, err error) uint64 {
	return res
}

// Convert string to uint64, panic when the string is not of a valid uint64 format.
func Str2Uint64(val string) uint64 {
	return uint64(wrapUint64Result(internalStr2Uint(val, 64, false)))
}

// Convert string to uint64, return 0 if the conversion fails
func Str2Uint64NoPanic(val string) uint64 {
	return uint64(wrapUint64Result(internalStr2Uint(val, 64, true)))
}

// Convert string to uint64, return the default value you set when the conversion fails
func Str2Uint64WithDefaultValue(val string, defaultVal uint64) uint64 {
	if res, err := internalStr2Uint(val, 64, true); err != nil {
		return defaultVal
	} else {
		return uint64(res)
	}
}

// Convert string to uint32, panic when the string is not of a valid uint32 format.
func Str2Uint32(val string) uint32 {
	return uint32(wrapUint64Result(internalStr2Uint(val, 32, false)))
}

// Convert string to uint32, return 0 if the conversion fails
func Str2Uint32NoPanic(val string) uint32 {
	return uint32(wrapUint64Result(internalStr2Uint(val, 32, true)))
}

// Convert string to uint32, return the default value you set when the conversion fails
func Str2Uint32WithDefaultValue(val string, defaultVal uint32) uint32 {
	if res, err := internalStr2Uint(val, 32, true); err != nil {
		return defaultVal
	} else {
		return uint32(res)
	}
}

// Convert string to uint16, panic when the string is not of a valid uint16 format.
func Str2Uint16(val string) uint16 {
	return uint16(wrapUint64Result(internalStr2Uint(val, 16, false)))
}

// Convert string to uint16, return 0 if the conversion fails
func Str2Uint16NoPanic(val string) uint16 {
	return uint16(wrapUint64Result(internalStr2Uint(val, 16, true)))
}

// Convert string to uint16, return the default value you set when the conversion fails
func Str2Uint16WithDefaultValue(val string, defaultVal uint16) uint16 {
	if res, err := internalStr2Uint(val, 16, true); err != nil {
		return defaultVal
	} else {
		return uint16(res)
	}
}

// Convert string to uint8, panic when the string is not of a valid uint8 format.
func Str2Uint8(val string) uint8 {
	return uint8(wrapUint64Result(internalStr2Uint(val, 8, false)))
}

// Convert string to uint8, return 0 if the conversion fails
func Str2Uint8NoPanic(val string) uint8 {
	return uint8(wrapUint64Result(internalStr2Uint(val, 8, true)))
}

// Convert string to uint16, return the default value you set when the conversion fails
func Str2Uint8WithDefaultValue(val string, defaultVal uint8) uint8 {
	if res, err := internalStr2Uint(val, 8, true); err != nil {
		return defaultVal
	} else {
		return uint8(res)
	}
}

// Convert string to uint, panic when the string is not of a valid uint format.
func Str2Uint(val string) uint {
	return uint(wrapUint64Result(internalStr2Uint(val, 0, false)))
}

// Convert string to uint, return 0 if the conversion fails
func Str2UintNoPanic(val string) uint {
	return uint(wrapUint64Result(internalStr2Uint(val, 0, true)))
}

// Convert string to uint, return the default value you set when the conversion fails
func Str2UintWithDefaultValue(val string, defaultVal uint) uint {
	if res, err := internalStr2Uint(val, 0, true); err != nil {
		return defaultVal
	} else {
		return uint(res)
	}
}

func internalStr2Float(val string, bitSize int, noPanic bool) (float64, error) {
	res, err := strconv.ParseFloat(val, bitSize)
	if err != nil {
		if noPanic {
			return 0, err
		} else {
			panic(err)
		}
	}
	return res, nil
}

// Simply return the conversion result.
// For conversions that may panic, the code won't come here.
// For conversions that won't panic, res will be zero value and err will come here as well.
func wrapFloat64Result(res float64, err error) float64 {
	return res
}

// Convert string to float64, panic when the string is not of a valid float64 format.
func Str2Float64(val string) float64 {
	return float64(wrapFloat64Result(internalStr2Float(val, 64, false)))
}

// Convert string to float64, return 0 if the conversion fails
func Str2Float64NoPanic(val string) float64 {
	return float64(wrapFloat64Result(internalStr2Float(val, 64, true)))
}

// Convert string to float64, return the default value you set when the conversion fails
func Str2Float64WithDefaultValue(val string, defaultVal float64) float64 {
	if res, err := internalStr2Float(val, 64, true); err != nil {
		return defaultVal
	} else {
		return float64(res)
	}
}

// Convert string to float32, panic when the string is not of a valid float32 format.
func Str2Float32(val string) float32 {
	return float32(wrapFloat64Result(internalStr2Float(val, 32, false)))
}

// Convert string to float32, return 0 if the conversion fails
func Str2Float32NoPanic(val string) float32 {
	return float32(wrapFloat64Result(internalStr2Float(val, 32, true)))
}

// Convert string to float32, return the default value you set when the conversion fails
func Str2Float32WithDefaultValue(val string, defaultVal float32) float32 {
	if res, err := internalStr2Float(val, 32, true); err != nil {
		return defaultVal
	} else {
		return float32(res)
	}
}

// Convert string to bool, panic when the string is not of a valid bool format
func Str2Bool(val string) bool {
	if val == "true" || val == "True" || val == "TRUE" {
		return true
	}
	if val == "false" || val == "False" || val == "FALSE" {
		return false
	}
	return false
}

// Convert bool to int, never panics
func Bool2Int(val bool) int {
	if val == true {
		return 1
	} else {
		return 0
	}
}

// Convert different types to byte slice using types and functions in unsafe and reflect package.
// It has higher performance, but notice that it may be not safe when garbage collection happens.
// Use it when you need to temporary convert a long string to a byte slice and won't keep it for long time.
func Str2ByteSliceNonCopy(val string) []byte {
	pSliceHeader := &reflect.SliceHeader{}
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&val))
	pSliceHeader.Data = strHeader.Data
	pSliceHeader.Len = strHeader.Len
	pSliceHeader.Cap = strHeader.Len
	return *(*[]byte)(unsafe.Pointer(pSliceHeader))
}

// Zero-copy convert from byte slice to a string
func BytesSlice2StrNonCopy(val []byte) string {
	return *(*string)(unsafe.Pointer(&val))
}
