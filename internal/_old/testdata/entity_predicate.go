package akin

// // Eval evaluates a value against a predicate, that is 𝑷❨𝒙❩.
// func Eval(p Predicate, x any) Evaluation {
// 	return eval(p, reflectx.ValueOf(x))
// }

// // Reduce returns the simplest form of p.
// func Reduce(p Predicate) Predicate {
// 	return reduce(p)
// }

// Same returns true if a and b are the "same" predicate.
// //
// // Two predicates are the same if they are the same type and have equivalent
// // parameters. This is a kind of weak equality that respects the commutative
// // properties of some predicate types.
// func Same(a, b Predicate) bool {
// 	return same(a, b)
// }

// type isPredicate[T Predicate] struct{}

// var (
// _ = isPredicate[Constant]{}
// _ = isPredicate[Equality]{}
// _ = isPredicate[Equivalence]{}
// _ = isPredicate[Negation]{}
// _ = isPredicate[Nilness]{}
// _ = isPredicate[Or]{}
// _ = isPredicate[TypeEquivalence]{}
// )
