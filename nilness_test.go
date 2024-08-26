package akin_test

import (
	"testing"
	"unsafe"

	. "github.com/dogmatiq/akin"
)

// nils contains a nil value for each kind that may be nilable.
var nils = []any{
	nil,
	unsafe.Pointer(nil),
	(*int)(nil),
	([]int)(nil),
	(map[int]int)(nil),
	(func())(nil),
	(chan int)(nil),
}

// nonNils contains a non-nil value for each kind that may be nilable.
var nonNils = []any{
	"", // any itself is a nilable type
	uintptr(0),
	unsafe.Pointer(new(int)),
	new(int),
	[]int{},
	map[int]int{},
	func() {},
	make(chan int),
}

func TestNilness(t *testing.T) {
	for _, v := range nils {
		AssertContains(t, Nil, v)
		AssertNotContains(t, NonNil, v)
	}

	for _, v := range nonNils {
		AssertContains(t, NonNil, v)
		AssertNotContains(t, Nil, v)
	}
}
