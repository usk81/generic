package generic

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

// Uint is generic uint type structure
type Uint struct {
	ValidFlag
	uint uint64
}

// Value implements the driver Valuer interface.
func (v Uint) Value() (driver.Value, error) {
	if !v.Valid() {
		return nil, nil
	}
	return v.uint, nil
}

// Scan implements the sql.Scanner interface.
func (v *Uint) Scan(x interface{}) (err error) {
	v.uint, v.ValidFlag, err = asUint(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

// Weak returns Uint.Uint, but if Uint.ValidFlag is false, returns nil.
func (v Uint) Weak() interface{} {
	i, _ := v.Value()
	return i
}

// Set sets a specified value.
func (v *Uint) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// Uint return uint value
func (v Uint) Uint() uint {
	if !v.Valid() {
		return 0
	}
	return uint(v.uint)
}

// Uint32 return uint32 value
func (v Uint) Uint32() uint32 {
	if !v.Valid() {
		return 0
	}
	return uint32(v.uint)
}

// Uint64 return uint64 value
func (v Uint) Uint64() uint64 {
	if !v.Valid() {
		return 0
	}
	return v.uint
}

// String implements the Stringer interface.
func (v Uint) String() string {
	if !v.Valid() {
		return ""
	}
	return strconv.FormatUint(v.uint, 10)
}

// MarshalJSON implements the json.Marshaler interface.
func (v Uint) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return nullBytes, nil
	}
	return []byte(strconv.FormatUint(v.uint, 10)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *Uint) UnmarshalJSON(data []byte) error {
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
