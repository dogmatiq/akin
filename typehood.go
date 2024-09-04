package akin

import "reflect"

// IsA returns a [Predicate] that is satisfied when 𝒙 is an implementation of T.
//
// If T is not an interface, the predicate matches only if 𝒙 is exactly T.
func IsA[T any]() Typehood {
	return Typehood{
		T: typeFor[T](),
	}
}

// Typehood is a [Predicate] that is satisfied when 𝒙 implements [Type] 𝐓.
//
// It is notated using ∈ (element of) and its inverse ∉ (not an element of). For
// example, given 𝑷 ≔ ❨𝒙 ∈ 𝐓❩, then 𝑷❨𝒙❩ = 𝓽 if 𝒙 has a [Type] that implements 𝐓.
type Typehood struct {
	T Type
}

func (p Typehood) visit(v PVisitor)     { v.Typehood(p) }
func (p Typehood) String() string       { return stringP(p, affirmative) }
func (s *stringer) Typehood(p Typehood) { render(s, "𝒙 {∈|∉} %s", p.T) }

func (e *evaluator) Typehood(p Typehood) {
	t := e.X.Type()

	if p.T.ref.Kind() == reflect.Interface {
		e.Px = truth(t.ref.Implements(p.T.ref))
	} else {
		e.Px = truth(t.ref == p.T.ref)
	}

	e.R = Ax{
		A:  TypeEq{t},
		Ax: true,
	}
}
