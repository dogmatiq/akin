package akin

import (
	"fmt"
)

// Or is a [Predicate] 𝑷 that is satisfied when 𝑥 satisfies any one of its 𝑛
// constituent predicates 𝐐ₙ.
//
// If there are no constituent predicates (𝑛 = 0), then 𝑷❨𝑥❩ is false.
type Or []Predicate

func (p Or) visitPredicate(v PredicateVisitor) { v.VisitOrPredicate(p) }

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
		return "❨𝐐 ∨ 𝐐❩"
	case 1:
		return sprintf("❨%s ∨ 𝐐❩", p[0])
	default:
		return "❨" + join(" ∨ ", p...) + "❩"
	}
}

func (p Or) human() string {
	switch len(p) {
	case 0:
		return "𝑷 has no constituent predicates"
	case 1:
		return sprintf("𝑥 satisfies %s", p[0])
	default:
		return "𝑥 satisfies " + join2(", ", " or ", p...)
	}
}

func (e *evaluator) VisitOrPredicate(p Or) {
	if len(p) == 0 {
		e.SetReason(false, NoConstituents{p})
		return
	}

	for _, qn := range p {
		en := eval(qn, e.X)

		if en.IsSatisfied {
			e.SetReason(true, ConstituentSatisfied{en})
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
