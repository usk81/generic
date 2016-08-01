package generic

import "encoding/json"

// Uint
type Uint struct {
	ValidFlag
	Uint uint64
}

// Value returns Uint.Uint, but if Uint.ValidFlag is false, returns nil.
func (v Uint) Value() interface{} {
	if !v.Valid() {
		return nil
	}
	return v.Uint
}

// Scan implements the sql.Scanner interface.
func (v *Uint) Scan(x interface{}) (err error) {
	v.Uint, v.ValidFlag, err = asUint(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

// Set sets a specified value.
func (v *Uint) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// MarshalJSON implements the json.Marshaler interface.
func (v Uint) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Uint)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *Uint) UnmarshalJSON(data []byte) error {
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
