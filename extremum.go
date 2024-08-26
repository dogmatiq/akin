package akin

import "fmt"

const (
	// Universe is the [Set] of all possible values.
	Universe extremum = true

	// Empty is a [Set] containing no values.
	Empty extremum = false
)

// An extremum is a [Set] that contains either all values or no values.
type extremum bool

var (
	_ Set = Universe
	_ Set = Empty
)

func (e extremum) Contains(any) bool {
	return bool(e)
}

func (e extremum) Eval(any) Membership {
	if e {
		return Membership{
			IsMember: true,
			Reason:   fmt.Sprintf("%s contains all values", e),
		}
	}

	return Membership{
		IsMember: false,
		Reason:   fmt.Sprintf("%s contains no values", e),
	}
}

func (e extremum) String() string {
	if e {
		return "Ω"
	}
	return "∅"
}
