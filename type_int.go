package generic

import "encoding/json"

// Int
type Int struct {
	ValidFlag
	Int int64
}

// Value returns Int.Int, but if Int.ValidFlag is false, returns nil.
func (v Int) Value() interface{} {
	if !v.Valid() {
		return nil
	}
	return v.Int
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

// Set sets a specified value.
func (v *Int) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// MarshalJSON implements the json.Marshaler interface.
func (v Int) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Int)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *Int) UnmarshalJSON(data []byte) error {
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
