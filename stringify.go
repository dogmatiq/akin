package akin

import (
	"fmt"
	"strings"
)

func stringP(p Predicate) string {
	var s identity
	p.VisitP(&s)
	return string(s)
}

func stringR(r Rationale) string {
	var s identity
	r.VisitR(&s)
	return string(s)
}

func stringA(a Attribute) string {
	var s identity
	a.VisitA(&s)
	return string(s)
}

type (
	identity string
	inverted string
)

func (s *identity) fmt(format string, args ...any) {
	*s = identity(fmt.Sprintf(format, args...))
}

func (s *inverted) fmt(format string, args ...any) {
	*s = inverted(fmt.Sprintf(format, args...))
}

// subscript renders n as a string using unicode subscript characters.
func subscript[
	T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64,
](n T) string {
	digits := []rune(fmt.Sprint(n))

	for i, r := range digits {
		switch r {
		case '-':
			digits[i] = '₋'
		default:
			digits[i] = r - '0' + '₀'
		}
	}

	return string(digits)
}

// parens adds mathematical parentheses to an expression if it contains spaces
// (and does not already have them).
func parens(s string) string {
	if strings.HasPrefix(s, "❨") {
		return s
	}
	if !strings.ContainsAny(s, " ") {
		return s
	}
	return "❨" + s + "❩"
}
