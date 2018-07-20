package generic

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

// Float is generic float type structure
type Float struct {
	ValidFlag
	float float64
}

// MarshalFloat return generic.Float converting of request data
func MarshalFloat(x interface{}) (Float, error) {
	v := Float{}
	err := v.Scan(x)
	return v, err
}

// Value implements the driver Valuer interface.
func (v Float) Value() (driver.Value, error) {
	if !v.Valid() {
		return nil, nil
	}
	return v.float, nil
}

// Scan implements the sql.Scanner interface.
func (v *Float) Scan(x interface{}) (err error) {
	v.float, v.ValidFlag, err = asFloat(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

// Weak returns Float.float, but if Float.ValidFlag is false, returns nil.
func (v Float) Weak() interface{} {
	i, _ := v.Value()
	return i
}

// Set sets a specified value.
func (v *Float) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// Float32 returns float32 value
func (v Float) Float32() float32 {
	return float32(v.Float64())
}

// Float64 returns float64 value
func (v Float) Float64() float64 {
	if !v.Valid() {
		return 0
	}
	return v.float
}

// String implements the Stringer interface.
func (v Float) String() string {
	if !v.Valid() {
		return ""
	}
	return strconv.FormatFloat(v.float, 'f', -1, 64)
}

// MarshalJSON implements the json.Marshaler interface.
func (v Float) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return nullBytes, nil
	}
	return []byte(strconv.FormatFloat(v.float, 'f', -1, 64)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *Float) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
