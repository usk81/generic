package generic

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"
)

func TestTimestampNanoJsonMarshal(t *testing.T) {
	v := time.Now()
	tn := TimestampNano{
		ValidFlag: true,
		time:      v,
	}
	expected := strconv.FormatInt(v.UnixNano(), 10)
	actual, err := json.Marshal(tn)
	if err != nil {
		t.Errorf("Not Expected error when json.Marshal. error:%v", err.Error())
	}
	if string(actual) != expected {
		t.Errorf("actual:%s, expected:%s", string(actual), expected)
	}
}

func TestTimestampNanoJsonMarshalValidFalse(t *testing.T) {
	tn := TimestampNano{
		ValidFlag: false,
		time:      time.Now(),
	}
	expected := []byte("null")
	actual, err := json.Marshal(tn)
	if err != nil {
		t.Errorf("Not Expected error when json.Marshal. error:%v", err.Error())
	}
	if string(actual) != string(expected) {
		t.Errorf("actual:%v, expected:%v", actual, expected)
	}
}

func TestTimestampNanoJsonUnmarshal(t *testing.T) {
	v := time.Now()
	in, _ := v.MarshalJSON()
	tn := TimestampNano{}
	if err := tn.UnmarshalJSON(in); err != nil {
		t.Errorf("Not Expected error when json.Unmarshal. error:%v", err.Error())
	}
	if !tn.Valid() {
		t.Error("ValidFlag should be TRUE")
	}
	if tn.Int64() != v.UnixNano() {
		t.Errorf("actual:%d, expected:%d", tn.Int64(), v.UnixNano())
	}
}

func TestTimestampNanoJsonUnmarshalNil(t *testing.T) {
	tn := TimestampNano{}
	if err := tn.UnmarshalJSON(nil); err != nil {
		t.Errorf("Not Expected error when json.Unmarshal. error:%v", err.Error())
	}
	if tn.Valid() {
		t.Error("ValidFlag should be FALSE")
	}
	if tn.Int64() != 0 {
		t.Errorf("actual:%d, expected:%d", tn.Int64(), 0)
	}
}

func TestTimestampNanoSetNil(t *testing.T) {
	tn := TimestampNano{}
	err := tn.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	actual, err := tn.Value()
	if err != nil {
		t.Errorf("This value should return nil. error:%s", err.Error())
	}
	if actual != nil {
		t.Errorf("actual:%d, expected:nil", actual)
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
	if tn.Weak() != expected.UnixNano() {
		t.Errorf("actual:%d, expected:%d", tn.Weak(), expected)
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
	if actual := tn.Weak(); actual != expected.UnixNano() {
		t.Errorf("actual:%v, expected:%v", actual, expected.UnixNano())
	}
}

func TestTimestampNanoSetNumericString(t *testing.T) {
	v := "1467059792"
	tn := TimestampNano{}
	err := tn.Set(v)
	if err == nil {
		t.Errorf("Expected error.")
	}
	if tn.Weak() != nil {
		t.Errorf("This value should return nil. value:%#v", tn.Weak())
	}
}

func TestTimestampNanoSetNonNumericString(t *testing.T) {
	v := "a"
	tn := TimestampNano{}
	err := tn.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if tn.Weak() != nil {
		t.Errorf("This value should return nil. value:%#v", tn.Weak())
	}
}

func TestTimestampNanoSetBool(t *testing.T) {
	v := true
	tn := TimestampNano{}
	err := tn.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if tn.Weak() != nil {
		t.Errorf("This value should return nil. value:%#v", tn.Weak())
	}
}

func TestTimestampNanoInt64(t *testing.T) {
	v := time.Now()
	expected := v.UnixNano()
	tm := TimestampNano{}
	err := tm.Set(v)
	if err != nil {
		t.Error("Not expected error.")
	}
	if tm.Int64() != expected {
		t.Errorf("This value should return %d. value:%d", expected, tm.Int())
	}
}

func TestTimestampNanoInt64Zero(t *testing.T) {
	v := time.Unix(0, 0)
	var expected int64
	tm := TimestampNano{}
	err := tm.Set(v)
	if err != nil {
		t.Error("Not expected error.")
	}
	if tm.Int64() != expected {
		t.Errorf("This value should return %d. value:%d", expected, tm.Int())
	}
}

func TestTimestampNanoInt(t *testing.T) {
	v := time.Now()
	expected := int(v.UnixNano())
	tm := TimestampNano{}
	err := tm.Set(v)
	if err != nil {
		t.Error("Not expected error.")
	}
	if tm.Int() != expected {
		t.Errorf("This value should return %d. value:%d", expected, tm.Int())
	}
}

func TestTimestampNanoString(t *testing.T) {
	v := time.Now()
	expected := strconv.FormatInt(v.UnixNano(), 10)
	tm := TimestampNano{}
	err := tm.Set(v)
	if err != nil {
		t.Error("Not expected error.")
	}
	if tm.String() != expected {
		t.Errorf("This value should return %s. value:%s", expected, tm.String())
	}
}
