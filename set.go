package akin

// A Set describes a (possibly infinite) set of Go values.
type Set interface {
	// Contains returns true if the set contains v.
	Contains(v any) bool

	// eval evaluates v's membership in the set.
	eval(v any) membership

	String() string
}

// membership describes a value's membership to a specific [Set].
type membership struct {
	// IsMember is true if the value is a member of the set.
	IsMember bool

	// For and Against are the reasons for and against the value being a member
	// of the set, respectively.
	For, Against []string
}
