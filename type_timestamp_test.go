package generic

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"
)

func TestMarshalTimestamp(t *testing.T) {
	v := time.Now()
	expected := v
	ts, err := MarshalTimestamp(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if ts.Weak() != expected {
		t.Errorf("actual:%v, expected:%v", ts.Weak(), expected)
	}
}

func TestTimestampJsonMarshal(t *testing.T) {
	v := time.Now()
	ts := Timestamp{
		ValidFlag: true,
		time:      v,
	}
	expected := strconv.FormatInt(v.Unix(), 10)
	actual, err := json.Marshal(ts)
	if err != nil {
		t.Errorf("Not Expected error when json.Marshal. error:%v", err.Error())
	}
	if string(actual) != expected {
		t.Errorf("actual:%s, expected:%s", string(actual), expected)
	}
}

func TestTimestampJsonMarshalValidFalse(t *testing.T) {
	ts := Timestamp{
		ValidFlag: false,
		time:      time.Now(),
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

func TestTimestampJsonUnmarshal(t *testing.T) {
	v := time.Now()
	in, _ := v.MarshalJSON()
	ts := Timestamp{}
	if err := ts.UnmarshalJSON(in); err != nil {
		t.Errorf("Not Expected error when json.Unmarshal. error:%v", err.Error())
	}
	if !ts.Valid() {
		t.Error("ValidFlag should be TRUE")
	}
	if ts.Int64() != v.Unix() {
		t.Errorf("actual:%v, expected:%v", ts.Int64(), v)
	}
}

func TestTimestampJsonUnmarshalNil(t *testing.T) {
	ts := Timestamp{}
	if err := ts.UnmarshalJSON(nil); err != nil {
		t.Errorf("Not Expected error when json.Unmarshal. error:%v", err.Error())
	}
	if ts.Valid() {
		t.Error("ValidFlag should be FALSE")
	}
	if ts.Int64() != 0 {
		t.Errorf("actual:%v, expected:%v", ts.Int64(), 0)
	}
}

func TestTimestampJsonUnmarshalInvalid(t *testing.T) {
	ts := Timestamp{}
	if err := ts.UnmarshalJSON([]byte(`"a`)); err == nil {
		t.Errorf("Expected error when json.Unmarshal, but not; %#v", ts)
	}
}

func TestTimestampSetNil(t *testing.T) {
	ts := Timestamp{}
	err := ts.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	actual, err := ts.Value()
	if err != nil {
		t.Errorf("This value should return nil. error:%s", err.Error())
	}
	if actual != nil {
		t.Errorf("actual:%d, expected:nil", actual)
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
	if ts.Weak() != expected {
		t.Errorf("actual:%v, expected:%v", ts.Weak(), expected)
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
	if ts.Weak() != expected {
		t.Errorf("actual:%v, expected:%v", ts.Weak(), expected)
	}
}

func TestTimestampSetNumericString(t *testing.T) {
	v := "1467059792"
	ts := Timestamp{}
	err := ts.Set(v)
	if err == nil {
		t.Errorf("Expected error.")
	}
	if ts.Weak() != nil {
		t.Errorf("This value should return nil. actual:%#v", ts.Weak())
	}
}

func TestTimestampSetNonNumericString(t *testing.T) {
	v := "a"
	ts := Timestamp{}
	err := ts.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if ts.Weak() != nil {
		t.Errorf("This value should return nil. actual:%#v", ts.Weak())
	}
}

func TestTimestampSetBool(t *testing.T) {
	v := true
	ts := Timestamp{}
	err := ts.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if ts.Weak() != nil {
		t.Errorf("This value should return nil. actual:%#v", ts.Weak())
	}
}

func TestTimestampInt64(t *testing.T) {
	v := time.Now()
	expected := v.Unix()
	ts := Timestamp{}
	err := ts.Set(v)
	if err != nil {
		t.Error("Not expected error.")
	}
	if ts.Int64() != expected {
		t.Errorf("This value should return %d. value:%d", expected, ts.Int())
	}
}

func TestTimestampInt64Zero(t *testing.T) {
	v := time.Unix(0, 0)
	var expected int64
	ts := Timestamp{}
	err := ts.Set(v)
	if err != nil {
		t.Error("Not expected error.")
	}
	if ts.Int64() != expected {
		t.Errorf("This value should return %d. value:%d", expected, ts.Int64())
	}
}

func TestTimestampInt(t *testing.T) {
	v := time.Now()
	expected := int(v.Unix())
	ts := Timestamp{}
	err := ts.Set(v)
	if err != nil {
		t.Error("Not expected error.")
	}
	if ts.Int() != expected {
		t.Errorf("This value should return %d. value:%d", expected, ts.Int())
	}
}

func TestTimestampString(t *testing.T) {
	v := time.Now()
	expected := strconv.FormatInt(v.Unix(), 10)
	ts := Timestamp{}
	err := ts.Set(v)
	if err != nil {
		t.Error("Not expected error.")
	}
	if ts.String() != expected {
		t.Errorf("This value should return %s. value:%s", expected, ts.String())
	}
}
