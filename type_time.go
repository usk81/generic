package generic

import (
	"database/sql/driver"
	"time"
)

// Time is generic time type structure
type Time struct {
	ValidFlag
	Time time.Time
}

// Value implements the driver Valuer interface.
func (v Time) Value() (driver.Value, error) {
	if !v.Valid() {
		return nil, nil
	}
	return v.Time, nil
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

// Weak returns Time.Time, but if Time.ValidFlag is false, returns nil.
func (v Time) Weak() interface{} {
	i, _ := v.Value()
	return i
}

// Set sets a specified value.
func (v *Time) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// MarshalJSON implements the json.Marshaler interface.
func (v Time) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return nullBytes, nil
	}
	return v.Time.MarshalJSON()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *Time) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}
	if err := v.Time.UnmarshalJSON(data); err != nil {
		return err
	}
	v.ValidFlag = true
	return nil
	// var in interface{}
	// if err := json.Unmarshal(data, &in); err != nil {
	// 	return err
	// }
	// return v.Scan(in)
}
