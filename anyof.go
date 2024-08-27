package akin

import (
	"fmt"
	"strings"
)

// AnyOf returns a [Predicate] that is satisfied if any of the given predicates
// are satisfied.
func AnyOf(predicates ...Predicate) Predicate {
	if len(predicates) == 0 {
		return Nothing
	}

	var filtered []Predicate

	for _, p := range predicates {
		if p == Anything {
			return Anything
		}
		if p != Nothing {
			filtered = append(filtered, p)
		}
	}

	return anyOf(filtered)
}

type anyOf []Predicate

func (p anyOf) String() string {
	var w strings.Builder

	for i, c := range p {
		if i > 0 {
			w.WriteString(" | ")
		}
		w.WriteString(c.String())
	}

	return w.String()
}

func (p anyOf) Eval(v any) Evaluation {
	if len(p) == 0 {
		panic("anyOf must contain at least one predicate")
	}

	var pe Evaluation

	for _, c := range p {
		ce := c.Eval(v)
		pe.Constituents = append(pe.Constituents, ce)

		if ce.IsSatisfied && !pe.IsSatisfied {
			pe.IsSatisfied = true
			pe.Reason = fmt.Sprintf("the constituent predicate %q is satisfied", c)
		}
	}

	if !pe.IsSatisfied {
		pe.Reason = "none of the constituent predicates are satisfied"
	}

	return pe
}
