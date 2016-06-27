package generic

import "encoding/json"

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

func (v *TypeFloat) Scan(x interface{}) (err error) {
	v.Float, v.ValidFlag, err = asFloat(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

func (v TypeFloat) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Float)
}

func (v *TypeFloat) UnmarshalJSON(data []byte) error {
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
