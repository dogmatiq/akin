package akin

import (
	"fmt"
)

// Expr is a psuedo-syntactic expression that describes how a specific
// [Predicate] or [Value] was "reached" during evaluation.
type Expr interface {
	fmt.Stringer
	acceptExprVisitor(ExprVisitor)
}

type (
	// NamedPredicateExpr is an [Expr] that identifies a predicate, such as 𝑷.
	NamedPredicateExpr uint

	// NamedValueExpr is an [Expr] that identifies a value, such as 𝒙.
	NamedValueExpr uint

	// PrimeExpr is an [Expr] that identifies the prime of another [Expr].
	//
	// For example, the first-order prime of 𝑷 is written as 𝑷′
	PrimeExpr struct {
		// Base is expression to which the prime is applied.
		Base Expr

		// N is the zero-based order of the prime. That is, when N is zero the
		// expression represents the first-order prime 𝑷′, and so on.
		//
		// The 1ˢᵗ through 4ᵗʰ order primes are represented by the unicode prime
		// characters, for example 𝑷′, 𝑷″, 𝑷‴, and 𝑷⁗. The 5ᵗʰ prime and
		// above are represented using superscript numbers, for example 𝑷⁽⁵⁾.
		N uint
	}
)

const (
	defaultPredicateExpr NamedPredicateExpr = 0
	defaultValueExpr     NamedValueExpr     = 0
)

// An ExprVisitor is an algorithm that applies different logic for each [Expr]
// type.
type ExprVisitor interface {
	VisitNamedPredicateExpr(NamedPredicateExpr)
	VisitNamedValueExpr(NamedValueExpr)
	VisitPrimeExpr(PrimeExpr)
}

func (e NamedPredicateExpr) acceptExprVisitor(v ExprVisitor) { v.VisitNamedPredicateExpr(e) }
func (e NamedValueExpr) acceptExprVisitor(v ExprVisitor)     { v.VisitNamedValueExpr(e) }
func (e PrimeExpr) acceptExprVisitor(v ExprVisitor)          { v.VisitPrimeExpr(e) }

func (e NamedPredicateExpr) String() string {
	return "𝑷"
}

func (e NamedValueExpr) String() string {
	return "𝒙"
}

func (e PrimeExpr) String() string {
	switch e.N {
	case 0:
		return fmt.Sprintf("%s′", e.Base)
	case 1:
		return fmt.Sprintf("%s″", e.Base)
	case 2:
		return fmt.Sprintf("%s‴", e.Base)
	case 3:
		return fmt.Sprintf("%s⁗", e.Base)
	default:
		return e.Base.String() + super("⁽%d⁾", e.N)
	}
}

// // LitExpr is a [ValueExpr] describing a literal value.
// LitExpr struct {
// 	Value Value
// }

// // IndexExpr is a [ValueExpr] describing the an element within a slice or
// // array.
// IndexExpr struct {
// 	Indexable ValueVarExpr
// 	Index     int
// }

// // KeyExpr is a [ValueExpr] describing the [Value] associated with a
// // specific key within a map.
// KeyExpr struct {
// 	Map ValueVarExpr
// 	Key Value
// }

// // FieldExpr is a [ValueExpr] describing a specific field within a struct.
// FieldExpr struct {
// 	Receiver ValueVarExpr
// 	Field    string
// }

// // DerefExpr is a [ValueExpr] describing the dereferenced value of a
// // pointer.
// DerefExpr struct {
// 	Pointer ValueVarExpr
// }
// )
//
// var primes = []rune{'′', '″', '‴', '⁗'}

// func (e PredicateVarExpr) String() string {
// 	prime := 0
// 	letter := '𝑷' + rune(e)

// 	for letter > '𝒁' {
// 		letter -= 26

// 		if letter > 'P' {
// 			prime++
// 		}
// 	}

// 	// ⁽x⁾

// 	if prime == 0 {
// 		return string(letter)
// 	}

// 	if prime < len(primes) {
// 		return string(letter) + string(primes[prime])
// 	}

// 	return string(letter) + superscript(prime)
// }

// // func (e VarExpr) visit(v ValueExprVisitor)   { v.VarExpr(e) }
// // func (e LitExpr) visit(v ValueExprVisitor)   { v.LitExpr(e) }
// // func (e IndexExpr) visit(v ValueExprVisitor) { v.IndexExpr(e) }
// // func (e KeyExpr) visit(v ValueExprVisitor)   { v.KeyExpr(e) }
// // func (e FieldExpr) visit(v ValueExprVisitor) { v.FieldExpr(e) }
// // func (e DerefExpr) visit(v ValueExprVisitor) { v.DerefExpr(e) }

// // func (e VarExpr) String() string {
// // 	return e.Name
// // }

// // func (e LitExpr) String() string {
// // 	return e.Value.String()
// // }

// // func (e IndexExpr) String() string {
// // 	return fmt.Sprintf("%s[%d]", e.Indexable, e.Index)
// // }

// // func (e KeyExpr) String() string {
// // 	return fmt.Sprintf("%s[%s]", e.Map, e.Key)
// // }

// // func (e FieldExpr) String() string {
// // 	r := e.Receiver

// // 	if d, ok := e.Receiver.(DerefExpr); ok {
// // 		// Render x.Field instead of (*x).Field, since the explicit dereference
// // 		// is not necessary in Go's syntax.
// // 		r = d.Pointer
// // 	}

// // 	return fmt.Sprintf("%s.%s", r, e.Field)
// // }

// // func (e DerefExpr) String() string {
// // 	return fmt.Sprintf("(*%s)", e.Pointer)
// // }
