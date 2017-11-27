package generic

import (
	"database/sql/driver"
	"encoding/json"
)

// String is generic string type structure
type String struct {
	ValidFlag
	String string
}

// Value implements the driver Valuer interface.
func (v String) Value() (driver.Value, error) {
	if !v.Valid() {
		return nil, nil
	}
	return v.String, nil
}

// Scan implements the sql.Scanner interface.
func (v *String) Scan(x interface{}) (err error) {
	v.String, v.ValidFlag, err = asString(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

// Weak returns string, but if String.ValidFlag is false, returns nil.
func (v String) Weak() interface{} {
	i, _ := v.Value()
	return i
}

// Set sets a specified value.
func (v *String) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// MarshalJSON implements the json.Marshaler interface.
func (v String) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return nullBytes, nil
	}
	s := `"` + v.String + `"`
	bs := make([]byte, 0, len(s))
	return append(bs, s...), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *String) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
