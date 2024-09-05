package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	. "github.com/dogmatiq/akin/internal/testx"
)

func TestConst(t *testing.T) {
	AssertRationale(t, Top, 42, "𝒙 ≔ 42, 𝑷 ≔ ⊤ ∴ 𝑷❨𝒙❩ = 𝓽 ∵ 𝑷 is constant")
	AssertRationale(t, Bottom, 42, "𝒙 ≔ 42, 𝑷 ≔ ⊥ ∴ 𝑷❨𝒙❩ = 𝓯 ∵ 𝑷 is constant")

	for _, c := range AllCases {
		t.Run(c.Name, func(t *testing.T) {
			AssertTrue(t, Top, c.X)
			AssertFalse(t, Bottom, c.X)
		})
	}

}
