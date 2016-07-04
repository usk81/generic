package generic

import (
	"encoding/json"
	"time"
)

// TypeTime
type TypeTime struct {
	ValidFlag
	Time time.Time
}

func (v TypeTime) Value() interface{} {
	if !v.Valid() {
		return nil
	}
	return v.Time
}

// Scan implements the sql.Scanner interface.
func (v *TypeTime) Scan(x interface{}) (err error) {
	v.Time, v.ValidFlag, err = asTime(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

func (v *TypeTime) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// MarshalJSON implements the json.Marshaler interface.
func (v TypeTime) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Time)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *TypeTime) UnmarshalJSON(data []byte) error {
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
