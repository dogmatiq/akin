package akin

import "github.com/dogmatiq/akin/internal/reflectx"

const (
	// IsNil is a [Predicate] that is satisfied by all nil values, regardless of
	// their type.
	IsNil nilness = true

	// IsNonNil is a [Predicate] that is satisfied by all non-nil values,
	// regardless of their type.
	IsNonNil nilness = false
)

type nilness bool

func (p nilness) String() string {
	if p {
		return "𝑥 = nil"
	}
	return "𝑥 ≠ nil"
}

func (p nilness) Eval(v any) Evaluation {
	r := reflectx.ValueOf(v)

	good, bad := "nil", "non-nil"
	if !p {
		good, bad = bad, good
	}

	// TODO: expand this language to read "x is a nil pointer", "x is a nil
	// slice", etc. Perhaps call out empty but non-nil slices, maps, etc.
	if reflectx.IsNil(r) == bool(p) {
		return satisfied(p, v, "𝑥 is %s", good)
	}

	return violated(p, v, "𝑥 is %s", bad)
}

func (p nilness) Is(q Predicate) bool {
	return p == q
}

func (p nilness) Reduce() Predicate {
	return p
}
