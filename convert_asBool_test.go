package generic

import "testing"

func TestAsBoolInt(t *testing.T) {
	var i int
	i = 100
	asBoolTest(i, t)
}

func TestAsBoolInt8(t *testing.T) {
	var i int8
	i = 100
	asBoolTest(i, t)
}

func TestAsBoolInt16(t *testing.T) {
	var i int16
	i = 100
	asBoolTest(i, t)
}

func TestAsBoolInt32(t *testing.T) {
	var i int32
	i = 100
	asBoolTest(i, t)
}

func TestAsBoolInt64(t *testing.T) {
	var i int64
	i = 100
	asBoolTest(i, t)
}

func TestAsBoolUint(t *testing.T) {
	var u uint
	u = 100
	asBoolTest(u, t)
}

func TestAsBoolUint8(t *testing.T) {
	var u uint8
	u = 100
	asBoolTest(u, t)
}

func TestAsBoolUint16(t *testing.T) {
	var u uint16
	u = 100
	asBoolTest(u, t)
}

func TestAsBoolUint32(t *testing.T) {
	var u uint32
	u = 100
	asBoolTest(u, t)
}

func TestAsBoolUint64(t *testing.T) {
	var u uint64
	u = 100
	asBoolTest(u, t)
}

func TestAsBoolFloat32(t *testing.T) {
	var f float32
	f = 1.0001
	asBoolTest(f, t)
}

func TestAsBoolFloat64(t *testing.T) {
	var f float64
	f = 1.0001
	asBoolTest(f, t)
}

func TestAsBoolTrue(t *testing.T) {
	b := true
	asBoolTest(b, t)
}

func TestAsBoolFalse(t *testing.T) {
	b := false
	r, v, err := asBool(b)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if v == false {
		t.Error("expected: true, actual: false")
	}
	if r == true {
		t.Error("expected: false, actual: true")
	}
}

func TestAsBoolNumericString(t *testing.T) {
	s := "1"
	asBoolTest(s, t)
}

func TestAsBoolUnumericString(t *testing.T) {
	s := "abd"
	_, _, err := asBool(s)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestAsBoolInvalidType(t *testing.T) {
	bs := []byte("true")
	_, _, err := asBool(bs)
	if err == nil {
		t.Error("Expected error")
	}
}

func asBoolTest(x interface{}, t *testing.T) {
	r, v, err := asBool(x)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if v == false {
		t.Error("expected: true, actual: false")
	}
	if r == false {
		t.Error("expected: true, actual: false")
	}
}
