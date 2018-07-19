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
	HTML      String `json:"html"`
	NullValue String `json:"null_value"`
}

func TestMarshalString(t *testing.T) {
	expected := String{
		ValidFlag: true,
		string:    "foobar",
	}

	s := "foobar"
	actual, err := MarshalString(s)
	if err != nil {
		t.Errorf("Not Expected error when MarshalString. error:%v", err.Error())
	}
	if actual != expected {
		t.Errorf("actual:%v, expected:%v", actual, expected)
	}
}

func TestStringJsonUnmarshalAndMarshal(t *testing.T) {
	var ts TestStringStruct
	jstr := `{"int":10,"float":1.1,"bool":false,"string":"qwertyuiopkjhgv876","html":"https://golang.org/src/encoding/json/encode.go?h=float64Encoder&foo=bar#L409","null_value":null}`
	expected := `{"int":"10","float":"1.1","bool":"false","string":"qwertyuiopkjhgv876","html":"https://golang.org/src/encoding/json/encode.go?h=float64Encoder\u0026foo=bar#L409","null_value":null}`
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

func TestStringUnmarshalNil(t *testing.T) {
	var actual String
	expected := String{}
	err := actual.UnmarshalJSON(nil)
	if err != nil {
		t.Errorf("Not Expected error when json.Unmarshal. error:%s", err.Error())
	}
	if actual != expected {
		t.Errorf("actual:%#v, expected:%#v", actual, expected)
	}
}

func TestStringUnmarshalNull(t *testing.T) {
	var actual String
	expected := String{}
	err := actual.UnmarshalJSON([]byte("null"))
	if err != nil {
		t.Errorf("Not Expected error when json.Unmarshal. error:%s", err.Error())
	}
	if actual != expected {
		t.Errorf("actual:%#v, expected:%#v", actual, expected)
	}
}

func TestStringSetNil(t *testing.T) {
	ts := String{}
	err := ts.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ts.Weak() != nil {
		t.Errorf("This value should return nil. error:%#v", ts.Weak())
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
	if ts.Weak() != expected {
		t.Errorf("actual:%v, expected:%v", ts.Weak(), expected)
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
	if ts.Weak() != expected {
		t.Errorf("actual:%v, expected:%v", ts.Weak(), expected)
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
	if ts.Weak() != expected {
		t.Errorf("actual:%v, expected:%v", ts.Weak(), expected)
	}
}

func TestString(t *testing.T) {
	expected := "vcrtyhjki876tfdews"
	ts := String{
		ValidFlag: true,
		string:    expected,
	}
	if ts.String() != expected {
		t.Errorf("actual:%s, expected:%s", ts.String(), expected)
	}
}

func TestStringInvalid(t *testing.T) {
	ts := String{
		ValidFlag: false,
		string:    "vcrtyhjki876tfdews",
	}
	if ts.String() != "" {
		t.Errorf("actual:%s, expected: (empty)", ts.String())
	}
}
