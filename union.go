package akin

import (
	"strings"
)

// Union returns the [Set] that is the union of all given sets.
func Union(sets ...Set) Set {
	filtered := make([]Set, 0, len(sets))

	for _, s := range sets {
		if s == Universe {
			return Universe
		}

		if s != Empty {
			filtered = append(filtered, s)
		}
	}

	if len(sets) == 0 {
		return Empty
	}

	return union(sets)
}

type union []Set

func (s union) String() string {
	var w strings.Builder

	for i, set := range s {
		if i > 0 {
			w.WriteString(" ∪ ")
		}
		w.WriteString(set.String())
	}

	return w.String()
}

func (s union) Contains(v any) bool {
	for _, set := range s {
		if set.Contains(v) {
			return true
		}
	}
	return false
}

func (s union) eval(v any) evaluation {
	e := evaluation{
		Set:   s,
		Value: v,
	}

	for _, set := range s {
		x := set.eval(v)

		if x.IsMember {
			e.IsMember = true
		}

		for _, p := range x.Predicates {
			e.Predicates = append(e.Predicates, p)
		}
	}

	return e
}
