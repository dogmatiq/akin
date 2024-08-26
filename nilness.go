package akin

import (
	"reflect"
)

var (
	// Nil is the [Set] of all nil values.
	Nil nilness = true

	// NonNil is the [Set] of all non-nil values.
	NonNil nilness = false
)

type nilness bool

var _ Set = Nil

func (s nilness) Contains(v any) bool {
	return bool(s) == isNil(v)
}

func (s nilness) Eval(v any) Membership {
	if s {
		return Membership{
			IsMember: s.Contains(v),
			Reason:   s.str(),
		}
	}
	return Membership{
		IsMember: s.Contains(v),
		Reason:   s.str(),
	}
}

func (s nilness) String() string {
	return "{" + s.str() + "}"
}

func (s nilness) str() string {
	if s {
		return "is nil"
	}
	return "is non-nil"
}

func isNil(v any) bool {
	if v == nil {
		return true
	}
	r := reflect.ValueOf(v)
	return r.IsNil()
}
