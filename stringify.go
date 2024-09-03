package akin

import "fmt"

func stringP(p Predicate) string {
	var s stringer
	p.VisitP(&s)
	return string(s)
}

func stringR(r Rationale) string {
	var s stringer
	r.VisitR(&s)
	return string(s)
}

func stringA(a Attribute) string {
	var s stringer
	a.VisitA(&s)
	return string(s)
}

type (
	stringer        string
	negatedStringer string
)

func (s *stringer) fmt(format string, args ...any) {
	*s = stringer(fmt.Sprintf(format, args...))
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
