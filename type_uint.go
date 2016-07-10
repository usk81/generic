package generic

import "encoding/json"

// TypeUint
type TypeUint struct {
	ValidFlag
	Uint uint64
}

// Value returns TypeUint.Uint, but if TypeUint.ValidFlag is false, returns nil.
func (v TypeUint) Value() interface{} {
	if !v.Valid() {
		return nil
	}
	return v.Uint
}

// Scan implements the sql.Scanner interface.
func (v *TypeUint) Scan(x interface{}) (err error) {
	v.Uint, v.ValidFlag, err = asUint(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

// Set sets a specified value.
func (v *TypeUint) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// MarshalJSON implements the json.Marshaler interface.
func (v TypeUint) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Uint)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *TypeUint) UnmarshalJSON(data []byte) error {
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
