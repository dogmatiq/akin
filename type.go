package akin

import (
	"reflect"
)

// HasType returns the [Set] of all values of type T.
func HasType[T any]() Set {
	return hasType{reflect.TypeFor[T]()}
}

// hasType is the [Set] of all values with a specific type.
type hasType struct {
	t reflect.Type
}

func (s hasType) Contains(v any) bool {
	return reflect.TypeOf(v) == s.t
}

func (s hasType) Eval(v any) Membership {
	if s.Contains(v) {
		return Membership{
			IsMember: true,
			Reason:   "has type " + s.t.String(),
		}
	}

	return Membership{
		IsMember: false,
		Reason:   "does not have type " + s.t.String(),
	}
}

func (s hasType) String() string {
	return "{has type " + s.t.String() + "}"
}
