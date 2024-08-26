package akin

import (
	"reflect"
)

const (
	// Nil is the [Set] of all nil values.
	Nil nilness = true

	// NonNil is the [Set] of all non-nil values.
	NonNil nilness = false
)

type nilness bool

var (
	_ Set = Nil
	_ Set = NonNil
)

func (s nilness) Contains(v any) bool {
	return bool(s) == isNil(valueOf(v))
}

func (s nilness) eval(v any) membership {
	if s.Contains(v) {
		return membership{
			IsMember: true,
			For:      []string{"it " + s.str()},
		}
	}
	return membership{
		IsMember: false,
		Against:  []string{"it " + (!s).str()},
	}
}

func (s nilness) String() string {
	return "{ " + s.str() + " }"
}

func (s nilness) str() string {
	if s {
		return "is nil"
	}
	return "is non-nil"
}

func isNil(v reflect.Value) bool {
	switch v.Kind() {
	default:
		return false
	case
		reflect.Interface,
		reflect.Pointer,
		reflect.UnsafePointer,
		reflect.Slice,
		reflect.Map,
		reflect.Func,
		reflect.Chan:
		return v.IsNil()
	}
}
