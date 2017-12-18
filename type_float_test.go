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

func TestFloatUnmarshalEmpty(t *testing.T) {
	tf := Float{}
	err := tf.UnmarshalJSON(nil)
	if err != nil {
		t.Errorf("Not Expected error when json.Marshal. error:%s", err.Error())
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
		t.Errorf("Not Expected error when json.Marshal. error:%s", err.Error())
	}
	actual := string(b)
	if actual != expected {
		t.Errorf("actual:%s, expected:%s", actual, expected)
	}
}

func TestFloatSetNil(t *testing.T) {
	tf := Float{}
	err := tf.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if tf.Weak() != nil {
		t.Errorf("This value should return nil. error:%#v", tf.Weak())
	}
}

func TestFloatSetInt64(t *testing.T) {
	var v int64 = 100
	var expected float64 = 100
	tf := Float{}
	err := tf.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if tf.Weak() != expected {
		t.Errorf("actual:%#v, expected:%#v", tf.Weak(), expected)
	}
}

func TestFloatSetNumericString(t *testing.T) {
	v := "56.0001"
	expected := 56.0001
	tf := Float{}
	err := tf.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if tf.Weak() != expected {
		t.Errorf("actual:%v, expected:%v", tf.Weak(), expected)
	}
}

func TestFloatSetNonNumericString(t *testing.T) {
	v := "a"
	var expected float64
	tf := Float{}
	err := tf.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if tf.Weak() == expected {
		t.Errorf("This value should return 0. value:%#v", tf.Weak())
	}
}

func TestFloatFloat32(t *testing.T) {
	var expected float32 = 56.0001
	tf := Float{
		ValidFlag: true,
		float:     float64(expected),
	}
	if tf.Float32() != expected {
		t.Errorf("actual:%v, expected:%v", tf.Float32(), expected)
	}
}

func TestFloatFloat32Invalid(t *testing.T) {
	var expected float32 = 56.0001
	tf := Float{
		ValidFlag: false,
		float:     float64(expected),
	}
	if tf.Float32() != 0 {
		t.Errorf("actual:%v, expected:0", tf.Float32())
	}
}

func TestFloatFloat64(t *testing.T) {
	var expected = 56.0001
	tf := Float{
		ValidFlag: true,
		float:     expected,
	}
	if tf.Float64() != expected {
		t.Errorf("actual:%v, expected:%v", tf.Float64(), expected)
	}
}

func TestFloatFloat64Invalid(t *testing.T) {
	tf := Float{
		ValidFlag: false,
		float:     56.0001,
	}
	if tf.Float64() != 0 {
		t.Errorf("actual:%v, expected:0", tf.Float64())
	}
}

func TestFloatString(t *testing.T) {
	var expected = "56.0001"
	tf := Float{}
	tf.Set(expected)
	if tf.String() != expected {
		t.Errorf("actual:%s, expected:%s", tf.String(), expected)
	}
}

func TestFloatStringInvalid(t *testing.T) {
	tf := Float{
		ValidFlag: false,
		float:     56.0001,
	}
	if tf.String() != "" {
		t.Errorf("expected empty string, actual:%s", tf.String())
	}
}
