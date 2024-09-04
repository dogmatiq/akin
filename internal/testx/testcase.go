package testx

import (
	"reflect"
)

var (
	// AllCases is the set of all test cases.
	AllCases = Union(
		NilableCases,
		NumericCases,
	)

	comparable, incomparable = AllCases.Split(reflect.Value.Comparable)

	// ComparableCases is the set of cases with values that can be compared
	// using the == operator.
	ComparableCases = comparable

	// IncomparableCases is the set of cases with values that cannot be compared
	// using the == operator.
	IncomparableCases = incomparable
)

// Cases is a set of related test cases.
type Cases map[string]any

// Split splits the set into two sub-sets based on a predicate function.
func (s Cases) Split(
	p func(reflect.Value) bool,
) (in, ex Cases) {
	for n, x := range s {
		s := &ex
		if p(reflect.ValueOf(x)) {
			s = &in
		}

		if *s == nil {
			*s = Cases{}
		}

		(*s)[n] = x
	}

	return in, ex
}

// Filter returns a new set containing only the test cases that match the
// predicate function.
func (s Cases) Filter(
	p func(reflect.Value) bool,
) Cases {
	var result Cases

	for n, x := range s {
		if p(reflect.ValueOf(x)) {
			if result == nil {
				result = Cases{}
			}
			result[n] = x
		}
	}

	return result
}

// Union returns a set containing all of the cases in the given sets.
func Union(sets ...Cases) Cases {
	var result Cases

	for _, s := range sets {
		for n, x := range s {
			if result == nil {
				result = Cases{}
			}
			result[n] = x
		}
	}

	return result
}
