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

// Case is a single test case.
type Case struct {
	Name  string
	Value any
}

// Set is a set of related [Case] values.
type Set []Case

// Split splits the set into two sub-sets based on a predicate function.
func (s Set) Split(
	p func(reflect.Value) bool,
) (in, ex Set) {
	for _, c := range s {
		if p(reflectx.ValueOf(c.Value)) {
			in = append(in, c)
		} else {
			ex = append(ex, c)
		}
	}
	return in, ex
}

// Filter returns a new set containing only the test cases that match the
// predicate function.
func (s Set) Filter(
	p func(reflect.Value) bool,
) Set {
	s, _ = s.Split(p)
	return s
}

// Union returns a set containing all of the cases in the given sets.
func Union[S ~[]Case](cases ...S) Set {
	seen := map[string]struct{}{}
	var out []Case

	for _, cc := range cases {
		for _, c := range cc {
			if _, ok := seen[c.Name]; !ok {
				seen[c.Name] = struct{}{}
				out = append(out, c)
			}
		}
	}

	return out
}
