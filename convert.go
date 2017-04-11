package generic

import (
	"strconv"
	"strings"
	"time"
)

// asBool converts a specified value to boolean value.
func asBool(x interface{}) (result bool, isValid ValidFlag, err error) {
	switch x.(type) {
	case nil:
		return result, false, nil
	case int:
		result = x.(int) != 0
	case int8:
		result = x.(int8) != 0
	case int16:
		result = x.(int16) != 0
	case int32:
		result = x.(int32) != 0
	case int64:
		result = x.(int64) != 0
	case uint:
		result = x.(uint) != 0
	case uint8:
		result = x.(uint8) != 0
	case uint16:
		result = x.(uint16) != 0
	case uint32:
		result = x.(uint32) != 0
	case uint64:
		result = x.(uint64) != 0
	case float32:
		result = x.(float32) != 0
	case float64:
		result = x.(float64) != 0
	case bool:
		result = x.(bool)
	case string:
		b, err := strconv.ParseBool(x.(string))
		if err != nil {
			return result, false, ErrInvalidGenericValue{Value: x}
		}
		result = b
	default:
		return result, false, ErrInvalidGenericValue{Value: x}
	}
	return result, true, nil
}

// asBool converts a specified value to float64 value.
func asFloat(x interface{}) (result float64, isValid ValidFlag, err error) {
	switch x.(type) {
	case nil:
		return result, false, nil
	case int:
		result = float64(x.(int))
	case int8:
		result = float64(x.(int8))
	case int16:
		result = float64(x.(int16))
	case int32:
		result = float64(x.(int32))
	case int64:
		result = float64(x.(int64))
	case uint:
		result = float64(x.(uint))
	case uint8:
		result = float64(x.(uint8))
	case uint16:
		result = float64(x.(uint16))
	case uint32:
		result = float64(x.(uint32))
	case uint64:
		result = float64(x.(uint64))
	case float32:
		result = float64(x.(float32))
	case float64:
		result = x.(float64)
	case bool:
		b := x.(bool)
		if b {
			result = 1
		} else {
			result = 0
		}
	case string:
		f, err := strconv.ParseFloat(x.(string), 64)
		if err != nil {
			return result, false, ErrInvalidGenericValue{Value: x}
		}
		result = f
	default:
		return result, false, ErrInvalidGenericValue{Value: x}
	}
	return result, true, nil
}

// asBool converts a specified value to int64 value.
func asInt(x interface{}) (result int64, isValid ValidFlag, err error) {
	switch x.(type) {
	case nil:
		return result, false, nil
	case int:
		result = int64(x.(int))
	case int8:
		result = int64(x.(int8))
	case int16:
		result = int64(x.(int16))
	case int32:
		result = int64(x.(int32))
	case int64:
		result = x.(int64)
	case uint:
		result = int64(x.(uint))
	case uint8:
		result = int64(x.(uint8))
	case uint16:
		result = int64(x.(uint16))
	case uint32:
		result = int64(x.(uint32))
	case uint64:
		result = int64(x.(uint64))
	case float32:
		result = int64(x.(float32))
	case float64:
		result = int64(x.(float64))
	case bool:
		b := x.(bool)
		if b == true {
			result = 1
		} else {
			result = 0
		}
	case string:
		result, err = strconv.ParseInt(x.(string), 10, 64)
		if err != nil {
			return 0, false, ErrInvalidGenericValue{Value: x}
		}
	default:
		return result, false, ErrInvalidGenericValue{Value: x}
	}
	return result, true, nil
}

// asBool converts a specified value to string value.
func asString(x interface{}) (result string, isValid ValidFlag, err error) {
	switch x.(type) {
	case nil:
		return result, false, nil
	case int:
		result = strconv.FormatInt(int64(x.(int)), 10)
	case int8:
		result = strconv.FormatInt(int64(x.(int8)), 10)
	case int16:
		result = strconv.FormatInt(int64(x.(int16)), 10)
	case int32:
		result = strconv.FormatInt(int64(x.(int32)), 10)
	case int64:
		result = strconv.FormatInt(x.(int64), 10)
	case uint:
		result = strconv.FormatUint(uint64(x.(uint)), 10)
	case uint8:
		result = strconv.FormatUint(uint64(x.(uint8)), 10)
	case uint16:
		result = strconv.FormatUint(uint64(x.(uint16)), 10)
	case uint32:
		result = strconv.FormatUint(uint64(x.(uint32)), 10)
	case uint64:
		result = strconv.FormatUint(x.(uint64), 10)
	case float32:
		fs := strconv.FormatFloat(float64(x.(float32)), 'f', 10, 64)
		result = strings.TrimRight(strings.TrimRight(fs, "0"), ".")
	case float64:
		fs := strconv.FormatFloat(x.(float64), 'f', 10, 64)
		result = strings.TrimRight(strings.TrimRight(fs, "0"), ".")
	case bool:
		result = strconv.FormatBool(x.(bool))
	case string:
		result = x.(string)
	default:
		return result, false, ErrInvalidGenericValue{Value: x}
	}
	return result, true, nil
}

// asBool converts a specified value to time.Time value.
func asTime(x interface{}) (result time.Time, isValid ValidFlag, err error) {
	switch x.(type) {
	case nil:
		return result, false, nil
	case time.Time:
		result = x.(time.Time)
		if result.IsZero() {
			return result, false, nil
		}
	default:
		return result, false, ErrInvalidGenericValue{Value: x}
	}
	return result, true, nil
}

