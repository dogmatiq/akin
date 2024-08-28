package akin

import (
	"reflect"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// IsEqualTo returns a [Predicate] that is satisfied by any value that compares
// as equal to v using the == operator.
//
// It panics if v is not comparable type.
func IsEqualTo(v any) Predicate {
	r := reflectx.ValueOf(v)

	if !r.Comparable() {
		panic(reflectx.Sprintf("%s is not comparable", r.Type()))
	}

	return equality{r}
}

type equality struct {
	want reflect.Value
}

func (p equality) String() string {
	return reflectx.Sprintf("𝑥 ≡ %s", p.want)
}

func (p equality) Eval(v any) Evaluation {
	r := reflectx.ValueOf(v)

	if r.Type() != p.want.Type() {
		return violated(p, v, "𝑥 is not of type %s", p.want.Type())
	}

	if v != p.want.Interface() {
		return violated(p, v, "𝑥 does not have the same value as %s", p.want)
	}

	return satisfied(p, v, "𝑥 has the same type and value as %s", p.want)
}

func (p equality) Is(q Predicate) bool {
	if q, ok := q.(equality); ok {
		return p.want.Interface() == q.want.Interface()
	}
	return false
}

func (p equality) Reduce() Predicate {
	return p
}
