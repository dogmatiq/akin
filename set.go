package akin

// A Set describes a (possibly infinite) set of Go values.
type Set interface {
	// Contains returns true if the set contains v.
	Contains(v any) bool

	// Eval evaluates v's membership in the set.
	Eval(v any) Membership

	String() string
}

// Membership describes a value's membership to a specific [Set].
type Membership struct {
	// IsMember is true if the value is a member of the set.
	IsMember bool

	// Reason is a human-readable explanation of why the value is or isn't a
	// member of the set.
	Reason string
}

// fn is a [Set] implemented as an eval function.
type fn func(value) error

func (s fn) Contains(v any) bool {
	return s.Eval(v).IsMember
}

func (s fn) Eval(v any) Membership {
	if err := s(valueOf(v)); err != nil {
		return Membership{IsMember: false, Reason: err.Error()}

	}
	return Membership{
		IsMember: true,
	}
}
