package akin

import "github.com/dogmatiq/akin/internal/reflectx"

const (
	// IsNil is a [Predicate] that is satisfied by all nil values, regardless of
	// their type.
	IsNil isNil = true

	// IsNonNil is a [Predicate] that is satisfied by all non-nil values,
	// regardless of their type.
	IsNonNil isNil = false
)

type isNil bool

var (
	_ Predicate = IsNil
	_ Predicate = IsNonNil
)

func (s isNil) String() string {
	if s {
		return "is nil"
	}
	return "is non-nil"
}

func (s isNil) Eval(v any) Evaluation {
	r := reflectx.ValueOf(v)

	if reflectx.IsNil(r) == bool(s) {
		return satisfied(s, v, "the value %s", s)
	}

	return violated(s, v, "the value %s", !s)
}
