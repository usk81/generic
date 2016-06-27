package generic

import (
	"encoding/json"
	"testing"
)

type TestUintStruct struct {
	Int       TypeUint `json:"int"`
	Float     TypeUint `json:"float"`
	Bool      TypeUint `json:"bool"`
	String    TypeUint `json:"string"`
	NullValue TypeUint `json:"null_value"`
}

func TestTypeUintJsonUnmarshalAndMarshal(t *testing.T) {
	var ts TestUintStruct
	jstr := `{"int":10,"float":1.0,"bool":true,"string":"50","null_value":null}`
	expected := `{"int":10,"float":1,"bool":1,"string":50,"null_value":null}`
	err := json.Unmarshal([]byte(jstr), &ts)
	if err != nil {
		t.Errorf("Not Expected error when json.Unmarshal. error:%v", err.Error())
	}
	b, err := json.Marshal(ts)
	if err != nil {
		t.Errorf("Not Expected error when json.Marshal. error:%v", err.Error())
	}
	actual := string(b)
	if actual != expected {
		t.Errorf("actual:%s, expected:%s", actual, expected)
	}
}

func TestTypeUintJsonError(t *testing.T) {
	var ts TestUintStruct
	jstr := `{"int":-10,"float":1.0,"bool":true,"string":"50","null_value":null}`
	expected := `{"int":null,"float":null,"bool":null,"string":null,"null_value":null}`
	err := json.Unmarshal([]byte(jstr), &ts)
	if err == nil {
		t.Error("Expected error when json.Unmarshal.")
	}
	b, err := json.Marshal(ts)
	if err != nil {
		t.Errorf("Not Expected error when json.Marshal. error:%v", err.Error())
	}
	actual := string(b)
	if actual != expected {
		t.Errorf("actual:%s, expected:%s", actual, expected)
	}
}

func TestTypeUintSetNil(t *testing.T) {
	tu := TypeUint{}
	err := tu.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if tu.Value() != nil {
		t.Errorf("This value should return nil. error:%#v", tu.Value())
	}
}

func TestTypeUintSetInt64(t *testing.T) {
	var v int64 = 100
	var expected uint64 = 100
	tu := TypeUint{}
	err := tu.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if tu.Value() != expected {
		t.Errorf("This value should return 100 (uint64). value:%#v", tu.Value())
	}
}

func TestTypeUintSetNumericString(t *testing.T) {
	var v string = "56"
	var expected uint64 = 56
	tu := TypeUint{}
	err := tu.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if tu.Value() != expected {
		t.Errorf("This value should return nil. error:%#v", tu.Value())
	}
}

func TestTypeUintSetNonNumericString(t *testing.T) {
	var v string = "a"
	var expected uint64 = 0
	tu := TypeUint{}
	err := tu.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if tu.Value() == expected {
		t.Errorf("This value should return 0. value:%#v", tu.Value())
	}
}
