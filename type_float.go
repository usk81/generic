package generic

import "encoding/json"

// TypeFloat
type TypeFloat struct {
	ValidFlag
	Float float64
}

func (v TypeFloat) Value() interface{} {
	if !v.Valid() {
		return nil
	}
	return v.Float
}

// Scan implements the sql.Scanner interface.
func (v *TypeFloat) Scan(x interface{}) (err error) {
	v.Float, v.ValidFlag, err = asFloat(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

func (v *TypeFloat) Set(x interface{}) (err error) {
	return v.Scan(x)
}

// MarshalJSON implements the json.Marshaler interface.
func (v TypeFloat) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Float)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *TypeFloat) UnmarshalJSON(data []byte) error {
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
