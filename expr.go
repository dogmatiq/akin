package akin

import (
	"fmt"
)

// Expr is an expression that describes how a specific [Predicate] or [Value]
// was obtained.
type Expr interface {
	fmt.Stringer
}

type (
	// PredicateExpr is an [Expr] for a predicate variable, such as 𝑷.
	PredicateExpr int
)

var primes = []rune{'′', '″', '‴', '⁗'}

func (e PredicateExpr) String() string {
	prime := 0
	letter := '𝑷' + rune(e)

	for letter > '𝒁' {
		letter -= 26

		if letter > 'P' {
			prime++
		}
	}

	// ⁽x⁾

	if prime == 0 {
		return string(letter)
	}

	if prime < len(primes) {
		return string(letter) + string(primes[prime])
	}

	return string(letter) + superscript(prime)
}
