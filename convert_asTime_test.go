package generic

import (
	"testing"
	"time"
)

func TestAsTimeTime(t *testing.T) {
	x := time.Date(2020, time.Month(7), 24, 20, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
	r, v, err := asTime(x)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if v == false {
		t.Error("expected: true, actual: false")
	}
	if s := r.String(); s != "2020-07-24 20:00:00 +0900 Asia/Tokyo" {
		t.Errorf("expected: 2020-07-24 20:00:00 +0900 Asia/Tokyo, actual: %s", s)
	}
}

func TestAsTimeZero(t *testing.T) {
	x := time.Time{}
	r, v, err := asTime(x)
	if err != nil {
		t.Errorf("Not Expected error. error:%s", err.Error())
	}
	if v == false {
		t.Error("expected: true, actual: false")
	}
	if !r.IsZero() {
		t.Errorf("expected: time.IsZero is true, actual: %s", r.String())
	}
}
