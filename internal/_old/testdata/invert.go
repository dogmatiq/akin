package akin

import "fmt"

func invert(p Predicate) Predicate {
	i := &inverter{}
	VisitPredicate(p, i)

	if i.Q == nil {
		panic(fmt.Sprintf("inverter.Visit%T() did not set Q", p))
	}

	return i.Q
}

type inverter struct {
	Q Predicate
}
