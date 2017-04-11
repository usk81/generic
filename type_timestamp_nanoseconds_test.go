package generic

import (
	"testing"
	"time"
)

func TestTimestampNanoSetNil(t *testing.T) {
	tn := TimestampNano{}
	err := tn.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if tn.Value() != nil {
		t.Errorf("This value should return nil. error:%#v", tn.Value())
	}
}

func TestTimestampNanoSetTime(t *testing.T) {
	v := time.Now()
	expected := v
	tn := TimestampNano{}
	err := tn.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if tn.Value().(time.Time).Unix() != expected.Unix() {
		t.Errorf("actual:%v, expected:%v", tn.Value(), expected)
	}
}

func TestTimestampNanoSetInt64(t *testing.T) {
	var v int64 = 1367059792
	expected := time.Unix(0, v)
	tn := TimestampNano{}
	err := tn.Set(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if tn.Value() != expected {
		t.Errorf("actual:%v, expected:%v", tn.Value(), expected)
	}
}

func TestTimestampNanoSetNumericString(t *testing.T) {
	v := "1467059792"
	tn := TimestampNano{}
	err := tn.Set(v)
	if err == nil {
		t.Errorf("Expected error.")
	}
	if tn.Value() != nil {
		t.Errorf("This value should return nil. value:%#v", tn.Value())
	}
}

func TestTimestampNanoSetNonNumericString(t *testing.T) {
	v := "a"
	tn := TimestampNano{}
	err := tn.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if tn.Value() != nil {
		t.Errorf("This value should return nil. value:%#v", tn.Value())
	}
}

func TestTimestampNanoSetBool(t *testing.T) {
	v := true
	tn := TimestampNano{}
	err := tn.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if tn.Value() != nil {
		t.Errorf("This value should return nil. value:%#v", tn.Value())
	}
}
