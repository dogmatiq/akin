package akin

// import "fmt"

// // Convertible is a [Set] that matches values that can be converted to the
// // specification value.
// type Convertible struct {
// 	Value value
// }

// // Compare returns a successful comparison result if the value compares as equal to s.Value when converted to the same type, and
// func (s Convertible) Compare(v value) result {
// 	if err := requireConvertibleTo(s.Value.dynamic, v); err != nil {
// 		return result{err}
// 	}

// 	if err := requireConvertibleTo(v.dynamic, s.Value); err != nil {
// 		return result{err}
// 	}

// 	if !v.r.CanConvert(s.Value.dynamic) {
// 		return result{
// 			fmt.Errorf("cannot convert %#v value to %s", v.v, s.Value.dynamic),
// 		}
// 	}

// 	// TODO: this doesn't belong here
// 	if isNegative(s.Value) != isNegative(v) {
// 		return result{notEqualErr(v, s.Value)}
// 	}

// 	converted := v.r.Convert(s.Value.dynamic)
// 	if !converted.Equal(s.Value.r) {
// 		return result{notEqualErr(v, s.Value)}
// 	}

// 	reverted := converted.Convert(v.dynamic)
// 	if !reverted.Equal(v.r) {
// 		return result{
// 			fmt.Errorf(
// 				"type conversion is lossy, %T(%v) != %T(%v)",
// 				reverted.Interface(),
// 				reverted.Interface(),
// 				v.v,
// 				v.v,
// 			),
// 		}
// 	}

// 	return result{}
// }
