package generic

import (
	"encoding/json"
	"testing"
)

type TestStringStruct struct {
	Int       String `json:"int"`
	Float     String `json:"float"`
	Bool      String `json:"bool"`
	String    String `json:"string"`
	NullValue String `json:"null_value"`
}

func TestStringJsonUnmarshalAndMarshal(t *testing.T) {
	var ts TestStringStruct
	jstr := `{"int":10,"float":1.1,"bool":false,"string":"qwertyuiopkjhgv876","null_value":null}`
	expected := `{"int":"10","float":"1.1","bool":"false","string":"qwertyuiopkjhgv876","null_value":null}`
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

func TestStringSetNil(t *testing.T) {
	ts := String{}
	err := ts.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ts.Value() != nil {
		t.Errorf("This value should return nil. error:%#v", ts.Value())
	}
}

func TestStringSetInt64(t *testing.T) {
	var v int64 = 100
	expected := "100"
	ts := String{}
	err := ts.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ts.Value() != expected {
		t.Errorf("actual:%v, expected:%v", ts.Value(), expected)
	}
}

func TestStringSetString(t *testing.T) {
	v := "vcrtyhjki876tfdews"
	expected := "vcrtyhjki876tfdews"
	ts := String{}
	err := ts.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ts.Value() != expected {
		t.Errorf("actual:%v, expected:%v", ts.Value(), expected)
	}
}

func TestStringSetBool(t *testing.T) {
	v := true
	expected := "true"
	ts := String{}
	err := ts.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ts.Value() != expected {
		t.Errorf("actual:%v, expected:%v", ts.Value(), expected)
	}
}
