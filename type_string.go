package generic

import "encoding/json"

type TypeString struct {
	ValidFlag
	String string
}

func (v TypeString) Value() interface{} {
	if !v.Valid() {
		return nil
	}
	return v.String
}

func (v *TypeString) Scan(x interface{}) (err error) {
	v.String, v.ValidFlag, err = asString(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

func (v TypeString) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return json.Marshal(nil)
	}
	return json.Marshal(v.String)
}

func (v *TypeString) UnmarshalJSON(data []byte) error {
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
