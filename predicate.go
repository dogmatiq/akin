package akin

import "fmt"

// A Predicate describes some criteria that a Go value may (or may not) satisfy.
//
// Within documentation and strings, 𝑷 (mathematical bold italic capital P) is
// used to represent a predicate. 𝒙 (mathematical bold italic small X)
// represents a value that is tested against the criteria described by 𝑷. When
// discussing multiple predicates, the letters 𝐐, 𝑹, and so on, are used.
//
// To determine if 𝒙 satisfies 𝑷, we "evaluate 𝑷 of 𝒙", written 𝑷❨𝒙❩. The
// result of an evaluation is one of [True], [False] or [Undefined], denoted
// 𝓽, 𝓯 and 𝓾, respectively (mathematical bold script small letters).
//
// The [Eval] function is used to evaluate 𝑷❨𝒙❩.
type Predicate interface {
	visit(PVisitor)
}

// PVisitor is an algorithm with logic specific to each [Predicate] type.
type PVisitor interface {
	Const(Const)
	Nilness(Nilness)
	Typehood(Typehood)
	Identity(Identity)
}

// Eval evaluates 𝑷❨𝒙❩.
func Eval(p Predicate, x any) (Truth, Rationale) {
	e := &evaluator{
		P: p,
		X: valueOf(x, VarExpr{"𝒙"}),
	}

	p.visit(e)

	if e.R == nil {
		panic(fmt.Sprintf(
			"%s ≔ %s, 𝑷 ≔ %s ∴ 𝑷❨%s❩ = %s has no rationale",
			e.X.Expr(),
			e.X,
			parens(stringP(e.P, affirmative)),
			e.X.Expr(),
			e.Px,
		))
	}

	return e.Px, Px(*e)
}

type evaluator Px
