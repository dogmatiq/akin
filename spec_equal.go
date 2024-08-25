package akin

type equalitySpec struct {
	v Value
}

func (s equalitySpec) Compare(v Value) ComparisonResult {
	if err := requireEqualTo(s.v, v); err != nil {
		return ComparisonResult{Error: err}
	}
	return ComparisonResult{}
}
