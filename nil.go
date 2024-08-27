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

func (p isNil) String() string {
	if p {
		return "𝑥 = nil"
	}
	return "𝑥 ≠ nil"
}

func (p isNil) Eval(v any) Evaluation {
	r := reflectx.ValueOf(v)

	good, bad := "nil", "non-nil"
	if !p {
		good, bad = bad, good
	}

	if reflectx.IsNil(r) == bool(p) {
		return satisfied(p, v, "the value is %s", good)
	}

	return violated(p, v, "the value is %s", bad)
}

func (p isNil) Is(q Predicate) bool {
	return p == q
}

func (p isNil) Reduce() Predicate {
	return p
}
