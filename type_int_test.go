package generic

import (
	"encoding/json"
	"testing"
)

type TestIntStruct struct {
	Int       Int `json:"int"`
	Float     Int `json:"float"`
	Bool      Int `json:"bool"`
	String    Int `json:"string"`
	NullValue Int `json:"null_value"`
	Empty     Int `json:"empty"`
}

func TestIntJsonUnmarshalAndMarshal(t *testing.T) {
	var ts TestIntStruct
	jstr := `{"int":10,"float":1.0,"bool":true,"string":"-50","null_value":null}`
	expected := `{"int":10,"float":1,"bool":1,"string":-50,"null_value":null,"empty":null}`
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

func TestIntJsonMarshal(t *testing.T) {
	ts := Int{
		ValidFlag: true,
		Int:       1000,
	}
	expected := `1000`
	actual, err := json.Marshal(ts)
	if err != nil {
		t.Errorf("Not Expected error when json.Marshal. error:%v", err.Error())
	}
	if string(actual) != expected {
		t.Errorf("actual:%s, expected:%s", string(actual), expected)
	}
}

func TestIntJsonMarshalValidFalse(t *testing.T) {
	ts := Int{
		ValidFlag: false,
		Int:       1000,
	}
	expected := []byte("null")
	actual, err := json.Marshal(ts)
	if err != nil {
		t.Errorf("Not Expected error when json.Marshal. error:%v", err.Error())
	}
	if string(actual) != string(expected) {
		t.Errorf("actual:%v, expected:%v", actual, expected)
	}
}

func TestIntJsonMarshalZero(t *testing.T) {
	ts := Int{
		ValidFlag: true,
		Int:       0,
	}
	expected := `0`
	actual, err := json.Marshal(ts)
	if err != nil {
		t.Errorf("Not Expected error when json.Marshal. error:%v", err.Error())
	}
	if string(actual) != expected {
		t.Errorf("actual:%s, expected:%s", string(actual), expected)
	}
}

func TestIntJsonUnmarshal(t *testing.T) {
	jstr := `{"int":10,"float":1.0,"bool":true,"string":"-50","null_value":null}`
	expected := TestIntStruct{
		Int: Int{
			ValidFlag: true,
			Int:       10,
		},
		Float: Int{
			ValidFlag: true,
			Int:       1,
		},
		Bool: Int{
			ValidFlag: true,
			Int:       1,
		},
		String: Int{
			ValidFlag: true,
			Int:       -50,
		},
		NullValue: Int{
			ValidFlag: false,
			Int:       0,
		},
		Empty: Int{
			ValidFlag: false,
			Int:       0,
		},
	}
	var actual TestIntStruct
	err := json.Unmarshal([]byte(jstr), &actual)
	if err != nil {
		t.Errorf("Not Expected error when json.Marshal. error:%s", err.Error())
	}
	if actual != expected {
		t.Errorf("actual:%#v, expected:%#v", actual, expected)
	}
}

func TestUnmarshalNil(t *testing.T) {
	var actual Int
	expected := Int{}
	err := actual.UnmarshalJSON(nil)
	if err != nil {
		t.Errorf("Not Expected error when json.Unmarshal. error:%s", err.Error())
	}
	if actual != expected {
		t.Errorf("actual:%#v, expected:%#v", actual, expected)
	}
}

func TestUnmarshalInvalidData(t *testing.T) {
	var actual Int
	err := actual.UnmarshalJSON([]byte(`"1`))
	if err == nil {
		t.Error("Expected error when json.Unmarshal")
	}
}

func TestIntJsonError(t *testing.T) {
	var ts TestIntStruct
	jstr := `{"int":10,"float":1.0,"bool":true,"string":"あ","null_value":null}`
	expected := `{"int":10,"float":1,"bool":1,"string":null,"null_value":null,"empty":null}`
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

func TestIntSetNil(t *testing.T) {
	ti := Int{}
	err := ti.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ti.Weak() != nil {
		t.Errorf("This value should return nil. error:%#v", ti.Weak())
	}
}

func TestIntSetInt64(t *testing.T) {
	var v int64 = 100
	var expected int64 = 100
	ti := Int{}
	err := ti.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ti.Weak() != expected {
		t.Errorf("actual:%v, expected:%v", ti.Weak(), expected)
	}
}

func TestIntSetNumericString(t *testing.T) {
	v := "56"
	var expected int64 = 56
	ti := Int{}
	err := ti.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ti.Weak() != expected {
		t.Errorf("actual:%v, expected:%v", ti.Weak(), expected)
	}
}

func TestIntSetNonNumericString(t *testing.T) {
	v := "a"
	var expected int64
	ti := Int{}
	err := ti.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if ti.Weak() == expected {
		t.Errorf("This value should return 0. value:%#v", ti.Weak())
	}
}
