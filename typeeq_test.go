package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	. "github.com/dogmatiq/akin/internal/testx"
)

func TestIs(t *testing.T) {
	p := Is[int]()

	AssertIsReduced(t, p)

	AssertTrue(t, p, 0)
	AssertFalse(t, p, uint(0))
	AssertFalse(t, p, float64(0))
	AssertFalse(t, p, "0")
}
