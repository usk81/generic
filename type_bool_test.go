package generic

import (
	"encoding/json"
	"testing"
)

type TestBoolStruct struct {
	Int       TypeBool `json:"int"`
	Float     TypeBool `json:"float"`
	Bool      TypeBool `json:"bool"`
	String    TypeBool `json:"string"`
	NullValue TypeBool `json:"null_value"`
}

func TestTypeBoolJsonUnmarshalAndMarshal(t *testing.T) {
	var ts TestBoolStruct
	jstr := `{"int":10,"float":1.1,"bool":false,"string":"1","null_value":null}`
	expected := `{"int":true,"float":true,"bool":false,"string":true,"null_value":null}`
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

func TestTypeBoolJsonError(t *testing.T) {
	var ts TestBoolStruct
	jstr := `{"int":10,"float":1.0,"bool":true,"string":"あ","null_value":null}`
	expected := `{"int":true,"float":true,"bool":true,"string":null,"null_value":null}`
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

func TestTypeBoolSetNil(t *testing.T) {
	ts := TypeBool{}
	err := ts.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ts.Value() != nil {
		t.Errorf("This value should return nil. error:%#v", ts.Value())
	}
}

func TestTypeBoolSetInt64(t *testing.T) {
	var v int64 = 100
	expected := true
	ts := TypeBool{}
	err := ts.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ts.Value() != expected {
		t.Errorf("actual:%v, expected:%v", ts.Value(), expected)
	}
}

func TestTypeBoolSetString(t *testing.T) {
	v := "false"
	expected := false
	ts := TypeBool{}
	err := ts.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ts.Value() != expected {
		t.Errorf("actual:%v, expected:%v", ts.Value(), expected)
	}
}
