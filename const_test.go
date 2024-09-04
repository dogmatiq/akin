package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	. "github.com/dogmatiq/akin/internal/testx"
)

func TestConst(t *testing.T) {
	AssertIsReduced(t, Top)
	AssertIsReduced(t, Bottom)

	for _, c := range AllCases {
		t.Run(c.Name, func(t *testing.T) {
			AssertTrue(t, Top, c.X)
			AssertFalse(t, Bottom, c.X)
		})
	}
}
