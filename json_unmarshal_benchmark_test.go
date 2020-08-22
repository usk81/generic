package generic

import (
	"testing"
	"time"
)

func BenchmarkUnmarshalJSONBoolFromBool(b *testing.B) {
	unmarshalJSONBoolBenchmark(b, []byte(`true`))
}

func BenchmarkUnmarshalJSONBoolFromFloat(b *testing.B) {
	unmarshalJSONBoolBenchmark(b, []byte(`1.0`))
}

func BenchmarkUnmarshalJSONBoolFromInt(b *testing.B) {
	unmarshalJSONBoolBenchmark(b, []byte(`1`))
}

func BenchmarkUnmarshalJSONBoolFromString(b *testing.B) {
	unmarshalJSONBoolBenchmark(b, []byte(`"true"`))
}

func BenchmarkUnmarshalJSONIntFromBool(b *testing.B) {
	unmarshalJSONIntBenchmark(b, []byte(`true`))
}

func BenchmarkUnmarshalJSONIntFromFloat(b *testing.B) {
	unmarshalJSONIntBenchmark(b, []byte(`1.0`))
}

func BenchmarkUnmarshalJSONIntFromInt(b *testing.B) {
	unmarshalJSONIntBenchmark(b, []byte(`1`))
}

func BenchmarkUnmarshalJSONIntFromString(b *testing.B) {
	unmarshalJSONIntBenchmark(b, []byte(`"1.0"`))
}

func BenchmarkUnmarshalJSONUintFromBool(b *testing.B) {
	unmarshalJSONUintBenchmark(b, []byte(`true`))
}

func BenchmarkUnmarshalJSONUintFromFloat(b *testing.B) {
	unmarshalJSONUintBenchmark(b, []byte(`1.0`))
}

func BenchmarkUnmarshalJSONUintFromInt(b *testing.B) {
	unmarshalJSONUintBenchmark(b, []byte(`1`))
}

func BenchmarkUnmarshalJSONUintFromString(b *testing.B) {
	unmarshalJSONUintBenchmark(b, []byte(`"1.0"`))
}

func BenchmarkUnmarshalJSONFloatFromBool(b *testing.B) {
	unmarshalJSONFloatBenchmark(b, []byte(`true`))
}

func BenchmarkUnmarshalJSONFloatFromFloat(b *testing.B) {
	unmarshalJSONFloatBenchmark(b, []byte(`1.0`))
}

func BenchmarkUnmarshalJSONFloatFromInt(b *testing.B) {
	unmarshalJSONFloatBenchmark(b, []byte(`1`))
}

func BenchmarkUnmarshalJSONFloatFromString(b *testing.B) {
	unmarshalJSONFloatBenchmark(b, []byte(`"1.0"`))
}

func BenchmarkUnmarshalJSONFloatFromUint(b *testing.B) {
	unmarshalJSONFloatBenchmark(b, []byte(`1`))
}

func BenchmarkUnmarshalJSONStringFromBool(b *testing.B) {
	unmarshalJSONStringBenchmark(b, []byte(`true`))
}

func BenchmarkUnmarshalJSONStringFromFloat(b *testing.B) {
	unmarshalJSONStringBenchmark(b, []byte(`1.0`))
}

func BenchmarkUnmarshalJSONStringFromInt(b *testing.B) {
	unmarshalJSONStringBenchmark(b, []byte(`1`))
}

func BenchmarkUnmarshalJSONStringFromString(b *testing.B) {
	unmarshalJSONStringBenchmark(b, []byte(`"true"`))
}

func BenchmarkUnmarshalJSONTimeFromString(b *testing.B) {
	now := time.Now()
	unmarshalJSONTimeBenchmark(b, []byte(`"`+now.String()+`"`))
}

func BenchmarkUnmarshalJSONURLFromString(b *testing.B) {
	unmarshalJSONURLBenchmark(b, []byte(`"https://google.com"`))
}

func unmarshalJSONBoolBenchmark(b *testing.B, bs []byte) {
	x := Bool{}
	for i := 0; i < b.N; i++ {
		x.UnmarshalJSON(bs) // nolint
	}
}

func unmarshalJSONFloatBenchmark(b *testing.B, bs []byte) {
	x := Float{}
	for i := 0; i < b.N; i++ {
		x.UnmarshalJSON(bs) // nolint
	}
}

func unmarshalJSONIntBenchmark(b *testing.B, bs []byte) {
	x := Int{}
	for i := 0; i < b.N; i++ {
		x.UnmarshalJSON(bs) // nolint
	}
}

func unmarshalJSONStringBenchmark(b *testing.B, bs []byte) {
	x := String{}
	for i := 0; i < b.N; i++ {
		x.UnmarshalJSON(bs) // nolint
	}
}

func unmarshalJSONTimeBenchmark(b *testing.B, bs []byte) {
	t := Time{}
	for i := 0; i < b.N; i++ {
		t.UnmarshalJSON(bs) // nolint
	}
}

func unmarshalJSONUintBenchmark(b *testing.B, bs []byte) {
	x := Uint{}
	for i := 0; i < b.N; i++ {
		x.UnmarshalJSON(bs) // nolint
	}
}

func unmarshalJSONURLBenchmark(b *testing.B, bs []byte) {
	x := URL{}
	for i := 0; i < b.N; i++ {
		x.UnmarshalJSON(bs) // nolint
	}
}