// asTimestamp converts a specified value to time.Time value.
func asTimestamp(x interface{}) (result time.Time, isValid ValidFlag, err error) {
	var i int64
	switch x.(type) {
	case nil:
		return result, false, nil
	case time.Time:
		result = x.(time.Time)
		if result.IsZero() {
			return result, false, nil
		}
		return result, true, nil
	case int:
		i = int64(x.(int))
	case int8:
		i = int64(x.(int8))
	case int16:
		i = int64(x.(int16))
	case int32:
		i = int64(x.(int32))
	case int64:
		i = x.(int64)
	case uint:
		i = int64(x.(uint))
	case uint8:
		i = int64(x.(uint8))
	case uint16:
		i = int64(x.(uint16))
	case uint32:
		i = int64(x.(uint32))
	case uint64:
		i = int64(x.(uint64))
	case float32:
		i = int64(x.(float32))
	case float64:
		i = int64(x.(float64))
	default:
		return result, false, ErrInvalidGenericValue{Value: x}
	}
	if i < 0 {
		return result, false, ErrInvalidGenericValue{Value: x}
	}
	return time.Unix(i, 0), true, nil
}

// asTimestampNanoseconds converts a specified value to time.Time value.
func asTimestampNanoseconds(x interface{}) (result time.Time, isValid ValidFlag, err error) {
	var i int64
	switch x.(type) {
	case nil:
		return result, false, nil
	case time.Time:
		result = x.(time.Time)
		if result.IsZero() {
			return result, false, nil
		}
		return result, true, nil
	case int:
		i = int64(x.(int))
	case int8:
		i = int64(x.(int8))
	case int16:
		i = int64(x.(int16))
	case int32:
		i = int64(x.(int32))
	case int64:
		i = x.(int64)
	case uint:
		i = int64(x.(uint))
	case uint8:
		i = int64(x.(uint8))
	case uint16:
		i = int64(x.(uint16))
	case uint32:
		i = int64(x.(uint32))
	case uint64:
		i = int64(x.(uint64))
	case float32:
		i = int64(x.(float32))
	case float64:
		i = int64(x.(float64))
	default:
		return result, false, ErrInvalidGenericValue{Value: x}
	}
	if i < 0 {
		return result, false, ErrInvalidGenericValue{Value: x}
	}
	return time.Unix(0, i), true, nil
}

// asTimestampMilliseconds converts a specified value to time.Time value.
func asTimestampMilliseconds(x interface{}) (result time.Time, isValid ValidFlag, err error) {
	var i int64
	switch x.(type) {
	case nil:
		return result, false, nil
	case time.Time:
		result = x.(time.Time)
		if result.IsZero() {
			return result, false, nil
		}
		return result, true, nil
	case int:
		i = int64(x.(int))
	case int8:
		i = int64(x.(int8))
	case int16:
		i = int64(x.(int16))
	case int32:
		i = int64(x.(int32))
	case int64:
		i = x.(int64)
	case uint:
		i = int64(x.(uint))
	case uint8:
		i = int64(x.(uint8))
	case uint16:
		i = int64(x.(uint16))
	case uint32:
		i = int64(x.(uint32))
	case uint64:
		i = int64(x.(uint64))
	case float32:
		i = int64(x.(float32))
	case float64:
		i = int64(x.(float64))
	default:
		return result, false, ErrInvalidGenericValue{Value: x}
	}
	if i < 0 {
		return result, false, ErrInvalidGenericValue{Value: x}
	}
	return time.Unix(0, i*1000000), true, nil
}

// asBool converts a specified value to uint64 value.
func asUint(x interface{}) (result uint64, isValid ValidFlag, err error) {
	switch x.(type) {
	case nil:
		return 0, false, nil
	case int:
		i := x.(int)
		if i < 0 {
			return result, false, ErrInvalidGenericValue{Value: x}
		}
		result = uint64(i)
	case int8:
		i8 := x.(int8)
		if i8 < 0 {
			return result, false, ErrInvalidGenericValue{Value: x}
		}
		result = uint64(i8)
	case int16:
		i16 := x.(int16)
		if i16 < 0 {
			return result, false, ErrInvalidGenericValue{Value: x}
		}
		result = uint64(i16)
	case int32:
		i32 := x.(int32)
		if i32 < 0 {
			return result, false, ErrInvalidGenericValue{Value: x}
		}
		result = uint64(i32)
	case int64:
		i64 := x.(int64)
		if i64 < 0 {
			return result, false, ErrInvalidGenericValue{Value: x}
		}
		result = uint64(i64)
	case uint:
		result = uint64(x.(uint))
	case uint8:
		result = uint64(x.(uint8))
	case uint16:
		result = uint64(x.(uint16))
	case uint32:
		result = uint64(x.(uint32))
	case uint64:
		result = x.(uint64)
	case float32:
		f32 := x.(float32)
		if f32 < 0 {
			return result, false, ErrInvalidGenericValue{Value: x}
		}
		result = uint64(f32)
	case float64:
		f64 := x.(float64)
		if f64 < 0 {
			return result, false, ErrInvalidGenericValue{Value: x}
		}
		result = uint64(f64)
	case bool:
		if x.(bool) {
			result = 1
		} else {
			result = 0
		}
	case string:
		u64, err := strconv.ParseUint(x.(string), 10, 64)
		if err != nil {
			return result, false, ErrInvalidGenericValue{Value: x}
		}
		result = u64
	default:
		return result, false, ErrInvalidGenericValue{Value: x}
	}
	return result, true, nil
}
