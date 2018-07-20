package generic

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

// Int is generic integer type structure
type Int struct {
	ValidFlag
	int int64
}

// MarshalInt return generic.Int converting of request data
func MarshalInt(x interface{}) (Int, error) {
	v := Int{}
	err := v.Scan(x)
	return v, err
}

// Value implements the driver Valuer interface.
func (v Int) Value() (driver.Value, error) {
	if !v.Valid() {
		return nil, nil
	}
	return v.int, nil
}

// Scan implements the sql.Scanner interface.
func (v *Int) Scan(x interface{}) (err error) {
	v.int, v.ValidFlag, err = asInt(x)
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

// Int return int value
func (v Int) Int() int {
	if !v.Valid() {
		return 0
	}
	return int(v.int)
}

// Int32 return int32 value
func (v Int) Int32() int32 {
	if !v.Valid() {
		return 0
	}
	return int32(v.int)
}

// Int64 return int64 value
func (v Int) Int64() int64 {
	if !v.Valid() {
		return 0
	}
	return v.int
}

// String implements the Stringer interface.
func (v Int) String() string {
	if !v.Valid() {
		return ""
	}
	return strconv.FormatInt(v.int, 10)
}

// MarshalJSON implements the json.Marshaler interface.
func (v Int) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return nullBytes, nil
	}
	return []byte(strconv.FormatInt(v.int, 10)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *Int) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
