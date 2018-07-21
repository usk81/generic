package generic

import (
	"encoding/json"
	"testing"
)

type TestUintStruct struct {
	Int       Uint `json:"int"`
	Float     Uint `json:"float"`
	Bool      Uint `json:"bool"`
	String    Uint `json:"string"`
	NullValue Uint `json:"null_value"`
}

func TestMarshalUint(t *testing.T) {
	expected := Uint{
		ValidFlag: true,
		uint:      100,
	}

	i := 100
	actual, err := MarshalUint(i)
	if err != nil {
		t.Errorf("Not Expected error when MarshalUint. error:%v", err.Error())
	}
	if actual != expected {
		t.Errorf("actual:%v, expected:%v", actual, expected)
	}
}

func TestUintJsonUnmarshalAndMarshal(t *testing.T) {
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

func TestUintJsonUnmarshalInvalid(t *testing.T) {
	u := Uint{}
	if err := u.UnmarshalJSON([]byte(`"0`)); err == nil {
		t.Errorf("Expected error when json.Unmarshal, but not; %#v", u)
	}
}

func TestUintJsonError(t *testing.T) {
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

func TestUintSetNil(t *testing.T) {
	tu := Uint{}
	err := tu.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if tu.Weak() != nil {
		t.Errorf("This value should return nil. error:%#v", tu.Weak())
	}
}

func TestUintSetInt64(t *testing.T) {
	var v int64 = 100
	var expected uint64 = 100
	tu := Uint{}
	err := tu.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if tu.Weak() != expected {
		t.Errorf("This value should return 100 (uint64). value:%#v", tu.Weak())
	}
}

func TestUintSetNumericString(t *testing.T) {
	v := "56"
	var expected uint64 = 56
	tu := Uint{}
	err := tu.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if tu.Weak() != expected {
		t.Errorf("This value should return nil. error:%#v", tu.Weak())
	}
}

func TestUintSetNonNumericString(t *testing.T) {
	v := "a"
	var expected uint64
	tu := Uint{}
	err := tu.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if tu.Weak() == expected {
		t.Errorf("This value should return 0. value:%#v", tu.Weak())
	}
}

func TestUintUint(t *testing.T) {
	var expected uint = 123456789
	ti := Uint{
		ValidFlag: true,
		uint:      uint64(expected),
	}
	if ti.Uint() != expected {
		t.Errorf("actual:%d, expected:%d", ti.Uint(), expected)
	}
}

func TestUintUintInvalid(t *testing.T) {
	ti := Uint{
		ValidFlag: false,
		uint:      123456789,
	}
	if ti.Uint() != 0 {
		t.Errorf("actual:%d, expected:0", ti.Uint())
	}
}

func TestUintUint32(t *testing.T) {
	var expected uint32 = 123456789
	ti := Uint{
		ValidFlag: true,
		uint:      uint64(expected),
	}
	if ti.Uint32() != expected {
		t.Errorf("actual:%d, expected:%d", ti.Uint32(), expected)
	}
}

func TestUintUint32Invalid(t *testing.T) {
	ti := Uint{
		ValidFlag: false,
		uint:      123456789,
	}
	if ti.Uint32() != 0 {
		t.Errorf("actual:%d, expected:0", ti.Uint32())
	}
}

func TestUintUint64(t *testing.T) {
	var expected uint64 = 123456789
	ti := Uint{
		ValidFlag: true,
		uint:      expected,
	}
	if ti.Uint64() != expected {
		t.Errorf("actual:%d, expected:%d", ti.Uint64(), expected)
	}
}

func TestUintUint64Invalid(t *testing.T) {
	ti := Uint{
		ValidFlag: false,
		uint:      123456789,
	}
	if ti.Uint64() != 0 {
		t.Errorf("actual:%d, expected:0", ti.Uint64())
	}
}

func TestUintString(t *testing.T) {
	var expected = "123456789"
	ti := Uint{}
	ti.Set(expected)
	if ti.String() != expected {
		t.Errorf("actual:%s, expected:%s", ti.String(), expected)
	}
}

func TestUintStringInvalid(t *testing.T) {
	ti := Uint{
		ValidFlag: false,
		uint:      123456789,
	}
	if ti.String() != "" {
		t.Errorf("expected empty string, actual:%s", ti.String())
	}
}
