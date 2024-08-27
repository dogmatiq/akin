package akin

import (
	"reflect"
)

// DomainFor returns the [Set] of all values of type T.
func DomainFor[T any]() Set {
	return domain{reflect.TypeFor[T]()}
}

// domainOf is the [Set] of all values of type T.
type domain struct {
	t reflect.Type
}

func (s domain) String() string {
	return "{ x | x has type " + renderT(s.t) + " }"
}

func (s domain) Contains(v any) bool {
	return reflect.TypeOf(v) == s.t
}

func (s domain) eval(v any) evaluation {
	return newEvaluation(
		s,
		v,
		s.Contains(v),
		hasType{s.t},
	)
}

type hasType struct {
	t reflect.Type
}

func (r hasType) String(isMember bool) string {
	if isMember {
		return "does not have type " + renderT(r.t)
	}
	return "has type " + renderT(r.t)
}
