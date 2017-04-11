package generic

import "testing"

func TestTimeSetNil(t *testing.T) {
	tt := Time{}
	err := tt.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%v", err.Error())
	}
	if tt.Value() != nil {
		t.Errorf("This value should return nil. error:%#v", tt.Value())
	}
}

func TestTimeSetInt64(t *testing.T) {
	var v int64 = 1367059792
	tt := Time{}
	err := tt.Set(v)
	if err == nil {
		t.Errorf("Not Expected error")
	}
	if tt.Value() != nil {
		t.Errorf("This value should return nil. error:%#v", tt.Value())
	}
}

func TestTimeSetNumericString(t *testing.T) {
	v := "1467059792"
	tt := Time{}
	err := tt.Set(v)
	if err == nil {
		t.Errorf("Expected error.")
	}
	if tt.Value() != nil {
		t.Errorf("This value should return nil. error:%#v", tt.Value())
	}
}

func TestTimeSetNonNumericString(t *testing.T) {
	v := "a"
	tt := Time{}
	err := tt.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if tt.Value() != nil {
		t.Errorf("This value should return nil. error:%#v", tt.Value())
	}
}

func TestTimeSetBool(t *testing.T) {
	v := true
	tt := Time{}
	err := tt.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if tt.Value() != nil {
		t.Errorf("This value should return nil. error:%#v", tt.Value())
	}
}
