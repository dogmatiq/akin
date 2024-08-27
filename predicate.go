package akin

import (
	"fmt"
	"reflect"
	"slices"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// A Predicate describes a condition that a value must satisfy.
type Predicate interface {
	fmt.Stringer

	// Eval evaluates v against the predicate.
	Eval(v any) Evaluation

	// Is returns true if p is the same as this predicate.
	Is(p Predicate) bool

	// Simplify returns the simplest equivalent of this predicate.
	//
	// It always returns a non-nil predicate, even if it's the same predicate.
	// It returns true if any simplification was possible.
	Simplify() (Predicate, bool)
}

// To returns a [Predicate] that matches values that are "akin to" the given
// model value.
func To(model any) Predicate {
	v := reflectx.ValueOf(model)
	return fromModel(v)
}

func fromModel(v reflect.Value) Predicate {
	if v.Type().PkgPath() != "" {
		return equalTo{v}
	}

	switch v.Kind() {
	case reflect.Array:
	case reflect.Chan:
	case reflect.Func:

	case reflect.Interface:
		if v.Type() == reflect.TypeFor[any]() && v.IsNil() {
			return AnyOf(IsNil, EqualTo(uintptr(0)))
		}

	case reflect.Map:
	case reflect.Pointer:
	case reflect.Slice:
	case reflect.String:
	case reflect.Struct:
	}

	return convertibleTo{v}
}

func sameConstituents[S ~[]Predicate](a, b S) bool {
	if len(a) != len(b) {
		return false
	}

	b = slices.Clone(b)
	n := len(b)

	for _, p := range a {
		i := slices.IndexFunc(b, p.Is)
		if i == -1 {
			return false
		}

		n--
		b[i] = b[n]
		b = b[:n]
	}

	return true
}

func hasConstituent[S ~[]Predicate](a S, p Predicate) bool {
	return slices.IndexFunc(a, p.Is) != -1
}
