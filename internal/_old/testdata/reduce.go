package akin

import "fmt"

func reduce(p Predicate) Predicate {
	r := &reducer{}
	VisitPredicate(p, r)

	if r.Q == nil {
		panic(fmt.Sprintf("inverter.Visit%T() did not set Q", p))
	}

	return r.Q
}

type reducer struct {
	Q Predicate
}
