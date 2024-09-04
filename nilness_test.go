package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	. "github.com/dogmatiq/akin/internal/testx"
)

func TestNilness(t *testing.T) {
	AssertIsReduced(t, IsNil)
	AssertIsReduced(t, IsNonNil)

	for n, x := range NilCases {
		t.Run(n, func(t *testing.T) {
			AssertTrue(t, IsNil, x)
			AssertFalse(t, IsNonNil, x)
		})
	}

	for n, x := range NonNilCases {
		t.Run(n, func(t *testing.T) {
			AssertFalse(t, IsNil, x)
			AssertTrue(t, IsNonNil, x)
		})
	}
}
