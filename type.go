package akin

import (
	"reflect"
)

// OfType returns a [Predicate] that is satisfied by any value of type T.
func OfType[T any]() Predicate {
	return ofType{reflect.TypeFor[T]()}
}

type ofType struct {
	want reflect.Type
}

func (s ofType) String() string {
	return "of type " + renderT(s.want)
}

func (s ofType) Eval(v any) Evaluation {
	t := reflect.TypeOf(v)
	if t == s.want {
		return satisfied(s, v, "the value is of the expected type")
	}
	return violated(s, v, "the value is not of the expected type")
}

func (s ofType) Simplify() (Predicate, bool) {
	return s, false
}
