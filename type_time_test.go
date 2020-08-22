package generic

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMarshalTime(t *testing.T) {
	v := time.Now()
	expected := v
	ts, err := MarshalTime(v)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if ts.Weak() != expected {
		t.Errorf("actual:%v, expected:%v", ts.Weak(), expected)
	}
}

func TestMustTime(t *testing.T) {
	v := time.Now()
	tests := []struct {
		name      string
		args      interface{}
		want      Time
		wantPanic bool
	}{
		{
			name: "valid",
			args: v,
			want: Time{
				ValidFlag: true,
				time:      v,
			},
			wantPanic: false,
		},
		{
			name: "panic",
			args: "valid paramenter",
			want: Time{
				ValidFlag: false,
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				p := assert.Panics(t, func() {
					MustTime(tt.args)
				})
				if !p {
					t.Errorf("MustTime() panic = %v, want panic %v", p, tt.wantPanic)
				}
				return
			}
			if got := MustTime(tt.args); got.Time() != v {
				t.Errorf("MustTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeJsonMarshal(t *testing.T) {
	v := time.Now()
	tt := Time{
		ValidFlag: true,
		time:      v,
	}
	expected := `"` + v.Format(time.RFC3339Nano) + `"`
	actual, err := json.Marshal(tt)
	if err != nil {
		t.Errorf("Not Expected error when json.Marshal. error:%v", err.Error())
	}
	if string(actual) != expected {
		t.Errorf("actual:%s, expected:%s", string(actual), expected)
	}
}

func TestTimeJsonMarshalValidFalse(t *testing.T) {
	tt := Time{
		ValidFlag: false,
		time:      time.Now(),
	}
	expected := []byte("null")
	actual, err := json.Marshal(tt)
	if err != nil {
		t.Errorf("Not Expected error when json.Marshal. error:%v", err.Error())
	}
	if string(actual) != string(expected) {
		t.Errorf("actual:%v, expected:%v", actual, expected)
	}
}

func TestTimeJsonUnmarshal(t *testing.T) {
	v := time.Now()
	in, _ := v.MarshalJSON()
	tt := Time{}
	if err := tt.UnmarshalJSON(in); err != nil {
		t.Errorf("Not Expected error when json.Unmarshal. error:%v", err.Error())
	}
	if !tt.Valid() {
		t.Error("ValidFlag should be TRUE")
	}
	if tt.Time().Format(time.RFC3339) != v.Format(time.RFC3339) {
		t.Errorf("actual:%v, expected:%v", tt.Time(), v)
	}
}

func TestTimeJsonUnmarshalNil(t *testing.T) {
	tt := Time{}
	if err := tt.UnmarshalJSON(nil); err != nil {
		t.Errorf("Not Expected error when json.Unmarshal. error:%v", err.Error())
	}
	if tt.Valid() {
		t.Error("ValidFlag should be FALSE")
	}
	if tt.Time() != time.Unix(0, 0) {
		t.Errorf("actual:%v, expected:%v", tt.Time(), time.Unix(0, 0))
	}
}

func TestTimeJsonUnmarshalInvalid(t *testing.T) {
	tt := Time{}
	if err := tt.UnmarshalJSON([]byte(`"a`)); err == nil {
		t.Errorf("Expected error when json.Unmarshal, but not; %#v", tt)
	}
}

func TestTimeSetNil(t *testing.T) {
	tt := Time{}
	err := tt.Set(nil)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if _, err = tt.Value(); err != nil {
		t.Errorf("This value should return nil. error:%s", err.Error())
	}
}

func TestTimeSetTime(t *testing.T) {
	v := time.Now()
	tt := Time{}
	err := tt.Set(v)
	if err != nil {
		t.Errorf("Not Expected error")
	}
	if tt.Weak() == nil {
		t.Errorf("This value should return nil. error:%#v", tt.Weak())
	}
}

func TestTimeSetInt64(t *testing.T) {
	var v int64 = 1367059792
	tt := Time{}
	err := tt.Set(v)
	if err == nil {
		t.Errorf("Not Expected error")
	}
	if tt.Weak() != nil {
		t.Errorf("This value should return nil. error:%#v", tt.Weak())
	}
}

func TestTimeSetNumericString(t *testing.T) {
	v := "1467059792"
	tt := Time{}
	err := tt.Set(v)
	if err == nil {
		t.Errorf("Expected error.")
	}
	if tt.Weak() != nil {
		t.Errorf("This value should return nil. error:%#v", tt.Weak())
	}
}

func TestTimeSetNonNumericString(t *testing.T) {
	v := "a"
	tt := Time{}
	err := tt.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if tt.Weak() != nil {
		t.Errorf("This value should return nil. error:%#v", tt.Weak())
	}
}

func TestTimeSetBool(t *testing.T) {
	v := true
	tt := Time{}
	err := tt.Set(v)
	if err == nil {
		t.Error("Expected error.")
	}
	if tt.Weak() != nil {
		t.Errorf("This value should return nil. error:%#v", tt.Weak())
	}
}

func TestTimeString(t *testing.T) {
	var expected = time.Now()
	tt := Time{}
	tt.Set(expected) // nolint
	if tt.String() != expected.String() {
		t.Errorf("actual:%s, expected:%s", tt.String(), expected.String())
	}
}

func TestTimeStringInvalid(t *testing.T) {
	tt := Time{
		ValidFlag: false,
		time:      time.Now(),
	}
	if tt.String() != "" {
		t.Errorf("expected empty string, actual:%s", tt.String())
	}
}
