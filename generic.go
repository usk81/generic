package generic

import (
	"bytes"
	"database/sql/driver"
	"reflect"
)

// Type is the interface used as the basis for generic types
type Type interface {
	Valid() bool
	Value() (driver.Value, error)
	Scan(interface{}) error
	Set(interface{}) error
	Reset()
}

// ErrInvalidGenericValue is used as error in generic types
type ErrInvalidGenericValue struct {
	Value interface{}
}

// ValidFlag is the flag to check that value is valid
type ValidFlag bool

var nullBytes = []byte("null")

// Reset resets ValidFlag
func (v *ValidFlag) Reset() {
	*v = false
}

// Valid validates the specified value is nil or not
func (v ValidFlag) Valid() bool {
	return bool(v)
}

// Error returns error message
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
