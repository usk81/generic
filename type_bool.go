package generic

import "encoding/json"

// Bool
type Bool struct {
	ValidFlag
	Bool bool
}

// Value returns Bool.Bool, but if Bool.ValidFlag is false, returns nil.
func (v Bool) Value() interface{} {
	if !v.Valid() {
		return nil
	}
	return v.Bool
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

// Set sets a specified value.
func (v *Bool) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// MarshalJSON implements the json.Marshaler interface.
func (v Bool) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Bool)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *Bool) UnmarshalJSON(data []byte) error {
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
