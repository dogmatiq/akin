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
