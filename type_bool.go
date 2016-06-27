package generic

import "encoding/json"

type TypeBool struct {
	ValidFlag
	Bool bool
}

func (v TypeBool) Value() interface{} {
	if !v.Valid() {
		return nil
	}
	return v.Bool
}

func (v *TypeBool) Scan(x interface{}) (err error) {
	v.Bool, v.ValidFlag, err = asBool(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

func (v TypeBool) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Bool)
}

func (v *TypeBool) UnmarshalJSON(data []byte) error {
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
