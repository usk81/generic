package generic

import "encoding/json"

// TypeString
type TypeString struct {
	ValidFlag
	String string
}

// Value returns TypeString.String, but if TypeString.ValidFlag is false, returns nil.
func (v TypeString) Value() interface{} {
	if !v.Valid() {
		return nil
	}
	return v.String
}

// Scan implements the sql.Scanner interface.
func (v *TypeString) Scan(x interface{}) (err error) {
	v.String, v.ValidFlag, err = asString(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

// Set sets a specified value.
func (v *TypeString) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// MarshalJSON implements the json.Marshaler interface.
func (v TypeString) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return json.Marshal(nil)
	}
	return json.Marshal(v.String)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *TypeString) UnmarshalJSON(data []byte) error {
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
