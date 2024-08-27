package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestNilness(t *testing.T) {
	for _, c := range nils {
		t.Run(c.Name, func(t *testing.T) {
			AssertIsMember(t, Nil, c.Value)
			AssertIsNotMember(t, NonNil, c.Value)
		})
	}

	for _, c := range nonNils {
		t.Run(c.Name, func(t *testing.T) {
			AssertIsMember(t, NonNil, c.Value)
			AssertIsNotMember(t, Nil, c.Value)
		})
	}
}

func TestNilness_uintptr(t *testing.T) {
	// The zero-valued uintptr is not technically nil, but it is "conceptually"
	// nil, so it is treated as such when using a model nil value, but not when
	// using the actual akin.Nil set.
	AssertIsMember(t, To(nil), uintptr(0))
	AssertIsNotMember(t, Nil, uintptr(0))
	AssertIsMember(t, NonNil, uintptr(0))
}
