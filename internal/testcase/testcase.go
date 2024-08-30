package testcase

import (
	"reflect"

	"github.com/dogmatiq/akin/internal/reflectx"
)

var (
	// All is the set of all test cases.
	All = Union(
		Nilable,
		Number,
	)

	comparable, incomparable = All.Split(reflect.Value.Comparable)

	// Comparable is the set of cases with values that can be compared using the
	// == operator.
	Comparable = comparable

	// Incomparable is the set of cases with values that cannot be compared
	// using the == operator.
	Incomparable = incomparable
)

// Set is a set of related test cases.
type Set map[string]any

// Split splits the set into two sub-sets based on a predicate function.
func (s Set) Split(
	p func(reflect.Value) bool,
) (in, ex Set) {
	for n, x := range s {
		s := &ex
		if p(reflectx.ValueOf(x)) {
			s = &in
		}

		if *s == nil {
			*s = Set{}
		}

		(*s)[n] = x
	}

	return in, ex
}

// Filter returns a new set containing only the test cases that match the
// predicate function.
func (s Set) Filter(
	p func(reflect.Value) bool,
) Set {
	var result Set

	for n, x := range s {
		if p(reflectx.ValueOf(x)) {
			if result == nil {
				result = Set{}
			}
			result[n] = x
		}
	}

	return result
}

// Union returns a set containing all of the cases in the given sets.
func Union(sets ...Set) Set {
	var result Set

	for _, s := range sets {
		for n, x := range s {
			if result == nil {
				result = Set{}
			}
			result[n] = x
		}
	}

	return result
}
