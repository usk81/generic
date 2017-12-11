package generic

import "testing"

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

func unmarshalJSONBoolBenchmark(b *testing.B, bs []byte) {
	x := Bool{}
	for i := 0; i < b.N; i++ {
		x.UnmarshalJSON(bs)
	}
}

func unmarshalJSONFloatBenchmark(b *testing.B, bs []byte) {
	x := Float{}
	for i := 0; i < b.N; i++ {
		x.UnmarshalJSON(bs)
	}
}

func unmarshalJSONIntBenchmark(b *testing.B, bs []byte) {
	x := Int{}
	for i := 0; i < b.N; i++ {
		x.UnmarshalJSON(bs)
	}
}

func unmarshalJSONStringBenchmark(b *testing.B, bs []byte) {
	x := String{}
	for i := 0; i < b.N; i++ {
		x.UnmarshalJSON(bs)
	}
}

func unmarshalJSONUintBenchmark(b *testing.B, bs []byte) {
	x := Uint{}
	for i := 0; i < b.N; i++ {
		x.UnmarshalJSON(bs)
	}
}
