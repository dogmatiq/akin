package akin

import (
	"fmt"
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
		panic(fmt.Sprintf(
			"%s is not comparable",
			renderT(r.Type()),
		))
	}

	return equalTo{r}
}

type equalTo struct {
	v reflect.Value
}

func (p equalTo) String() string {
	return "equal to " + render(p.v)
}

func (p equalTo) Eval(v any) Evaluation {
	r := reflectx.ValueOf(v)

	if r.Type() != p.v.Type() {
		return violated(
			p,
			v,
			"the value has a different type",
		)
	}

	if !r.Equal(p.v) {
		return violated(p, v, "the value is not equal")
	}

	return satisfied(p, v, "the value is equal")
}
