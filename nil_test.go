package akin_test

import (
	"testing"
	"unsafe"
)

func Test_nil(t *testing.T) {
	nils := []any{
		nil,
		(*int)(nil),
		uintptr(0),
		unsafe.Pointer(nil),
		(chan int)(nil),
		([]int)(nil),
		(map[int]int)(nil),
		(func())(nil),
	}

	for i, a := range nils {
		for j, b := range nils {
			if i == j {
				assertAkin(t, a, b)
			} else {
				assertNotAkin(t, a, b)
			}
		}
	}
}
