package generic

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTimeJsonMarshal(t *testing.T) {
	v := time.Now()
	tt := Time{
		ValidFlag: true,
		time:      v,
	}
	expected := `"` + v.Format(time.RFC3339Nano) + `"`
	actual, err := json.Marshal(tt)
	if err != nil {
		t.Errorf("Not Expected error when json.Marshal. error:%v", err.Error())
	}
	if string(actual) != expected {
		t.Errorf("actual:%s, expected:%s", string(actual), expected)
	}
}

func TestTimeJsonMarshalValidFalse(t *testing.T) {
	tt := Time{
		ValidFlag: false,
		time:      time.Now(),
	}
	expected := []byte("null")
	actual, err := json.Marshal(tt)
	if err != nil {
		t.Errorf("Not Expected error when json.Marshal. error:%v", err.Error())
	}
	if string(actual) != string(expected) {
		t.Errorf("actual:%v, expected:%v", actual, expected)
	}
}

func TestTimeSetNil(t *testing.T) {
	tt := Time{}
	err := tt.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if tt.Weak() != nil {
		t.Errorf("This value should return nil. error:%#v", tt.Weak())
	}
}

func TestTimeSetTime(t *testing.T) {
	v := time.Now()
	tt := Time{}
	err := tt.Set(v)
	if err != nil {
		t.Errorf("Not Expected error")
	}
	if tt.Weak() == nil {
		t.Errorf("This value should return nil. error:%#v", tt.Weak())
	}
}

func TestTimeSetInt64(t *testing.T) {
	var v int64 = 1367059792
	tt := Time{}
	err := tt.Set(v)
	if err == nil {
		t.Errorf("Not Expected error")
	}
	if tt.Weak() != nil {
		t.Errorf("This value should return nil. error:%#v", tt.Weak())
	}
}

func TestTimeSetNumericString(t *testing.T) {
	v := "1467059792"
	tt := Time{}
	err := tt.Set(v)
	if err == nil {
		t.Errorf("Expected error.")
	}
	if tt.Weak() != nil {
		t.Errorf("This value should return nil. error:%#v", tt.Weak())
	}
}

func TestTimeSetNonNumericString(t *testing.T) {
	v := "a"
	tt := Time{}
	err := tt.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if tt.Weak() != nil {
		t.Errorf("This value should return nil. error:%#v", tt.Weak())
	}
}

func TestTimeSetBool(t *testing.T) {
	v := true
	tt := Time{}
	err := tt.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if tt.Weak() != nil {
		t.Errorf("This value should return nil. error:%#v", tt.Weak())
	}
}

func TestTimeString(t *testing.T) {
	var expected = time.Now()
	tt := Time{}
	tt.Set(expected)
	if tt.String() != expected.String() {
		t.Errorf("actual:%s, expected:%s", tt.String(), expected.String())
	}
}

func TestTimeStringInvalid(t *testing.T) {
	tt := Time{
		ValidFlag: false,
		time:      time.Now(),
	}
	if tt.String() != "" {
		t.Errorf("expected empty string, actual:%s", tt.String())
	}
}
