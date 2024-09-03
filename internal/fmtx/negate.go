package fmtx

// NStringer is an interface for types that can be rendered in their "negated"
// form.
type NStringer interface {
	NString() string
}

// N returns the negated string representation of a value.
func N(v any) string {
	if v, ok := v.(NStringer); ok {
		return v.NString()
	}
	return "¬" + Parens(v)
}
