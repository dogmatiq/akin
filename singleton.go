package akin

import (
	"fmt"
	"reflect"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// Singleton returns a [Set] containing v.
//
// It panics if v's type is not comparable.
func Singleton(v any) Set {
	r := reflectx.ValueOf(v)

	if !r.Comparable() {
		panic(fmt.Sprintf(
			"%s is not comparable",
			renderT(r.Type()),
		))
	}

	return singleton{r}
}

type singleton struct {
	v reflect.Value
}

func (s singleton) String() string {
	return "{ " + renderTV(s.v) + " }"
}

func (s singleton) Contains(v any) bool {
	return reflectx.ValueOf(v).Equal(s.v)
}

func (s singleton) eval(v any) evaluation {
	return newEvaluation(
		s,
		v,
		s.Contains(v),
		isEqual{s.v},
	)
}

type isEqual struct {
	v reflect.Value
}

func (r isEqual) String(inverse bool) string {
	if inverse {
		return "it is not equal to " + renderTV(r.v)
	}
	return "it is equal to " + renderTV(r.v)
}
