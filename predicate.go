package akin

import (
	"fmt"
	"slices"
)

// A Predicate describes a condition that a value must satisfy.
type Predicate interface {
	fmt.Stringer

	// Eval evaluates v against the predicate.
	Eval(v any) Evaluation

	// Is returns true if p is equivalent to this predicate.
	//
	// It does NOT reduce the predicates to their simplest form before
	// comparison. For example [Top] is NOT equivalent to [Or]([Top], [Bottom]),
	// even though both predicates match all values.
	Is(p Predicate) bool

	// Reduce returns the simplest equivalent of this predicate.
	Reduce() Predicate
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

func hasConstituent[S ~[]Predicate](p S, q Predicate) bool {
	return slices.IndexFunc(p, q.Is) != -1
}

// equal returns true if q has type P and compares as equal to p.
func equal[P interface {
	comparable
	Predicate
}](p P, q Predicate) bool {
	if q, ok := q.(P); ok {
		return p == q
	}
	return false
}
