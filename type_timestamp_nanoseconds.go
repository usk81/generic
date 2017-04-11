package generic

import (
	"encoding/json"
	"strconv"
	"time"
)

// Timestamp
type TimestampNano struct {
	ValidFlag
	Time time.Time
}

// Value returns Time.Time, but if Time.ValidFlag is false, returns nil.
func (v TimestampNano) Value() interface{} {
	if !v.Valid() {
		return nil
	}
	return v.Time
}

// Scan implements the sql.Scanner interface.
func (v *TimestampNano) Scan(x interface{}) (err error) {
	v.Time, v.ValidFlag, err = asTimestampNanoseconds(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

// Set sets a specified value.
func (v *TimestampNano) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// MarshalJSON implements the json.Marshaler interface.
func (v TimestampNano) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return json.Marshal(nil)
	}
	return []byte(strconv.FormatInt(v.Time.UnixNano(), 10)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *TimestampNano) UnmarshalJSON(data []byte) error {
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
