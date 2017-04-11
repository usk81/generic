package generic

import (
	"encoding/json"
	"strconv"
	"time"
)

// TimestampMS
type TimestampMS struct {
	ValidFlag
	Time time.Time
}

// Value returns Time.Time, but if Time.ValidFlag is false, returns nil.
func (v TimestampMS) Value() interface{} {
	if !v.Valid() {
		return nil
	}
	return v.Time
}

// Scan implements the sql.Scanner interface.
func (v *TimestampMS) Scan(x interface{}) (err error) {
	v.Time, v.ValidFlag, err = asTimestampMilliseconds(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

// Set sets a specified value.
func (v *TimestampMS) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// MarshalJSON implements the json.Marshaler interface.
func (v TimestampMS) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return json.Marshal(nil)
	}
	return []byte(strconv.FormatInt(v.Time.Unix(), 10)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *TimestampMS) UnmarshalJSON(data []byte) error {
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
