package generic

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"time"
)

// TimestampNano is a wrapped time type structure
type TimestampNano struct {
	ValidFlag
	time time.Time
}

// Value returns timestamp with nanoseconds, but if TimestampNano.ValidFlag is false, returns nil.
func (v TimestampNano) Value() (driver.Value, error) {
	if !v.Valid() {
		return nil, nil
	}
	return v.time.UnixNano(), nil
}

// Scan implements the sql.Scanner interface.
func (v *TimestampNano) Scan(x interface{}) (err error) {
	v.time, v.ValidFlag, err = asTimestampNanoseconds(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

// Weak returns timestamp with nano seconds, but if TimestampNano.ValidFlag is false, returns nil.
func (v TimestampNano) Weak() interface{} {
	i, _ := v.Value()
	return i
}

// Set sets a specified value.
func (v *TimestampNano) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// String implements the Stringer interface.
func (v TimestampNano) String() string {
	return strconv.FormatInt(v.Int64(), 10)
}

// Int return int value
func (v TimestampNano) Int() int {
	return int(v.Int64())
}

// Int64 return int64 value
func (v TimestampNano) Int64() int64 {
	if !v.Valid() || v.time.UnixNano() == 0 {
		return 0
	}
	return v.time.UnixNano()
}

// MarshalJSON implements the json.Marshaler interface.
func (v TimestampNano) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return nullBytes, nil
	}
	return []byte(strconv.FormatInt(v.time.UnixNano(), 10)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *TimestampNano) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
