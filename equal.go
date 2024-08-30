package akin

import (
	"fmt"
	"reflect"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// IsEqualTo returns a [Predicate] that is satisfied when 𝑥 == v.
func IsEqualTo(v any) Equal {
	r := reflectx.ValueOf(v)

	// TODO: make the predicate always fail if v is not comparable, rather than
	// panicking.
	if !r.Comparable() {
		panic(sprintf("%s is not comparable", r.Type()))
	}

	return Equal{r}
}

// Equal is a [Predicate] and a [Property] that is satisfied when 𝑥 compares as
// equal to a fixed value when using the == operator. This implies that the
// values have the same type.
type Equal struct {
	V reflect.Value
}

// Format implements the [fmt.Formatter] interface.
func (p Equal) Format(s fmt.State, v rune) {
	format(p, s, v)
}

func (p Equal) hide() any {
	type T = Equal
	type Equal T
	return Equal(p)
}

func (p Equal) formal() string {
	return sprintf("𝑥 ≡ %s", p.V)
}

func (p Equal) human() string {
	return sprintf("𝑥 is equal to %s", p.V)
}

func (p Equal) visitPredicate(v PredicateVisitor) {
	v.VisitEqualPredicate(p)
}

func (e *evaluator) VisitEqualPredicate(p Equal) {
	e.IsSatisfied = e.X.Interface() == p.V.Interface()
}

// 	wantT := p.V.Type()

// 	isSameType := got.Type() == wantT
// 	isSameValue := isSameType && v == p.V.Interface()

// 	return buildEvaluation(p, v).
// 		Property(TypeIsProp{wantT}, isSameType).
// 		Property(ValueIs{p.V}, isSameValue).
// 		Build(isSameValue)
// }

// func (p Equal) Is(q Predicate) bool {
// 	if q, ok := q.(Equal); ok {
// 		return p.V.Interface() == q.V.Interface()
// 	}
// 	return false
// }

// func (p Equal) Reduce() Predicate {
// 	return p
// }
