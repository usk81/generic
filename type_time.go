package generic

import (
	"encoding/json"
	"time"
)

// Time
type Time struct {
	ValidFlag
	Time time.Time
}

// Value returns Time.Time, but if Time.ValidFlag is false, returns nil.
func (v Time) Value() interface{} {
	if !v.Valid() {
		return nil
	}
	return v.Time
}

// Scan implements the sql.Scanner interface.
func (v *Time) Scan(x interface{}) (err error) {
	v.Time, v.ValidFlag, err = asTime(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

// Set sets a specified value.
func (v *Time) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// MarshalJSON implements the json.Marshaler interface.
func (v Time) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Time)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *Time) UnmarshalJSON(data []byte) error {
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
