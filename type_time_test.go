package generic

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type TestTimeStruct struct {
	Int       Time `json:"int"`
	Float     Time `json:"float"`
	String    Time `json:"string"`
	NullValue Time `json:"null_value"`
}

func TestTimeJsonUnmarshalAndMarshal(t *testing.T) {
	var ts TestTimeStruct
	jstr := `{"int":10,"float":1.0,"string":"50","null_value":null}`
	expected := fmt.Sprintf(`{"int":"%s","float":"%s","string":"%s","null_value":null}`, time.Unix(10, 0).Format(time.RFC3339), time.Unix(1, 0).Format(time.RFC3339), time.Unix(50, 0).Format(time.RFC3339))
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

func TestTimeJsonError(t *testing.T) {
	var ts TestTimeStruct
	jstr := `{"int":0,"float":1.0,"string":"あ","null_value":null}`
	expected := fmt.Sprintf(`{"int":"%s","float":"%s","string":null,"null_value":null}`, time.Unix(0, 0).Format(time.RFC3339), time.Unix(1, 0).Format(time.RFC3339))
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

func TestTimeSetNil(t *testing.T) {
	tt := Uint{}
	err := tt.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if tt.Value() != nil {
		t.Errorf("This value should return nil. error:%#v", tt.Value())
	}
}

func TestTimeSetInt64(t *testing.T) {
	var v int64 = 1367059792
	expected := time.Unix(v, 0)
	tt := Time{}
	err := tt.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if tt.Value() != expected {
		t.Errorf("actual:%v, expected:%v", tt.Value(), expected)
	}
}

func TestTimeSetNumericString(t *testing.T) {
	v := "1467059792"
	expected := time.Unix(1467059792, 0)
	tt := Time{}
	err := tt.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if tt.Value() != expected {
		t.Errorf("actual:%v, expected:%v", tt.Value(), expected)
	}
}

func TestTimeSetNonNumericString(t *testing.T) {
	v := "a"
	expected := time.Unix(0, 0)
	tt := Time{}
	err := tt.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if tt.Value() == expected {
		t.Errorf("This value should return 0. value:%#v", tt.Value())
	}
}

func TestTimeSetBool(t *testing.T) {
	v := true
	expected := time.Unix(0, 0)
	tt := Time{}
	err := tt.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if tt.Value() == expected {
		// t.Errorf("This value should return 0. value:%#v", tt.Value())
		t.Errorf("actual:%v, expected:%v", tt.Value(), expected)
	}
}
