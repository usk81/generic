package generic

import (
	"bytes"
	"reflect"
)

type GenericType interface {
	Valid() bool
	Value() interface{}
	Scan(interface{}) error
	Reset()
}

type ErrInvalidGenericValue struct {
	Value interface{}
}

type ValidFlag bool

func (v *ValidFlag) Reset() {
	*v = false
}

func (v ValidFlag) Valid() bool {
	return bool(v)
}

func (e ErrInvalidGenericValue) Error() string {
	buf := bytes.Buffer{}
	buf.WriteString("invalid value: ")
	t := reflect.TypeOf(e.Value)
	switch t {
	case nil:
		buf.WriteString("(nil)")
	default:
		buf.WriteByte('(')
		buf.WriteString(t.String())
		buf.WriteByte(')')
	}

	return buf.String()
}
