package generic

import (
	"bytes"
	"reflect"
)

// GenericType is the interface used as the basis for generic types
type GenericType interface {
	Valid() bool
	Value() interface{}
	Scan(interface{}) error
	Set(interface{}) error
	Reset()
}

// ErrInvalidGenericValue is used as error in generic types
type ErrInvalidGenericValue struct {
	Value interface{}
}

// ValidFlag
type ValidFlag bool

// Reset resets ValidFlag
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
