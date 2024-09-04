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
type Cases []Case

// Case is a single test case.
type Case struct {
	Name string
	X    any
}

// Split splits the set into two sub-sets based on a predicate function.
func (cc Cases) Split(
	p func(reflect.Value) bool,
) (in, ex Cases) {
	for _, c := range cc {
		if p(reflect.ValueOf(c)) {
			in = append(in, c)
		} else {
			ex = append(ex, c)
		}
	}

	return in, ex
}

// Filter returns a new set containing only the test cases that match the
// predicate function.
func (cc Cases) Filter(
	p func(reflect.Value) bool,
) Cases {
	var result Cases

	for _, c := range cc {
		if p(reflect.ValueOf(c)) {
			result = append(result, c)
		}
	}

	return result
}

// Union returns all cases in the given sets of cases.
func Union(cases ...Cases) Cases {
	var result Cases
	seen := map[string]struct{}{}

	for _, cc := range cases {
		for _, c := range cc {
			if _, ok := seen[c.Name]; !ok {
				result = append(result, c)
				seen[c.Name] = struct{}{}
			}
		}
	}

	return result
}
