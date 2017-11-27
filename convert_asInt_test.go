package generic

import "testing"

func TestAsIntInt(t *testing.T) {
	var i int
	i = int(100)
	asIntTest(i, t)
}

func TestAsIntInt8(t *testing.T) {
	var i int8
	i = int8(100)
	asIntTest(i, t)
}

func TestAsIntInt16(t *testing.T) {
	var i int16
	i = int16(100)
	asIntTest(i, t)
}

func TestAsIntInt32(t *testing.T) {
	var i int32
	i = int32(100)
	asIntTest(i, t)
}

func TestAsIntInt64(t *testing.T) {
	var i int64
	i = int64(100)
	asIntTest(i, t)
}

func TestAsIntUint(t *testing.T) {
	var u uint
	u = uint(100)
	asIntTest(u, t)
}

func TestAsIntUint8(t *testing.T) {
	var u uint8
	u = uint8(100)
	asIntTest(u, t)
}

func TestAsIntUint16(t *testing.T) {
	var u uint16
	u = uint16(100)
	asIntTest(u, t)
}

func TestAsIntUint32(t *testing.T) {
	var u uint32
	u = uint32(100)
	asIntTest(u, t)
}

func TestAsIntUint64(t *testing.T) {
	var u uint64
	u = uint64(100)
	asIntTest(u, t)
}

func TestAsIntFloat32(t *testing.T) {
	var f float32
	f = float32(100.0001)
	asIntTest(f, t)
}

func TestAsIntFloat64(t *testing.T) {
	var f float64
	f = float64(100.0001)
	asIntTest(f, t)
}

func TestAsIntTrue(t *testing.T) {
	b := true
	r, v, err := asInt(b)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if v == false {
		t.Error("expected: true, actual: false")
	}
	if r != 1 {
		t.Errorf("expected: 1, actual: %d", r)
	}
}

func TestAsIntFalse(t *testing.T) {
	b := false
	r, v, err := asInt(b)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if v == false {
		t.Error("expected: true, actual: false")
	}
	if r != 0 {
		t.Errorf("expected: 0, actual: %d", r)
	}
}

func TestAsIntNumericString(t *testing.T) {
	s := "100"
	asIntTest(s, t)
}

func TestAsIntUnumericString(t *testing.T) {
	s := "abd"
	_, _, err := asInt(s)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestAsIntInvalidType(t *testing.T) {
	bs := []byte("1")
	_, _, err := asInt(bs)
	if err == nil {
		t.Error("Expected error")
	}
}

func asIntTest(x interface{}, t *testing.T) {
	r, v, err := asInt(x)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if v == false {
		t.Error("expected: true, actual: false")
	}
	if r != 100 {
		t.Errorf("expected: 100, actual: %d", r)
	}
}
