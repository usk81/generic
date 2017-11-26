package generic

import (
	"encoding/json"
	"strconv"
	"time"
)

// Timestamp is a wrapped time type structure
type Timestamp struct {
	ValidFlag
	Time time.Time
}

// Value returns Time.Time, but if Time.ValidFlag is false, returns nil.
func (v Timestamp) Value() interface{} {
	if !v.Valid() {
		return nil
	}
	return v.Time
}

// Scan implements the sql.Scanner interface.
func (v *Timestamp) Scan(x interface{}) (err error) {
	v.Time, v.ValidFlag, err = asTimestamp(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

// Set sets a specified value.
func (v *Timestamp) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// MarshalJSON implements the json.Marshaler interface.
func (v Timestamp) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return nullBytes, nil
	}
	return []byte(strconv.FormatInt(v.Time.Unix(), 10)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *Timestamp) UnmarshalJSON(data []byte) error {
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
