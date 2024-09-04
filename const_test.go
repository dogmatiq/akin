package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	. "github.com/dogmatiq/akin/internal/testx"
)

func TestConst(t *testing.T) {
	AssertIsReduced(t, Top)
	AssertIsReduced(t, Bottom)

	for n, x := range AllCases {
		t.Run(n, func(t *testing.T) {
			AssertTrue(t, Top, x)
			AssertFalse(t, Bottom, x)
		})
	}
}
