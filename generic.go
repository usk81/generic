package generic

import (
	"bytes"
	"reflect"
)

// GenericType
type GenericType interface {
	Valid() bool
	Value() interface{}
	Scan(interface{}) error
	Set(interface{}) error
	Reset()
}

// ErrInvalidGenericValue
type ErrInvalidGenericValue struct {
	Value interface{}
}

// ValidFlag
type ValidFlag bool

// Reset
func (v *ValidFlag) Reset() {
	*v = false
}

// Valid validates the specified value is nil or not.
func (v ValidFlag) Valid() bool {
	return bool(v)
}

// Error
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
