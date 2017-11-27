package generic

import "testing"

func TestAsFloatInt(t *testing.T) {
	var i int
	i = int(100)
	asFloatTest(i, t)
}

func TestAsFloatInt8(t *testing.T) {
	var i int8
	i = int8(100)
	asFloatTest(i, t)
}

func TestAsFloatInt16(t *testing.T) {
	var i int16
	i = int16(100)
	asFloatTest(i, t)
}

func TestAsFloatInt32(t *testing.T) {
	var i int32
	i = int32(100)
	asFloatTest(i, t)
}

func TestAsFloatInt64(t *testing.T) {
	var i int64
	i = int64(100)
	asFloatTest(i, t)
}

func TestAsFloatUint(t *testing.T) {
	var u uint
	u = uint(100)
	asFloatTest(u, t)
}

func TestAsFloatUint8(t *testing.T) {
	var u uint8
	u = uint8(100)
	asFloatTest(u, t)
}

func TestAsFloatUint16(t *testing.T) {
	var u uint16
	u = uint16(100)
	asFloatTest(u, t)
}

func TestAsFloatUint32(t *testing.T) {
	var u uint32
	u = uint32(100)
	asFloatTest(u, t)
}

func TestAsFloatUint64(t *testing.T) {
	var u uint64
	u = uint64(100)
	asFloatTest(u, t)
}

func TestAsFloatFloat32(t *testing.T) {
	var f float32
	f = float32(100.0001)
	r, v, err := asFloat(f)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if v == false {
		t.Error("expected: true, actual: false")
	}
	if float32(r) != 100.0001 {
		t.Errorf("expected: 100.0001, actual: %v", r)
	}
}

func TestAsFloatFloat64(t *testing.T) {
	var f float64
	f = float64(100.0001)
	r, v, err := asFloat(f)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if v == false {
		t.Error("expected: true, actual: false")
	}
	if r != 100.0001 {
		t.Errorf("expected: 100, actual: %v", r)
	}
}

func TestAsFloatTrue(t *testing.T) {
	b := true
	r, v, err := asFloat(b)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if v == false {
		t.Error("expected: true, actual: false")
	}
	if r != 1 {
		t.Errorf("expected: 1, actual: %v", r)
	}
}

func TestAsFloatFalse(t *testing.T) {
	b := false
	r, v, err := asFloat(b)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if v == false {
		t.Error("expected: true, actual: false")
	}
	if r != 0 {
		t.Errorf("expected: 0, actual: %v", r)
	}
}

func TestAsFloatNumericString(t *testing.T) {
	s := "100"
	asFloatTest(s, t)
}

func TestAsFloatUnumericString(t *testing.T) {
	s := "abd"
	_, _, err := asFloat(s)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestAsFloatInvalidType(t *testing.T) {
	bs := []byte("1")
	_, _, err := asFloat(bs)
	if err == nil {
		t.Error("Expected error")
	}
}

func asFloatTest(x interface{}, t *testing.T) {
	r, v, err := asFloat(x)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if v == false {
		t.Error("expected: true, actual: false")
	}
	if r != 100 {
		t.Errorf("expected: 100, actual: %v", r)
	}
}
