package akin

import (
	"reflect"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// OfType returns a [Predicate] that is satisfied by any value of type T.
func OfType[T any]() Predicate {
	return ofType{reflect.TypeFor[T]()}
}

type ofType struct {
	want reflect.Type
}

func (s ofType) String() string {
	return reflectx.Sprintf("𝑥 ⦂ %s", s.want)
}

func (s ofType) Eval(v any) Evaluation {
	t := reflect.TypeOf(v)
	if t == s.want {
		return satisfied(s, v, "the value is of the expected type")
	}
	return violated(s, v, "the value is not of the expected type")
}

func (s ofType) Is(q Predicate) bool {
	if q, ok := q.(ofType); ok {
		return s.want == q.want
	}
	return false
}

func (s ofType) Reduce() Predicate {
	return s
}
