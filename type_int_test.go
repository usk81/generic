package generic

import (
	"encoding/json"
	"testing"
)

type TestIntStruct struct {
	Int       TypeInt `json:"int"`
	Float     TypeInt `json:"float"`
	Bool      TypeInt `json:"bool"`
	String    TypeInt `json:"string"`
	NullValue TypeInt `json:"null_value"`
}

func TestTypeIntJsonUnmarshalAndMarshal(t *testing.T) {
	var ts TestIntStruct
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

func TestTypeIntJsonError(t *testing.T) {
	var ts TestIntStruct
	jstr := `{"int":10,"float":1.0,"bool":true,"string":"あ","null_value":null}`
	expected := `{"int":10,"float":1,"bool":1,"string":null,"null_value":null}`
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

func TestTypeIntSetNil(t *testing.T) {
	ti := TypeUint{}
	err := ti.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ti.Value() != nil {
		t.Errorf("This value should return nil. error:%#v", ti.Value())
	}
}

func TestTypeIntSetInt64(t *testing.T) {
	var v int64 = 100
	var expected int64 = 100
	ti := TypeInt{}
	err := ti.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ti.Value() != expected {
		t.Errorf("actual:%v, expected:%v", ti.Value(), expected)
	}
}

func TestTypeIntSetNumericString(t *testing.T) {
	v := "56"
	var expected int64 = 56
	ti := TypeInt{}
	err := ti.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ti.Value() != expected {
		t.Errorf("actual:%v, expected:%v", ti.Value(), expected)
	}
}

func TestTypeIntSetNonNumericString(t *testing.T) {
	v := "a"
	var expected int64
	ti := TypeInt{}
	err := ti.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if ti.Value() == expected {
		t.Errorf("This value should return 0. value:%#v", ti.Value())
	}
}
