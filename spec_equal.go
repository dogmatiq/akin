package akin

import "reflect"

type equalTo struct {
	v reflect.Value
}

func (s equalTo) Compare(v reflect.Value) ComparisonResult {
	if err := requireType(s.v.Type(), v); err != nil {
		return ComparisonResult{Error: err}
	}
	if err := requireEquality(s.v, v); err != nil {
		return ComparisonResult{Error: err}
	}
	return ComparisonResult{}
}
