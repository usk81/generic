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
		int:       1000,
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
		int:       1000,
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
		int:       0,
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
			int:       10,
		},
		Float: Int{
			ValidFlag: true,
			int:       1,
		},
		Bool: Int{
			ValidFlag: true,
			int:       1,
		},
		String: Int{
			ValidFlag: true,
			int:       -50,
		},
		NullValue: Int{
			ValidFlag: false,
			int:       0,
		},
		Empty: Int{
			ValidFlag: false,
			int:       0,
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

func TestIntInt(t *testing.T) {
	expected := 123456789
	ti := Int{
		ValidFlag: true,
		int:       int64(expected),
	}
	if ti.Int() != expected {
		t.Errorf("actual:%d, expected:%d", ti.Int(), expected)
	}
}

func TestIntIntInvalid(t *testing.T) {
	ti := Int{
		ValidFlag: false,
		int:       123456789,
	}
	if ti.Int() != 0 {
		t.Errorf("actual:%d, expected:0", ti.Int())
	}
}

func TestIntInt32(t *testing.T) {
	var expected int32 = 123456789
	ti := Int{
		ValidFlag: true,
		int:       int64(expected),
	}
	if ti.Int32() != expected {
		t.Errorf("actual:%d, expected:%d", ti.Int32(), expected)
	}
}

func TestIntInt32Invalid(t *testing.T) {
	ti := Int{
		ValidFlag: false,
		int:       123456789,
	}
	if ti.Int32() != 0 {
		t.Errorf("actual:%d, expected:0", ti.Int32())
	}
}

func TestIntInt64(t *testing.T) {
	var expected int64 = 123456789
	ti := Int{
		ValidFlag: true,
		int:       expected,
	}
	if ti.Int64() != expected {
		t.Errorf("actual:%d, expected:%d", ti.Int64(), expected)
	}
}

func TestIntInt64Invalid(t *testing.T) {
	ti := Int{
		ValidFlag: false,
		int:       123456789,
	}
	if ti.Int64() != 0 {
		t.Errorf("actual:%d, expected:0", ti.Int64())
	}
}

func TestIntString(t *testing.T) {
	var expected = "123456789"
	ti := Int{}
	ti.Set(expected)
	if ti.String() != expected {
		t.Errorf("actual:%s, expected:%s", ti.String(), expected)
	}
}

func TestIntStringInvalid(t *testing.T) {
	ti := Int{
		ValidFlag: false,
		int:       123456789,
	}
	if ti.String() != "" {
		t.Errorf("expected empty string, actual:%s", ti.String())
	}
}
