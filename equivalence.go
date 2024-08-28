package akin

import (
	"reflect"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// IsEquivalentTo returns a [Predicate] that is satisfied by any value that
// compares as equal to v when converted to the same type, provided it can be
// converted back to the original type without loss of information.
func IsEquivalentTo(v any) Predicate {
	r := reflectx.ValueOf(v)

	if !r.Comparable() {
		panic(reflectx.Sprintf("%s is not comparable", r.Type()))
	}

	return equivalence{r}
}

type equivalence struct {
	want reflect.Value
}

func (p equivalence) String() string {
	return reflectx.Sprintf("𝑥 ≅ %s", p.want)
}

func (p equivalence) Eval(v any) Evaluation {
	got := reflectx.ValueOf(v)

	gotT := got.Type()
	wantT := p.want.Type()

	if gotT == wantT {
		if v == p.want.Interface() {
			return satisfied(p, v, "𝑥 has the same type and value as %s", p.want)
		}
		return violated(p, v, "𝑥 does not have the same value as %s", p.want)
	}

	// Check that we can convert the value to the predicate's value.
	if !gotT.ConvertibleTo(wantT) {
		return violated(
			p,
			v,
			"values of type %s cannot be converted to %s",
			gotT,
			wantT,
		)
	}

	// Check that we can convert the converted value back to its original type.
	if !wantT.ConvertibleTo(gotT) {
		return violated(
			p,
			v,
			"values of type %s cannot be converted to %s",
			wantT,
			gotT,
		)
	}

	// Otherwise, check for any unanticipated reasons the conversion might fail.
	// If this condition passes it means the logic above is incomplete.
	//
	// At the time of writing (2024-08-28), the only additional logic in
	// [reflect.Value.CanConvert] is related to converting slices to arrays.
	// Since conversion back to a slice from an array is not possible these
	// cases are already convered by the checks above.
	if !got.CanConvert(wantT) {
		return violated(
			p,
			v,
			"𝑥 can not be converted to type %s",
			wantT,
		)
	}

	converted := got.Convert(wantT)

	if converted.Interface() != p.want.Interface() {
		return violated(p, v, "𝑥 does not have the same value as %s after type conversion", p.want)
	}

	if reflectx.IsNeg(got) != reflectx.IsNeg(converted) {
		return violated(p, v, "the value of 𝑥 changed after type conversion")
	}

	reverted := converted.Convert(gotT)
	if reverted.Interface() != v {
		return violated(p, v, "the value of 𝑥 changed after type conversion")
	}

	return satisfied(p, v, "𝑥 has the same value as %s after type conversion", p.want)
}

func (p equivalence) Is(q Predicate) bool {
	if q, ok := q.(equivalence); ok {
		return p.want.Interface() == q.want.Interface()
	}
	return false
}

func (p equivalence) Reduce() Predicate {
	return p
}
