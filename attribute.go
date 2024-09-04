package akin

// An Attribute describes some indivisible quality of a value.
//
// Unlike a [Predicate], an [Attribute] cannot be evaluated against a Go value
// directly. This is because attributes can represent abstract qualities. For
// example, the [ValueEq] attribute can be used to represent "the number one",
// without any specific Go type, such as int or float64.
//
// Within documentation and strings, 𝛂 (mathematical bold italic small alpha)
// is used to represent an attribute. When discussing multiple attributes, the
// letters 𝛃, 𝜸, and so on, are used.
type Attribute interface {
	visit(AVisitor)
}

// AVisitor is an algorithm with logic specific to each [Attribute] type.
type AVisitor interface {
	TypeEq(TypeEq)
	ValueEq(ValueEq)
}
