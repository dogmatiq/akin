package akin

// Tribool is a tri-state boolean.
//
// It can be either [False], [True], any other value is considered
// indeterminate.
type Tribool uint8

const (
	// False is the false value for a [Tribool].
	False Tribool = iota

	// True is the true value for a [Tribool].
	True
)

// IsIndeterminate returns true if the is neither [False] nor [True].
func (t Tribool) IsIndeterminate() bool {
	return t != True && t != False
}

// IsTrue returns true if t is [True].
func (t Tribool) IsTrue() bool {
	return t == True
}

// IsFalse returns true if t is [False].
func (t Tribool) IsFalse() bool {
	return t == False
}

// AsBool returns the boolean value of t, or panics if t is indeterminate.
func (t Tribool) AsBool() bool {
	if t.IsIndeterminate() {
		panic("cannot collapse indeterminate tribool")
	}
	return t.IsTrue()
}

func (t Tribool) String() string {
	switch t {
	case False:
		return "false"
	case True:
		return "true"
	default:
		return "indeterminate"
	}
}
