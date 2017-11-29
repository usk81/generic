package generic

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

// Int is generic integer type structure
type Int struct {
	ValidFlag
	Int int64
}

// Value implements the driver Valuer interface.
func (v Int) Value() (driver.Value, error) {
	if !v.Valid() {
		return nil, nil
	}
	return v.Int, nil
}

// Scan implements the sql.Scanner interface.
func (v *Int) Scan(x interface{}) (err error) {
	v.Int, v.ValidFlag, err = asInt(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

// Weak returns Int.Int, but if Int.ValidFlag is false, returns nil.
func (v Int) Weak() interface{} {
	i, _ := v.Value()
	return i
}

// Set sets a specified value.
func (v *Int) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// MarshalJSON implements the json.Marshaler interface.
func (v Int) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return nullBytes, nil
	}
	return []byte(strconv.FormatInt(v.Int, 10)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *Int) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
