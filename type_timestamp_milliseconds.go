package generic

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"time"
)

// TimestampMS is a wrapped time type structure
type TimestampMS struct {
	ValidFlag
	time time.Time
}

// MarshalTimestampMS return generic.TimestampMS converting of request data
func MarshalTimestampMS(x interface{}) (TimestampMS, error) {
	v := TimestampMS{}
	err := v.Scan(x)
	return v, err
}

// Value returns timestamp with milliseconds, but if TimestampMS.ValidFlag is false, returns nil.
func (v TimestampMS) Value() (driver.Value, error) {
	if !v.Valid() {
		return nil, nil
	}
	return v.time.UnixNano() / 1000000, nil
}

// Scan implements the sql.Scanner interface.
func (v *TimestampMS) Scan(x interface{}) (err error) {
	v.time, v.ValidFlag, err = asTimestampMilliseconds(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

// Weak returns timestamp int value, but if TimestampMS.ValidFlag is false, returns nil.
func (v TimestampMS) Weak() interface{} {
	i, _ := v.Value()
	return i
}

// Set sets a specified value.
func (v *TimestampMS) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// String implements the Stringer interface.
func (v TimestampMS) String() string {
	return strconv.FormatInt(v.Int64(), 10)
}

// Int return int value
func (v TimestampMS) Int() int {
	return int(v.Int64())
}

// Int64 return int64 value
func (v TimestampMS) Int64() int64 {
	if !v.Valid() || v.time.UnixNano() == 0 {
		return 0
	}
	return v.time.UnixNano() / 1000000
}

// MarshalJSON implements the json.Marshaler interface.
func (v TimestampMS) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return nullBytes, nil
	}
	return []byte(v.String()), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *TimestampMS) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
