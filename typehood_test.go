package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	. "github.com/dogmatiq/akin/internal/testx"
)

type errT struct{}

func (errT) Error() string {
	return "<error>"
}

func TestTypehood(t *testing.T) {
	t.Run("concrete type", func(t *testing.T) {
		p := IsA[int]()

		AssertRationale(t, p, 42, "𝒙 ≔ 42, 𝑷 ≔ ❨𝒙 ∈ int❩ ∴ 𝑷❨𝒙❩ = 𝓽 ∵ 𝒙 ⦂ int")

		AssertTrue(t, p, 0)
		AssertFalse(t, p, uint(0))
		AssertFalse(t, p, float64(0))
		AssertFalse(t, p, "0")
	})

	t.Run("interface type", func(t *testing.T) {
		p := IsA[error]()

		AssertRationale(t, p, errT{}, "𝒙 ≔ akin_test.errT(<error>), 𝑷 ≔ ❨𝒙 ∈ error❩ ∴ 𝑷❨𝒙❩ = 𝓽 ∵ 𝒙 ⦂ akin_test.errT")

		AssertTrue(t, p, errT{})
		AssertFalse(t, p, 1)
		AssertFalse(t, p, nil)
	})
}
