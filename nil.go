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

func (p isNil) String() string {
	if p {
		return "is nil"
	}
	return "is non-nil"
}

func (p isNil) Eval(v any) Evaluation {
	r := reflectx.ValueOf(v)

	if reflectx.IsNil(r) == bool(p) {
		return satisfied(p, v, "the value %s", p)
	}

	return violated(p, v, "the value %s", !p)
}

func (p isNil) Is(q Predicate) bool {
	if q, ok := q.(isNil); ok {
		return p == q
	}
	return false
}

func (p isNil) Simplify() (Predicate, bool) {
	return p, false
}
