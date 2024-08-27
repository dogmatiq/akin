package akin_test

import (
	"unsafe"

	"github.com/dogmatiq/akin/internal/reflectx"
)

var (
	all = joinCases(
		nilables,
		numbers,
	)
)

var (
	nilables = joinCases(
		nils,
		nonNils,
	)

	nils = []testCase{
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

	nonNils = []testCase{
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

var (
	numbers = joinCases(
		zeroNumbers,
		positiveNumbers,
		negativeNumbers,
	)

	zeroNumbers = []testCase{
		{"zero int", int(0)},
		{"zero int8", int8(0)},
		{"zero int16", int16(0)},
		{"zero int32", int32(0)},
		{"zero int64", int64(0)},
		{"zero uint", uint(0)},
		{"zero uint8", uint8(0)},
		{"zero uint16", uint16(0)},
		{"zero uint32", uint32(0)},
		{"zero uint64", uint64(0)},
		{"zero float32", float32(0)},
		{"zero float64", float64(0)},
		{"zero complex64", complex64(0)},
		{"zero complex128", complex128(0)},
		{"zero uintptr", uintptr(0)},
	}

	positiveNumbers = []testCase{
		{"positive int", int(1)},
		{"positive int8", int8(1)},
		{"positive int16", int16(1)},
		{"positive int32", int32(1)},
		{"positive int64", int64(1)},
		{"positive uint", uint(1)},
		{"positive uint8", uint8(1)},
		{"positive uint16", uint16(1)},
		{"positive uint32", uint32(1)},
		{"positive uint64", uint64(1)},
		{"positive float32", float32(1)},
		{"positive float64", float64(1)},
		{"positive complex64", complex64(1)},
		{"positive complex128", complex128(1)},
		{"positive uintptr", uintptr(1)},
	}

	negativeNumbers = []testCase{
		{"negative int", int(-1)},
		{"negative int8", int8(-1)},
		{"negative int16", int16(-1)},
		{"negative int32", int32(-1)},
		{"negative int64", int64(-1)},
		{"negative float32", float32(-1)},
		{"negative float64", float64(-1)},
		{"negative complex64", complex64(-1)},
		{"negative complex128", complex128(-1)},
	}

	comparable, incomparable = splitCases(
		all,
		func(c testCase) bool {
			return reflectx.ValueOf(c.Value).Comparable()
		},
	)
)

type testCase struct {
	Name  string
	Value any
}

func joinCases(cases ...[]testCase) []testCase {
	seen := map[string]struct{}{}
	var out []testCase

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

func splitCases(
	cases []testCase,
	p func(testCase) bool,
) (in, out []testCase) {
	for _, c := range cases {
		if p(c) {
			in = append(in, c)
		} else {
			out = append(out, c)
		}
	}

	return in, out
}
