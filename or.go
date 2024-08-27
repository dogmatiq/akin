package akin

import (
	"fmt"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// Or returns a [Predicate] that is satisfied if any of the given predicates are
// satisfied.
func Or(predicates ...Predicate) Predicate {
	return or(predicates)
}

type or []Predicate

func (p or) String() string {
	switch len(p) {
	case 0:
		return "❨⁇ ∨ ⁇❩"
	case 1:
		return reflectx.Sprintf("❨%s ∨ ⁇❩", p[0])
	default:
		return "❨" + reflectx.SprintList(" ∨ ", p...) + "❩"
	}
}

func (p or) Eval(v any) Evaluation {
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

func (p or) Is(q Predicate) bool {
	if q, ok := q.(or); ok {
		return sameConstituents(p, q)
	}
	return false
}

func (p or) Reduce() Predicate {
	var reduced or

	for _, c := range p {
		flattened, ok := c.(or)
		if !ok {
			flattened = or{c}
		}

		for _, c := range flattened {
			c = c.Reduce()

			if c.Is(Top) {
				return Top
			}

			if c.Is(Bottom) {
				continue
			}

			if hasConstituent(reduced, c) {
				continue
			}

			reduced = append(reduced, c)
		}
	}

	if len(reduced) == 0 {
		return Bottom
	}

	if len(reduced) == 1 {
		return reduced[0]
	}

	return reduced
}
