package akin

import (
	"reflect"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// EqualTo returns a [Predicate] that is satisfied by any value that compares as
// equal to v.
//
// It panics if v is not comparable type.
func EqualTo(v any) Predicate {
	r := reflectx.ValueOf(v)

	if !r.Comparable() {
		panic(reflectx.Sprintf("%s is not comparable", r.Type()))
	}

	return equalTo{r}
}

type equalTo struct {
	want reflect.Value
}

func (p equalTo) String() string {
	return reflectx.Sprintf("𝑥 ≡ %s", p.want)
}

func (p equalTo) Eval(v any) Evaluation {
	r := reflectx.ValueOf(v)

	if r.Type() != p.want.Type() {
		return violated(
			p,
			v,
			"the values have different types",
		)
	}

	if v != p.want.Interface() {
		return violated(p, v, "the values are not equal")
	}

	return satisfied(p, v, "the values are equal")
}

func (p equalTo) Is(q Predicate) bool {
	if q, ok := q.(equalTo); ok {
		return p.want.Interface() == q.want.Interface()
	}
	return false
}

func (p equalTo) Reduce() Predicate {
	return p
}
