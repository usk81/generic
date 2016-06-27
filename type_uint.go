package generic

import "encoding/json"

type TypeUint struct {
	ValidFlag
	Uint uint64
}

func (v TypeUint) Value() interface{} {
	if !v.Valid() {
		return nil
	}
	return v.Uint
}

func (v *TypeUint) Scan(x interface{}) (err error) {
	v.Uint, v.ValidFlag, err = asUint(x)
	if err != nil {
		v.ValidFlag = false
		return err
	}
	return
}

func (v TypeUint) MarshalJSON() ([]byte, error) {
	if !v.Valid() {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Uint)
}

func (v *TypeUint) UnmarshalJSON(data []byte) error {
	var in interface{}
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Scan(in)
}
