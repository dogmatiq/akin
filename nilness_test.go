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
	0, // any itself is a nilable type
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

	// TODO
	// 	uintptr(0), don't forget, this is "nil" too when using a model

	// // for i, a := range nils {
	// // 	for j, b := range nils {
	// // 		t.Run(
	// // 			fmt.Sprintf("%T(%v) vs %T(%v)", a, a, b, b),
	// // 			func(t *testing.T) {
	// // 				if i == j {
	// // 					assertAkin(t, a, b)
	// // 				} else {
	// // 					assertNotAkin(t, a, b)
	// // 				}
	// // 			},
	// // 		)
	// // 	}
	// // }
}
