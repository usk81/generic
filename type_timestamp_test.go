package generic

import (
	"testing"
	"time"
)

func TestTimestampSetNil(t *testing.T) {
	ts := Timestamp{}
	err := ts.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if ts.Value() != nil {
		t.Errorf("This value should return nil. error:%#v", ts.Value())
	}
}

func TestTimestampSetTime(t *testing.T) {
	v := time.Now()
	expected := v
	ts := Timestamp{}
	err := ts.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if ts.Value() != expected {
		t.Errorf("actual:%v, expected:%v", ts.Value(), expected)
	}
}

func TestTimestampSetInt64(t *testing.T) {
	var v int64 = 1367059792
	expected := time.Unix(v, 0)
	ts := Timestamp{}
	err := ts.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if ts.Value() != expected {
		t.Errorf("actual:%v, expected:%v", ts.Value(), expected)
	}
}

func TestTimestampSetNumericString(t *testing.T) {
	v := "1467059792"
	ts := Timestamp{}
	err := ts.Set(v)
	if err == nil {
		t.Errorf("Expected error.")
	}
	if ts.Value() != nil {
		t.Errorf("This value should return nil. error:%#v", ts.Value())
	}
}

func TestTimestampSetNonNumericString(t *testing.T) {
	v := "a"
	ts := Timestamp{}
	err := ts.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if ts.Value() != nil {
		t.Errorf("This value should return nil. error:%#v", ts.Value())
	}
}

func TestTimestampSetBool(t *testing.T) {
	v := true
	ts := Timestamp{}
	err := ts.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if ts.Value() != nil {
		t.Errorf("This value should return nil. error:%#v", ts.Value())
	}
}
