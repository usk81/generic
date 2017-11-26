package generic

import (
	"encoding/json"
	"strconv"
	"testing"
)

func BenchmarkConvJSON(b *testing.B) {
	var n int64 = 1000
	for i := 0; i < b.N; i++ {
		json.Marshal(n)
	}
}

func BenchmarkConvStrconv(b *testing.B) {
	var n int64 = 1000
	for i := 0; i < b.N; i++ {
		_ = []byte(strconv.FormatInt(n, 10))
	}
}

func BenchmarkUnmarshalInterface(b *testing.B) {
	jb := []byte(`"\"Fran\""`)
	var in interface{}
	for i := 0; i < b.N; i++ {
		json.Unmarshal(jb, &in)
	}
	// fmt.Println(in)
}

func BenchmarkUnmarshalRawMessage(b *testing.B) {
	jb := []byte(`"\"Fran\""`)
	for i := 0; i < b.N; i++ {
		s := string(jb)
		strconv.Unquote(s)
		// fmt.Println(s)
	}
}
