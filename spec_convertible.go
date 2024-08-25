package akin

import "fmt"

type losslessConversionSpec struct {
	v Value
}

func (s losslessConversionSpec) Compare(v Value) ComparisonResult {
	if err := requireConvertibleTo(s.v.rtype, v); err != nil {
		return ComparisonResult{err}
	}

	if err := requireConvertibleTo(v.rtype, s.v); err != nil {
		return ComparisonResult{err}
	}

	if !v.rvalue.CanConvert(s.v.rtype) {
		return ComparisonResult{
			fmt.Errorf("cannot convert %#v value to %s", v.value, s.v.rtype),
		}
	}

	// TODO: this doesn't belong here
	if isNegative(s.v) != isNegative(v) {
		return ComparisonResult{notEqualErr(v, s.v)}
	}

	converted := v.rvalue.Convert(s.v.rtype)
	if !converted.Equal(s.v.rvalue) {
		return ComparisonResult{notEqualErr(v, s.v)}
	}

	reverted := converted.Convert(v.rtype)
	if !reverted.Equal(v.rvalue) {
		return ComparisonResult{
			fmt.Errorf(
				"type conversion is lossy, %T(%v) != %T(%v)",
				reverted.Interface(),
				reverted.Interface(),
				v.value,
				v.value,
			),
		}
	}

	return ComparisonResult{}
}
