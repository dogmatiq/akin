package akin

import (
	"fmt"
	"reflect"
)

// Singleton returns a [Set] containing v.
//
// It panics if v's type is not comparable.
func Singleton(v any) Set {
	r := valueOf(v)

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

func (s singleton) Contains(v any) bool {
	return valueOf(v).Equal(s.v)
}

func (s singleton) eval(v any) membership {
	if s.Contains(v) {
		return membership{
			IsMember: true,
			For:      []string{"is equal to " + renderV(s.v)},
		}
	}

	return membership{
		IsMember: false,
		Against:  []string{"is not equal to " + renderV(s.v)},
	}
}

func (s singleton) String() string {
	return "{ " + renderV(s.v) + " }"
}
