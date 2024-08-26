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
