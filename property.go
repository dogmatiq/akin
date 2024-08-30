package akin

import "fmt"

// A Property describes some specific quality of 𝑥 that is a component in
// determining whether a [Predicate] is satisfied.
//
// Many [Predicate] types are composed of a single [Property] implemented by the
// same type. For example, the [TypeEquivalence] predicate is a direct assertion that 𝑥
// has the property of having a specific type. In contrast, the [Equivalence]
// predicate is composed of separate [TypeEquivalence] and [ValueEquivalence] properties.
type Property interface {
	formatter

	visitProperty(PropertyVisitor)
	inverse() string
}

type isProperty[T interface {
	fmt.Formatter
	Property
}] struct{}

var (
	_ = isProperty[TypeEquivalence]{}
	_ = isProperty[ValueEquivalence]{}
)

// A PropertyVisitor encapsulates logic specific to each [Property] type.
type PropertyVisitor interface {
	VisitEqualProperty(Equal)
	VisitTypeEquivalenceProperty(TypeEquivalence)
	VisitValueEquivalenceProperty(ValueEquivalence)
}

// VisitProperty calls the appropriate method on v for the given property.
func VisitProperty(p Property, v PropertyVisitor) {
	p.visitProperty(v)
}
