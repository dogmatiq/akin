package akin

import (
	"fmt"
	"reflect"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// IsEquivalentTo returns a [Predicate] that is satisfied when 𝒙 == v, after
// conversion to the same type.
func IsEquivalentTo(v any) Equivalence {
	r := reflectx.ValueOf(v)

	// TODO: make the predicate always fail if v is not comparable, rather than
	// panicking.
	if !r.Comparable() {
		panic(renderf("%s is not comparable", r.Type()))
	}

	return Equivalence{r}
}

// Equivalence is a [Predicate] that is satisfied when 𝒙 compares as equal to a
// fixed value when using the == operator, after conversion to the same type.
type Equivalence struct {
	V reflect.Value
}

// Format implements the [fmt.Formatter] interface.
func (p Equivalence) Format(s fmt.State, verb rune) {
	format(p, s, verb)
}

func (p Equivalence) hide() any {
	type T = Equivalence
	type Equivalent T
	return Equivalent(p)
}
func (p Equivalence) formal(neg bool) string {
	return "𝒙" + choose(neg, " ≆ ", " ≅ ") + render(p.V)
}

func (p Equivalence) human(neg bool) string {
	return "𝒙 is" + insert(neg, "not") + " equivalent to " + render(p.V) + " when converted to " + render(p.V.Type())
}

func (p Equivalence) visitPredicate(v PredicateVisitor) {
	v.VisitEquivalentPredicate(p)
}

func (i *inverter) VisitEquivalentPredicate(p Equivalence) {
	i.Q = Not(p)
}

func (r *reducer) VisitEquivalentPredicate(p Equivalence) {
	r.Q = p
}

func (e *evaluator) VisitEquivalentPredicate(p Equivalence) {
	e.IsSatisfied = e.T == p.V.Type() && e.X.Interface() == p.V.Interface()
}

// func (p Equivalent) Eval(v any) Evaluation {
// 	got := reflectx.ValueOf(v)
// 	gotT := got.Type()
// 	wantT := p.V.Type()

// 	e := buildEvaluation(p, v)

// 	isSameType := gotT == wantT
// 	isConvertible := got.CanConvert(wantT)
// 	isSameValue := isSameType && v == p.V.Interface()

// 	if !isSameValue {
// 		if isConvertible {
// 			converted := got.Convert(wantT)

// 			if converted.CanConvert(gotT) {
// 				reverted := converted.Convert(gotT)

// 				isSameValue = converted.Interface() == p.V.Interface() &&
// 					reverted.Interface() == v &&
// 					reflectx.IsNeg(got) == reflectx.IsNeg(converted)
// 			}
// 		}
// 	}

// 	return e.
// 		Property(TypeIsProp{wantT}, isSameType).
// 		Property(ValueIsConvertibleTo{wantT}, isConvertible).
// 		Property(ValueIs{p.V}, isSameValue).
// 		Build(isSameValue)
// }
