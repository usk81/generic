package generic

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestBoolStruct struct {
	Int       Bool `json:"int"`
	Float     Bool `json:"float"`
	Bool      Bool `json:"bool"`
	String    Bool `json:"string"`
	NullValue Bool `json:"null_value"`
}

func TestMarshalBool(t *testing.T) {
	expected := Bool{
		ValidFlag: true,
		bool:      true,
	}

	v := true
	actual, err := MarshalBool(v)
	if err != nil {
		t.Errorf("Not Expected error when MarshalBool. error:%s", err.Error())
	}
	if actual != expected {
		t.Errorf("actual:%v, expected:%v", actual, expected)
	}
}

func TestMustBool(t *testing.T) {
	tests := []struct {
		name      string
		args      interface{}
		want      Bool
		wantPanic bool
	}{
		{
			name: "valid",
			args: true,
			want: Bool{
				ValidFlag: true,
				bool:      true,
			},
			wantPanic: false,
		},
		{
			name: "panic",
			args: "valid paramenter",
			want: Bool{
				ValidFlag: false,
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				p := assert.Panics(t, func() {
					MustBool(tt.args)
				})
				if !p {
					t.Errorf("MustBool() panic = %v, want panic %v", p, tt.wantPanic)
				}
				return
			}
			if got := MustBool(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolJsonUnmarshalAndMarshal(t *testing.T) {
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

func TestBoolJsonError(t *testing.T) {
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

func TestBoolUnmarshalNil(t *testing.T) {
	var actual Bool
	expected := Bool{}
	err := actual.UnmarshalJSON(nil)
	if err != nil {
		t.Errorf("Not Expected error when json.Unmarshal. error:%s", err.Error())
	}
	if actual != expected {
		t.Errorf("actual:%#v, expected:%#v", actual, expected)
	}
}

func TestBoolJsonUnmarshalInvalid(t *testing.T) {
	tb := Bool{}
	if err := tb.UnmarshalJSON([]byte(`"true`)); err == nil {
		t.Errorf("Expected error when json.Unmarshal, but not; %#v", tb)
	}
}

func TestBoolSetNil(t *testing.T) {
	ts := Bool{}
	err := ts.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ts.Weak() != nil {
		t.Errorf("This value should return nil. error:%#v", ts.Weak())
	}
}

func TestBoolSetInt64(t *testing.T) {
	var v int64 = 100
	expected := true
	ts := Bool{}
	err := ts.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ts.Weak() != expected {
		t.Errorf("actual:%v, expected:%v", ts.Weak(), expected)
	}
}

func TestBoolSetString(t *testing.T) {
	v := "false"
	expected := false
	ts := Bool{}
	err := ts.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if ts.Weak() != expected {
		t.Errorf("actual:%v, expected:%v", ts.Weak(), expected)
	}
}

func TestBoolTrue(t *testing.T) {
	ts := Bool{
		ValidFlag: true,
		bool:      true,
	}
	if !ts.Bool() {
		t.Errorf("actual:%v, expected:true", ts.Bool())
	}
}

func TestBoolFalse(t *testing.T) {
	ts := Bool{
		ValidFlag: true,
		bool:      false,
	}
	if ts.Bool() {
		t.Errorf("actual:%v, expected:false", ts.Bool())
	}
}

func TestBoolInvalid(t *testing.T) {
	ts := Bool{
		ValidFlag: false,
		bool:      true,
	}
	if ts.Bool() {
		t.Errorf("actual:%v, expected:false", ts.Bool())
	}
}

func TestBoolStringTrue(t *testing.T) {
	ts := Bool{
		ValidFlag: true,
		bool:      true,
	}
	if ts.String() != "true" {
		t.Errorf("actual:%s, expected:true", ts.String())
	}
}

func TestBoolStringFalse(t *testing.T) {
	ts := Bool{
		ValidFlag: true,
		bool:      false,
	}
	if ts.String() != "false" {
		t.Errorf("actual:%s, expected:false", ts.String())
	}
}

func TestBoolStringInvalid(t *testing.T) {
	ts := Bool{
		ValidFlag: false,
		bool:      true,
	}
	if ts.String() != "false" {
		t.Errorf("actual:%s, expected:false", ts.String())
	}
}
