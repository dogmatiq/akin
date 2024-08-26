package akin

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

func (e extremum) eval(any) membership {
	if e {
		return membership{
			IsMember: true,
			For:      []string{"everything is a member of " + e.String()},
		}
	}

	return membership{
		IsMember: false,
		Against:  []string{"nothing is a member of " + e.String()},
	}
}

func (e extremum) String() string {
	if e {
		return "Ω"
	}
	return "∅"
}
