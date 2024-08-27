package akin

import (
	"reflect"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// EquivalentTo returns a [Predicate] that is satisfied by any value that
// compares as equal to v when converted to the same type, provided it can be
// converted back to the original type without loss of information.
func EquivalentTo(v any) Predicate {
	r := reflectx.ValueOf(v)

	if !r.Comparable() {
		panic(reflectx.Sprintf("%s is not comparable", r.Type()))
	}

	return equivalent{r}
}

type equivalent struct {
	want reflect.Value
}

func (p equivalent) String() string {
	return reflectx.Sprintf("𝑥 = %s", p.want)
}

func (p equivalent) Eval(v any) Evaluation {
	got := reflectx.ValueOf(v)

	gotT := got.Type()
	wantT := p.want.Type()

	if gotT == wantT {
		if v == p.want.Interface() {
			return satisfied(p, v, "the values are equal")
		}
		return violated(p, v, "the values have the same type but are not equal")
	}

	// Check that we can convert the value to the predicate's value.
	if !gotT.ConvertibleTo(wantT) {
		return violated(
			p,
			v,
			"%s values cannot be converted to %s",
			gotT,
			wantT,
		)
	}

	// Check that we can convert the converted value back to its original type.
	if !wantT.ConvertibleTo(gotT) {
		return violated(
			p,
			v,
			"%s values cannot be converted to %s",
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
			"this specific value can not be converted to %s",
			wantT,
		)
	}

	converted := got.Convert(wantT)

	if converted.Interface() != p.want.Interface() {
		return violated(p, v, "the values are not equal after type conversion")
	}

	// HACK(jmalloc): Do we really want numeric-specific comparisons here?
	if reflectx.IsNeg(got) != reflectx.IsNeg(converted) {
		return violated(p, v, "the values have different numeric signs after type conversion")
	}

	reverted := converted.Convert(gotT)
	if reverted.Interface() != v {
		return violated(p, v, "information is lost when converting back to the original type")
	}

	return satisfied(p, v, "the values are equal after type conversion")
}

func (p equivalent) Is(q Predicate) bool {
	if q, ok := q.(equivalent); ok {
		return p.want.Interface() == q.want.Interface()
	}
	return false
}

func (p equivalent) Reduce() Predicate {
	return p
}
