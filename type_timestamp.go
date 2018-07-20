package generic

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"time"
)

// Timestamp is a wrapped time type structure
type Timestamp struct {
	ValidFlag
	time time.Time
}

// MarshalTimestamp return generic.Timestamp converting of request data
func MarshalTimestamp(x interface{}) (Timestamp, error) {
	v := Timestamp{}
	err := v.Scan(x)
	return v, err
}

// Value returns Time.Time, but if Time.ValidFlag is false, returns nil.
func (v Timestamp) Value() (driver.Value, error) {
	if !v.Valid() {
		return nil, nil
	}
	return v.time, nil
}

// Scan implements the sql.Scanner interface.
func (v *Timestamp) Scan(x interface{}) (err error) {
	v.time, v.ValidFlag, err = asTimestamp(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

// Weak returns timestamp, but if Timestamp.ValidFlag is false, returns nil.
func (v Timestamp) Weak() interface{} {
	i, _ := v.Value()
	return i
}

// Set sets a specified value.
func (v *Timestamp) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// String implements the Stringer interface.
func (v Timestamp) String() string {
	return strconv.FormatInt(v.Int64(), 10)
}

// Int return int value
func (v Timestamp) Int() int {
	return int(v.Int64())
}

// Int64 return int64 value
func (v Timestamp) Int64() int64 {
	if !v.Valid() || v.time.Unix() == 0 {
		return 0
	}
	return v.time.Unix()
}

// MarshalJSON implements the json.Marshaler interface.
func (v Timestamp) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return nullBytes, nil
	}
	return []byte(strconv.FormatInt(v.time.Unix(), 10)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *Timestamp) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
