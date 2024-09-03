package slicex

import "slices"

// Map applies a function to each element of a slice and returns a new slice
// containing the results.
//
// If the function returns false, the element is unchanged. If no elements are
// changed s is returned.
func Map[E any, S ~[]E](s S, f func(E) (E, bool)) S {
	cloned := false

	for i, e := range s {
		if e, ok := f(e); ok {
			if !cloned {
				s = slices.Clone(s)
			}
			s[i] = e
		}
	}

	return s
}
