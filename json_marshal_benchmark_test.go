package generic

import (
	"net/url"
	"testing"
	"time"
)

func BenchmarkMarshalJSONBool(b *testing.B) {
	x := Bool{
		ValidFlag: true,
		bool:      true,
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON() // nolint
	}
}

func BenchmarkMarshalJSONFloat(b *testing.B) {
	x := Float{
		ValidFlag: true,
		float:     1000.000001,
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON() // nolint
	}
}

func BenchmarkMarshalJSONInt(b *testing.B) {
	x := Int{
		ValidFlag: true,
		int:       10000,
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON() // nolint
	}
}

func BenchmarkMarshalJSONString(b *testing.B) {
	x := String{
		ValidFlag: true,
		string:    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON() // nolint
	}
}

func BenchmarkMarshalJSONStringLarge(b *testing.B) {
	x := String{
		ValidFlag: true,
		string:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON() // nolint
	}
}

func BenchmarkMarshalJSONTime(b *testing.B) {
	x := Time{
		ValidFlag: true,
		time:      time.Now(),
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON() // nolint
	}
}

func BenchmarkMarshalJSONTimestampMS(b *testing.B) {
	x := TimestampMS{
		ValidFlag: true,
		time:      time.Now(),
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON() // nolint
	}
}

func BenchmarkMarshalJSONTimestampNano(b *testing.B) {
	x := TimestampNano{
		ValidFlag: true,
		time:      time.Now(),
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON() // nolint
	}
}

func BenchmarkMarshalJSONTimestamp(b *testing.B) {
	x := Timestamp{
		ValidFlag: true,
		time:      time.Now(),
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON() // nolint
	}
}

func BenchmarkMarshalJSONUint(b *testing.B) {
	x := Uint{
		ValidFlag: true,
		uint:      10000,
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON() // nolint
	}
}

func BenchmarkMarshalJSONURL(b *testing.B) {
	x := URL{
		ValidFlag: true,
		url: &url.URL{
			Scheme: "https",
			Host:   "www.google.com",
		},
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON() // nolint
	}
}
