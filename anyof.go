package akin

import (
	"fmt"
	"strings"
)

// AnyOf returns a [Predicate] that is satisfied if any of the given predicates
// are satisfied.
func AnyOf(predicates ...Predicate) Predicate {
	if len(predicates) == 0 {
		panic("akin.AnyOf(): at least one predicate must be provided")
	}
	return anyOf(predicates)
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
	var pe Evaluation

	for _, c := range p {
		ce := c.Eval(v)
		pe.Constituents = append(pe.Constituents, ce)

		if ce.IsSatisfied && !pe.IsSatisfied {
			pe.IsSatisfied = true
			pe.Reason = fmt.Sprintf("the constituent %q is satisfied", c)
		}
	}

	if !pe.IsSatisfied {
		pe.Reason = "none of the constituents are satisfied"
	}

	return pe
}

func (p anyOf) Simplify() (Predicate, bool) {
	var filtered []Predicate

	for _, c := range p {
		if c == Anything {
			return Anything, true
		}

		if c != Nothing {
			filtered = append(filtered, c)
		}
	}

	if len(filtered) == len(p) {
		return p, false
	}

	if len(filtered) == 0 {
		return Nothing, true
	}

	if len(filtered) == 1 {
		return filtered[0], true
	}

	return anyOf(filtered), true
}
