package akin

import (
	"fmt"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// A Set describes a (possibly infinite) set of Go values.
type Set interface {
	fmt.Stringer

	// Contains returns true if the set contains v.
	Contains(v any) bool

	// eval evaluates v's membership in the set.
	eval(v any) evaluation
}

// evaluation is the result evaluating a value's membership in a [Set].
type evaluation struct {
	Set        Set
	Value      any
	IsMember   bool
	Predicates []predicateEvaluation
}

type predicateEvaluation struct {
	Set       Set
	Predicate predicate
	Satisfied bool
}

func newEvaluation(
	s Set,
	v any,
	isMember bool,
	p predicate,
) evaluation {
	if isMember {
		return evaluation{
			Set:        s,
			Value:      v,
			IsMember:   true,
			Predicates: []predicateEvaluation{{s, p, true}},
		}
	}

	return evaluation{
		Set:        s,
		Value:      v,
		IsMember:   false,
		Predicates: []predicateEvaluation{{s, p, false}},
	}
}

func (e evaluation) String() string {
	symbol := "∉"
	if e.IsMember {
		symbol = "∈"
	}

	return fmt.Sprintf(
		"%s %s %s",
		renderTV(reflectx.ValueOf(e.Value)),
		symbol,
		e.Set,
	)
}

type predicate interface {
	String(inverse bool) string
}
