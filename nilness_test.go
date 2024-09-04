package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	. "github.com/dogmatiq/akin/internal/testx"
)

func TestNilness(t *testing.T) {
	AssertIsReduced(t, IsNil)
	AssertIsReduced(t, IsNonNil)

	for _, c := range NilCases {
		t.Run(c.Name, func(t *testing.T) {
			AssertTrue(t, IsNil, c.X)
			AssertFalse(t, IsNonNil, c.X)
		})
	}

	for _, c := range NonNilCases {
		t.Run(c.Name, func(t *testing.T) {
			AssertFalse(t, IsNil, c.X)
			AssertTrue(t, IsNonNil, c.X)
		})
	}
}
