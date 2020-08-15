# Generic
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/usk81/generic)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://github.com/usk81/generic/blob/master/LICENSE)
[![Build Status](http://img.shields.io/travis/usk81/generic.svg?style=flat-square)](https://travis-ci.org/usk81/generic)
[![Coverage Status](https://img.shields.io/coveralls/usk81/generic.svg?style=flat-square)](https://coveralls.io/github/usk81/generic?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/usk81/generic)](https://goreportcard.com/report/github.com/usk81/generic)

flexible data type for Go

support: Go 1.8+

## Install

standard `go get`:

```
go get -u github.com/usk81/generic/v2
```

## Usage

encode/decode:

```go
package main

import (
	"encoding/json"
	"github.com/usk81/generic"
)

type User struct {
	Name String      `json:"name"`
	Age  generic.Int `json:"age"`
}

var user1 User
u1 := []byte(`{"name":"Daryl Dixon","age":"40"}`)
json.Unmarshal([]byte(u1), &user1)
b, _ := json.Marshal(user1)
Println(string(b))
// {"name":"Daryl Dixon","age":40}

var user2 User
u2 := []byte(`{"name":"Rick Grimes"}`)
json.Unmarshal([]byte(u2), &user2)
b, _ := json.Marshal(user2)
Println(string(b))
// {"name":"Rick Grimes","age":null}
```

set:

```go
package main

import (
	"fmt"
	"github.com/usk81/generic"
)

func main() {
	v := 1.0

	var tb generic.Bool
	tb.Set(v)
	vb := tb.Weak()
	fmt.Printf("%v, (%T)\n", vb, vb)
	// true, (bool)

	var tf generic.Float
	tf.Set(v)
	vf := tf.Weak()
	fmt.Printf("%v, (%T)\n", vf, vf)
	// 1, (float64)

	var ti generic.Int
	ti.Set(v)
	vi := ti.Weak()
	fmt.Printf("%v, (%T)\n", vi, vi)
	// 1, (int64)

	var ts generic.String
	ts.Set(v)
	vs := ts.Weak()
	fmt.Printf("%v, (%T)\n", vs, vs)
	// 1, (string)

	var tt generic.Time
	tt.Set(v)
	vt := tt.Weak()
	fmt.Printf("%v, (%T)\n", vt.UTC(), vt)
	// 1970-01-01 09:00:01 +0900 JST, (time.Time)

	var tu generic.Uint
	tu.Set(v)
	vu := tu.Weak()
	fmt.Printf("%v, (%T)\n", vu, vu)
	// 1, (uint64)
}
```

## Benchmarks

### Marshal

#### Bool

| version | requests | /op | B/op | allocs/op |
|---|---|---|---|---|
| 1.0.0  | 5000000 | 240 ns | 185 | 3 |
| 2.0.0  | 200000000 | 6.69 ns | 0 | 0 |

#### Float

| version | requests | /op | B/op | allocs/op |
|---|---|---|---|---|
| 1.0.0  | 3000000 | 425 ns | 192 | 3 |
| 2.0.0  | 5000000 | 260 ns | 64 | 3 |

#### Int

| version | requests | /op | B/op | allocs/op |
|---|---|---|---|---|
| 1.0.0  | 5000000 | 265 ns | 192 | 3 |
| 2.0.0  | 20000000 | 70.5 ns | 16 | 2 |

#### String (small)

| version | requests | /op | B/op | allocs/op |
|---|---|---|---|---|
| 1.0.0  | 3000000 | 382 ns | 200 | 3 |
| 2.0.0  | 20000000 | 89.0 ns | 128 | 2 |

#### String (Large)

| version | requests | /op | B/op | allocs/op |
|---|---|---|---|---|
| 1.0.0  | 1000000 | 1056 ns | 776 | 4 |
| 2.0.0  | 5000000 | 237 ns | 896 | 2 |

#### Time

| version | requests | /op | B/op | allocs/op |
|---|---|---|---|---|
| 1.0.0  | 1000000 | 1122 ns | 360 | 5 |
| 2.0.0  | 3000000 | 401 ns | 48 | 1 |

#### TimestampMS

| version | requests | /op | B/op | allocs/op |
|---|---|---|---|---|
| 1.0.0  | 20000000 | 97.9 ns | 32 | 2 |
| 2.0.0  | 20000000 | 91.2 ns | 32 | 2 |

#### TimestampNano

| version | requests | /op | B/op | allocs/op |
|---|---|---|---|---|
| 1.0.0  | 10000000 | 114 ns | 64 | 2 |
| 2.0.0  | 10000000 | 112 ns | 64 | 2 |

#### Timestamp

| version | requests | /op | B/op | allocs/op |
|---|---|---|---|---|
| 1.0.0  | 20000000 | 88.4 ns | 32 | 2 |
| 2.0.0  | 20000000 | 86.7 ns | 32 | 2 |

#### Uint

| version | requests | /op | B/op | allocs/op |
|---|---|---|---|---|
| 1.0.0  | 5000000 | 277 ns | 192 | 3 |
| 2.0.0  | 20000000 | 64.2 ns | 16 | 2 |

## Licence

[MIT](https://github.com/usk81/generic/blob/master/LICENSE)

## Author

[Yusuke Komatsu](https://github.com/usk81)
