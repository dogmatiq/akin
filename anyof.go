package akin

import (
	"fmt"
	"strings"
)

// AnyOf returns a [Predicate] that is satisfied if any of the given predicates
// are satisfied.
func AnyOf(predicates ...Predicate) Predicate {
	return anyOf(predicates)
}

type anyOf []Predicate

func (p anyOf) String() string {
	var w strings.Builder

	w.WriteString("(")
	for i, c := range p {
		if i > 0 {
			w.WriteString(" | ")
		}
		w.WriteString(c.String())
	}
	w.WriteString(")")

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
		if len(p) == 0 {
			pe.Reason = "there are no constituents"
		} else {
			pe.Reason = "none of the constituents are satisfied"
		}
	}

	return pe
}

func (p anyOf) Is(q Predicate) bool {
	if q, ok := q.(anyOf); ok {
		sameConstituents(p, q)
	}
	return false
}

func (p anyOf) Simplify() (Predicate, bool) {
	if len(p) == 0 {
		return Nothing, true
	}

	var (
		simple     anyOf
		simplified bool
	)

	for _, c := range p {
		c, ok := c.Simplify()
		if ok {
			simplified = true
		}

		if c.Is(Anything) {
			return Anything, true
		}

		if constituents, ok := c.(anyOf); ok {
			simple.merge(constituents...)
			simplified = true
		} else {
			if !simple.merge(c) {
				simplified = true
			}
		}
	}

	if !simplified {
		return p, false
	}

	if len(simple) == 0 {
		return Nothing, true
	}

	if len(simple) == 1 {
		return simple[0], true
	}

	return simple, true
}

func (p *anyOf) merge(constituents ...Predicate) bool {
	all := true

	for _, c := range constituents {
		if c.Is(Nothing) || hasConstituent(*p, c) {
			all = false
		} else {
			*p = append(*p, c)
		}
	}

	return all
}
