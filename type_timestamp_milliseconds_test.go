package generic

import (
	"testing"
	"time"
)

func TestTimestampMSSetNil(t *testing.T) {
	tm := TimestampMS{}
	err := tm.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if tm.Value() != nil {
		t.Errorf("This value should return nil. error:%#v", tm.Value())
	}
}

func TestTimestampMSSetTime(t *testing.T) {
	v := time.Now()
	expected := v
	tm := TimestampMS{}
	err := tm.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if tm.Value().(time.Time).Unix() != expected.Unix() {
		t.Errorf("actual:%v, expected:%v", tm.Value(), expected)
	}
}

func TestTimestampMSSetInt64(t *testing.T) {
	var v int64 = 1367059792
	expected := time.Unix(0, v*1000000)
	tm := TimestampMS{}
	err := tm.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if tm.Value() != expected {
		t.Errorf("actual:%v, expected:%v", tm.Value(), expected)
	}
}

func TestTimestampMSSetNumericString(t *testing.T) {
	v := "1467059792"
	tm := TimestampMS{}
	err := tm.Set(v)
	if err == nil {
		t.Errorf("Expected error.")
	}
	if tm.Value() != nil {
		t.Errorf("This value should return nil. value:%#v", tm.Value())
	}
}

func TestTimestampMSSetNonNumericString(t *testing.T) {
	v := "a"
	tm := TimestampMS{}
	err := tm.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if tm.Value() != nil {
		t.Errorf("This value should return nil. value:%#v", tm.Value())
	}
}

func TestTimestampMSSetBool(t *testing.T) {
	v := true
	tm := TimestampMS{}
	err := tm.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if tm.Value() != nil {
		t.Errorf("This value should return nil. value:%#v", tm.Value())
	}
}
