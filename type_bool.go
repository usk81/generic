package generic

import (
	"database/sql/driver"
	"encoding/json"
)

// Bool is generic boolean type structure
type Bool struct {
	ValidFlag
	Bool bool
}

// Value implements the driver Valuer interface.
func (v Bool) Value() (driver.Value, error) {
	if !v.Valid() {
		return nil, nil
	}
	return v.Bool, nil
}

// Scan implements the sql.Scanner interface.
func (v *Bool) Scan(x interface{}) (err error) {
	v.Bool, v.ValidFlag, err = asBool(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

// Weak returns Bool.Bool, but if Bool.ValidFlag is false, returns nil.
func (v Bool) Weak() interface{} {
	i, _ := v.Value()
	return i
}

// Set sets a specified value.
func (v *Bool) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// MarshalJSON implements the json.Marshaler interface.
func (v Bool) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return nullBytes, nil
	}
	if v.Bool {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *Bool) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
