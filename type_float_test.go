package generic

import (
	"encoding/json"
	"testing"
)

type TestFloatStruct struct {
	Int       Float `json:"int"`
	Float     Float `json:"float"`
	Bool      Float `json:"bool"`
	String    Float `json:"string"`
	NullValue Float `json:"null_value"`
}

func TestFloatJsonUnmarshalAndMarshal(t *testing.T) {
	var ts TestFloatStruct
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

func TestFloatJsonError(t *testing.T) {
	var ts TestFloatStruct
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

func TestFloatSetNil(t *testing.T) {
	ti := Float{}
	err := ti.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ti.Weak() != nil {
		t.Errorf("This value should return nil. error:%#v", ti.Weak())
	}
}

func TestFloatSetInt64(t *testing.T) {
	var v int64 = 100
	var expected float64 = 100
	ti := Float{}
	err := ti.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ti.Weak() != expected {
		t.Errorf("actual:%#v, expected:%#v", ti.Weak(), expected)
	}
}

func TestFloatSetNumericString(t *testing.T) {
	v := "56.0001"
	expected := 56.0001
	ti := Float{}
	err := ti.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ti.Weak() != expected {
		t.Errorf("actual:%v, expected:%v", ti.Weak(), expected)
	}
}

func TestFloatSetNonNumericString(t *testing.T) {
	v := "a"
	var expected float64
	ti := Float{}
	err := ti.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if ti.Weak() == expected {
		t.Errorf("This value should return 0. value:%#v", ti.Weak())
	}
}
