package akin

import "reflect"

// IsA returns a [PredicateP] that is satisfied when 𝒙 is an implementation of T.
//
// If T is not an interface, the predicate matches only if 𝒙 is exactly T.
func IsA[T any]() Typehood {
	return Typehood{
		T: typeFor[T](),
	}
}

// Typehood is a [PredicateP] that is satisfied when 𝒙 implements [Type] 𝐓.
//
// It is notated using ∈ (element of) and its inverse ∉ (not an element of). For
// example, given 𝑷 ≔ ❨𝒙 ∈ 𝐓❩, then 𝑷❨𝒙❩ = 𝓽 if 𝒙 has a [Type] that implements 𝐓.
type Typehood struct {
	T Type
}

func (p Typehood) acceptPredicateVisitor(v PredicateVisitor) { v.VisitTypehood(p) }
func (p Typehood) acceptAssertionVisitor(v AssertionVisitor) { v.VisitTypehood(p) }
func (p Typehood) String() string                            { return predicateToString(p) }

func (r *predicateRenderer) VisitTypehood(p Typehood) {
	r.Render("𝒙 {∈|∉} %s", p.T)
}

func (e *evaluator) VisitTypehood(p Typehood) {
	t := e.Value.Type()

	if p.T.ref.Kind() == reflect.Interface {
		e.Result = asResult(t.ref.Implements(p.T.ref))
	} else {
		e.Result = asResult(t.ref == p.T.ref)
	}

	e.Rationale = IntrinsicRationale{
		TypeEq{t},
		e.PredicateExpr, // TODO
		e.Value,
		e.ValueExpr,
		true,
	}
}
