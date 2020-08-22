package generic

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// asBool converts a specified value to boolean value.
func asBool(x interface{}) (result bool, isValid ValidFlag, err error) {
	switch t := x.(type) {
	case nil:
		return result, false, nil
	case int, int8, int16, int32, int64:
		result = reflect.ValueOf(t).Int() != 0
	case uint, uint8, uint16, uint32, uint64:
		result = reflect.ValueOf(t).Uint() != 0
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
	switch v := x.(type) {
	case nil:
		return result, false, nil
	case int:
		result = float64(v)
	case int8:
		result = float64(v)
	case int16:
		result = float64(v)
	case int32:
		result = float64(v)
	case int64:
		result = float64(v)
	case uint:
		result = float64(v)
	case uint8:
		result = float64(v)
	case uint16:
		result = float64(v)
	case uint32:
		result = float64(v)
	case uint64:
		result = float64(v)
	case float32:
		result = float64(v)
	case float64:
		result = v
	case bool:
		if v {
			result = 1
		} else {
			result = 0
		}
	case string:
		f, err := strconv.ParseFloat(v, 64)
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
	switch t := x.(type) {
	case nil:
		return result, false, nil
	case int, int8, int16, int32, int64:
		result = reflect.ValueOf(t).Int()
	case uint, uint8, uint16, uint32, uint64:
		result = int64(reflect.ValueOf(t).Uint())
	case float32:
		result = int64(x.(float32))
	case float64:
		result = int64(x.(float64))
	case bool:
		b := x.(bool)
		if b {
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
	switch t := x.(type) {
	case nil:
		return result, false, nil
	case int, int8, int16, int32, int64:
		result = strconv.FormatInt(reflect.ValueOf(t).Int(), 10)
	case uint, uint8, uint16, uint32, uint64:
		result = strconv.FormatUint(reflect.ValueOf(t).Uint(), 10)
	case float32, float64:
		fs := strconv.FormatFloat(reflect.ValueOf(t).Float(), 'f', 10, 64)
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
	switch v := x.(type) {
	case nil:
		return result, false, nil
	case time.Time:
		result = v
		if result.IsZero() {
			return result, true, nil
		}
	default:
		return result, false, ErrInvalidGenericValue{Value: x}
	}
	return result, true, nil
}

// asTimestamp converts a specified value to time.Time value.
func asTimestamp(x interface{}) (result time.Time, isValid ValidFlag, err error) {
	return asTimestampWithFunc(x, func(i int64) time.Time {
		return time.Unix(i, 0)
	})
}

// asTimestampNanoseconds converts a specified value to time.Time value.
func asTimestampNanoseconds(x interface{}) (result time.Time, isValid ValidFlag, err error) {
	return asTimestampWithFunc(x, func(i int64) time.Time {
		return time.Unix(0, i)
	})
}

// asTimestampMilliseconds converts a specified value to time.Time value.
func asTimestampMilliseconds(x interface{}) (result time.Time, isValid ValidFlag, err error) {
	return asTimestampWithFunc(x, func(i int64) time.Time {
		return time.Unix(0, i*1000000)
	})
}

// asBool converts a specified value to uint64 value.
func asUint(x interface{}) (result uint64, isValid ValidFlag, err error) {
	switch t := x.(type) {
	case nil:
		return 0, false, nil
	case int, int8, int16, int32, int64:
		i := reflect.ValueOf(t).Int()
		if i < 0 {
			return result, false, ErrInvalidGenericValue{Value: x}
		}
		result = uint64(i)
	case uint, uint8, uint16, uint32, uint64:
		result = reflect.ValueOf(t).Uint()
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

func asTimestampWithFunc(x interface{}, f func(i int64) time.Time) (result time.Time, isValid ValidFlag, err error) {
	var i int64
	switch t := x.(type) {
	case nil:
		return result, false, nil
	case time.Time:
		result = x.(time.Time)
		if result.IsZero() {
			return result, true, nil
		}
		return result, true, nil
	case string:
		result, err = time.Parse(time.RFC3339Nano, x.(string))
		return result, err == nil, err
	case int, int8, int16, int32, int64:
		i = reflect.ValueOf(t).Int()
	case uint, uint8, uint16, uint32, uint64:
		i = int64(reflect.ValueOf(t).Uint())
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
	return f(i), true, nil
}

func asURL(x interface{}) (result *url.URL, isValid ValidFlag, err error) {
	switch v := x.(type) {
	case nil:
		return nil, false, nil
	case *url.URL:
		result = v
	case string:
		result, err = url.Parse(v)
	default:
		err = ErrInvalidGenericValue{Value: x}
	}
	return result, (err == nil), err
}
