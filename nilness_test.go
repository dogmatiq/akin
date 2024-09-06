package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	. "github.com/dogmatiq/akin/internal/testx"
)

func TestNilness(t *testing.T) {
	AssertRationale(t, IsNil, 42, "𝒙 ≔ 42, 𝑷 ≔ ❨𝒙 = nil❩ ∴ 𝑷❨𝒙❩ = 𝓯 ∵ 𝒙 ⦂ int")
	AssertRationale(t, IsNonNil, 42, "𝒙 ≔ 42, 𝑷 ≔ ❨𝒙 ≠ nil❩ ∴ 𝑷❨𝒙❩ = 𝓽 ∵ 𝒙 ⦂ int")
	AssertRationale(t, IsNil, (*int)(nil), "𝒙 ≔ (*int)(nil), 𝑷 ≔ ❨𝒙 = nil❩ ∴ 𝑷❨𝒙❩ = 𝓽 ∵ 𝒙 ≍ nil")
	AssertRationale(t, IsNonNil, (*int)(nil), "𝒙 ≔ (*int)(nil), 𝑷 ≔ ❨𝒙 ≠ nil❩ ∴ 𝑷❨𝒙❩ = 𝓯 ∵ 𝒙 ≍ nil")

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
