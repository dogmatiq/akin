package akin

import (
	"reflect"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// InhabitsType returns a [Predicate] that is satisfied by any value of type T.
func InhabitsType[T any]() Predicate {
	return inhabitance{reflect.TypeFor[T]()}
}

type inhabitance struct {
	want reflect.Type
}

func (p inhabitance) String() string {
	return reflectx.Sprintf("𝑥 ⦂ %s", p.want)
}

func (p inhabitance) Eval(v any) Evaluation {
	t := reflect.TypeOf(v)
	if t == p.want {
		return satisfied(p, v, "𝑥 is of type %s", p.want)
	}
	return violated(p, v, "𝑥 is not of type %s", p.want)
}

func (p inhabitance) Is(q Predicate) bool {
	if q, ok := q.(inhabitance); ok {
		return p.want == q.want
	}
	return false
}

func (p inhabitance) Reduce() Predicate {
	return p
}
