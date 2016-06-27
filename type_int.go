package generic

import "encoding/json"

type TypeInt struct {
	ValidFlag
	Int int64
}

func (v TypeInt) Value() interface{} {
	if !v.Valid() {
		return nil
	}
	return v.Int
}

func (v *TypeInt) Scan(x interface{}) (err error) {
	v.Int, v.ValidFlag, err = asInt(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

func (v TypeInt) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Int)
}

func (v *TypeInt) UnmarshalJSON(data []byte) error {
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
