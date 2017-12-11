package generic

import (
	"testing"
	"time"
)

func BenchmarkMarshalJSONBool(b *testing.B) {
	x := Bool{
		ValidFlag: true,
		Bool:      true,
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON()
	}
}

func BenchmarkMarshalJSONFloat(b *testing.B) {
	x := Float{
		ValidFlag: true,
		Float:     1000.000001,
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON()
	}
}

func BenchmarkMarshalJSONInt(b *testing.B) {
	x := Int{
		ValidFlag: true,
		Int:       10000,
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON()
	}
}

func BenchmarkMarshalJSONString(b *testing.B) {
	x := String{
		ValidFlag: true,
		String:    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON()
	}
}

func BenchmarkMarshalJSONStringLarge(b *testing.B) {
	x := String{
		ValidFlag: true,
		String:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON()
	}
}

func BenchmarkMarshalJSONTime(b *testing.B) {
	x := Time{
		ValidFlag: true,
		Time:      time.Now(),
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON()
	}
}

func BenchmarkMarshalJSONTimestampMS(b *testing.B) {
	x := TimestampMS{
		ValidFlag: true,
		Time:      time.Now(),
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON()
	}
}

func BenchmarkMarshalJSONTimestampNano(b *testing.B) {
	x := TimestampNano{
		ValidFlag: true,
		Time:      time.Now(),
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON()
	}
}

func BenchmarkMarshalJSONTimestamp(b *testing.B) {
	x := Timestamp{
		ValidFlag: true,
		Time:      time.Now(),
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON()
	}
}

func BenchmarkMarshalJSONUint(b *testing.B) {
	x := Uint{
		ValidFlag: true,
		Uint:      10000,
	}
	for i := 0; i < b.N; i++ {
		x.MarshalJSON()
	}
}
