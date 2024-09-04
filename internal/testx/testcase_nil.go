package testx

import "unsafe"

var (
	// NilableCases is a list of test cases that contain values of nilable
	// types.
	NilableCases = Union(
		NilCases,
		NonNilCases,
	)

	// NilCases is the set of cases for nil values.
	NilCases = Cases{
		{"nil interface", nil},
		{"nil unsafe pointer", unsafe.Pointer(nil)},
		{"nil pointer", (*int)(nil)},
		{"nil slice", ([]int)(nil)},
		{"nil map", (map[int]int)(nil)},
		{"nil func", (func())(nil)},
		{"nil chan", (chan int)(nil)},
		{"nil readable chan", (<-chan int)(nil)},
		{"nil writable chan", (chan<- int)(nil)},
	}

	// NonNilCases is the set of cases for non-nil values of nilable types.
	NonNilCases = Cases{
		{"non-nil interface", 123},
		{"non-nil unsafe pointer", unsafe.Pointer(new(int))},
		{"non-nil pointer", new(int)},
		{"empty slice", []int{}},
		{"non-empty slice", []int{1, 2}},
		{"empty map", map[int]int{}},
		{"non-empty map", map[int]int{1: 2}},
		{"non-nil func", func() {}},
		{"non-nil chan", make(chan int)},
		{"non-nil readable chan", make(<-chan int)},
		{"non-nil writable chan", make(chan<- int)},
	}
)
