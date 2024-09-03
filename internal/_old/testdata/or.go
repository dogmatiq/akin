package akin

import (
	"fmt"
)

// Or is a [Predicate] 𝑷 that is satisfied when 𝒙 satisfies any one of its 𝑛
// constituent predicates 𝐐ₙ.
//
// If there are no constituent predicates (𝑛 = 0), then 𝑷❨𝒙❩ is false.
type Or []Predicate

// Format implements the [fmt.Formatter] interface.
func (p Or) Format(f fmt.State, v rune) {
	format(p, f, v)
}

func (p Or) hide() any {
	type T = Or
	type Or T
	return T(p)
}

func (p Or) formal() string {
	switch len(p) {
	case 0:
		return parensf("𝐐 ∨ 𝐐")
	case 1:
		return parensf("%s ∨ 𝐐", p[0])
	default:
		return parens(join(" ∨ ", p...))
	}
}

func (p Or) human() string {
	switch len(p) {
	case 0:
		return "𝑷 has no constituent predicates"
	case 1:
		return renderf("𝒙 satisfies %s", p[0])
	default:
		return "𝒙 satisfies " + join2(", ", " or ", p...)
	}
}

func (p Or) visitPredicate(v PredicateVisitor) {
	v.VisitOrPredicate(p)
}

func (i *inverter) VisitOrPredicate(p Or) {
	i.Q = Not(p)
}

func (r *reducer) VisitOrPredicate(p Or) {
	// TODO
	r.Q = p
}

func (e *evaluator) VisitOrPredicate(p Or) {
	if len(p) == 0 {
		e.SetReason(false, NoConstituents{p})
		return
	}

	for _, q := range p {
		eq := eval(q, e.X)

		if eq.IsSatisfied {
			e.SetReason(true, ConstituentSatisfied{eq})
			return
		}
	}

	e.SetReason(false, AllConstituentsViolated{p})
}

// func (p or) Eval(v any) Evaluation {
// 	e := buildEvaluation(p, v)

// 	isSatisfied := false

// 	for _, c := range p {
// 		ce := c.Eval(v)

// 		if ce.IsSatisfied {
// 			isSatisfied = true
// 			e.For(ConstituentEvaluation{ce})
// 		} else {
// 			e.Against(ConstituentEvaluation{ce})
// 		}
// 	}

// 	if len(p) == 0 {
// 		e.Against(PredicateIsConstant{p, false})
// 	} else {
// 		e.Against(PredicateHasNoConstituents{p})
// 	}

// 	return e.Build(isSatisfied)
// }

// func (p or) Is(q Predicate) bool {
// 	if q, ok := q.(or); ok {
// 		return samePredicates(p, q)
// 	}
// 	return false
// }

// func (p or) Reduce() Predicate {
// 	var reduced or

// 	for _, c := range p {
// 		flattened, ok := c.(or)
// 		if !ok {
// 			flattened = or{c}
// 		}

// 		for _, c := range flattened {
// 			c = c.Reduce()

// 			if c.Is(Top) {
// 				return Top
// 			}

// 			if c.Is(Bottom) {
// 				continue
// 			}

// 			if containsPredicate(reduced, c) {
// 				continue
// 			}

// 			reduced = append(reduced, c)
// 		}
// 	}

// 	if len(reduced) == 0 {
// 		return Bottom
// 	}

// 	if len(reduced) == 1 {
// 		return reduced[0]
// 	}

// 	return reduced
// }
