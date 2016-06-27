# Generic
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/usk81/generic)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://github.com/usk81/generic/blob/master/LICENSE)
[![Build Status](http://img.shields.io/travis/usk81/generic.svg?style=flat-square)](https://travis-ci.org/usk81/generic)
[![Coverage Status](https://img.shields.io/coveralls/usk81/generic.svg?style=flat-square)](https://coveralls.io/github/usk81/generic?branch=master)
[![Gratipay User](https://img.shields.io/gratipay/user/YusukeKomatsu.svg?style=flat-square)](https://gratipay.com/YusukeKomatsu/)

flexible data type for Go

## Install

standard `go get`:

```
go get -u github.com/usk81/generic
```

## Usage

encode:

```go
package main

import (
	"encoding/json"
	"github.com/usk81/generic"
)

type User struct {
	Name String          `json:"name"`
	Age  generic.TypeInt `json:"age"`
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

## Licence

[MIT](https://github.com/usk81/generic/blob/master/LICENSE)

## Author

[Yusuke Komatsu](https://github.com/usk81)
