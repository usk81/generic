package generic

import "testing"

func TestAsUintInt(t *testing.T) {
	i := int(100)
	asUintTest(i, t)
}

func TestAsUintIntMinus(t *testing.T) {
	x := int(-100)
	_, _, err := asUint(x)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestAsUintInt8(t *testing.T) {
	i := int8(100)
	asUintTest(i, t)
}

func TestAsUintInt8Minus(t *testing.T) {
	x := int8(-100)
	_, _, err := asUint(x)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestAsUintInt16(t *testing.T) {
	i := int16(100)
	asUintTest(i, t)
}

func TestAsUintInt16Minus(t *testing.T) {
	x := int16(-100)
	_, _, err := asUint(x)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestAsUintInt32(t *testing.T) {
	i := int32(100)
	asUintTest(i, t)
}

func TestAsUintInt32Minus(t *testing.T) {
	x := int32(-100)
	_, _, err := asUint(x)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestAsUintInt64(t *testing.T) {
	i := int64(100)
	asUintTest(i, t)
}

func TestAsUintInt64Minus(t *testing.T) {
	x := int64(-100)
	_, _, err := asUint(x)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestAsUintUint(t *testing.T) {
	u := uint(100)
	asUintTest(u, t)
}

func TestAsUintUint8(t *testing.T) {
	u := uint8(100)
	asUintTest(u, t)
}

func TestAsUintUint16(t *testing.T) {
	u := uint16(100)
	asUintTest(u, t)
}

func TestAsUintUint32(t *testing.T) {
	u := uint32(100)
	asUintTest(u, t)
}

func TestAsUintUint64(t *testing.T) {
	u := uint64(100)
	asUintTest(u, t)
}

func TestAsUintFloat32(t *testing.T) {
	f := float32(100.001)
	asUintTest(f, t)
}

func TestAsUintFloat32Minus(t *testing.T) {
	x := float32(-100.001)
	_, _, err := asUint(x)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestAsUintFloat64(t *testing.T) {
	f := float64(100.001)
	asUintTest(f, t)
}

func TestAsUintFloat64Minus(t *testing.T) {
	x := float64(-100.001)
	_, _, err := asUint(x)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestAsUintTrue(t *testing.T) {
	b := true
	r, v, err := asUint(b)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if !v {
		t.Error("expected: true, actual: false")
	}
	if r != 1 {
		t.Errorf("expected: 1, actual: %d", r)
	}
}

func TestAsUintFalse(t *testing.T) {
	b := false
	r, v, err := asUint(b)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if !v {
		t.Error("expected: true, actual: false")
	}
	if r != 0 {
		t.Errorf("expected: 0, actual: %d", r)
	}
}

func TestAsUintNumericString(t *testing.T) {
	s := "100"
	asUintTest(s, t)
}

func TestAsUintUnumericString(t *testing.T) {
	s := "abd.01"
	_, _, err := asUint(s)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestAsUintInvalidType(t *testing.T) {
	bs := []byte("100")
	_, _, err := asUint(bs)
	if err == nil {
		t.Error("Expected error")
	}
}

func asUintTest(x interface{}, t *testing.T) {
	r, v, err := asUint(x)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if !v {
		t.Error("expected: true, actual: false")
	}
	if r != 100 {
		t.Errorf("expected: 100, actual: %v", r)
	}
}
