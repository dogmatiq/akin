package akin

import (
	"reflect"
)

// OfType returns a [Predicate] that is satisfied by any value of type T.
func OfType[T any]() Predicate {
	return ofType{reflect.TypeFor[T]()}
}

type ofType struct {
	t reflect.Type
}

func (s ofType) String() string {
	return "of type " + renderT(s.t)
}

func (s ofType) Eval(v any) Evaluation {
	t := reflect.TypeOf(v)
	if t == s.t {
		return satisfied(s, v, "the value has the same type")
	}
	return violated(s, v, "the value has a different type")
}
