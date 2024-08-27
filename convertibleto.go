package akin

import (
	"fmt"
	"reflect"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// ConvertibleTo returns a [Predicate] that is satisfied by any value that
// compares as equal to v when converted to the same type.
func ConvertibleTo(v any) Predicate {
	r := reflectx.ValueOf(v)

	if !r.Comparable() {
		panic(fmt.Sprintf(
			"%s is not comparable",
			renderT(r.Type()),
		))
	}

	return convertibleTo{r}
}

type convertibleTo struct {
	want reflect.Value
}

func (p convertibleTo) String() string {
	return "convertible to " + render(p.want)
}

func (p convertibleTo) Eval(v any) Evaluation {
	got := reflectx.ValueOf(v)

	if got.Equal(p.want) {
		return satisfied(p, v, "the values are equal without requiring conversion")
	}

	gotT := got.Type()
	wantT := p.want.Type()

	if !gotT.ConvertibleTo(wantT) {
		return violated(
			p,
			v,
			"%s is not convertible to %s",
			renderT(gotT),
			renderT(wantT),
		)
	}

	// The logic below is lifted directly from [reflect.Value.CanConvert], so
	// that we can describe the reason for the conversion failure.
	switch {
	case gotT.Kind() == reflect.Slice &&
		wantT.Kind() == reflect.Array:
		if wantT.Len() > got.Len() {
			return violated(p, v, "the value only has %d elements", got.Len())
		}

	case gotT.Kind() == reflect.Slice &&
		wantT.Kind() == reflect.Pointer &&
		wantT.Elem().Kind() == reflect.Array:
		if wantT.Elem().Len() > got.Len() {
			return violated(p, v, "the value only has %d elements", got.Len())
		}
	}

	// Otherwise, we fallback to any conditions we haven't anticipated. If this
	// block is reached it means the logic above is incomplete.
	if !got.CanConvert(wantT) {
		return violated(
			p,
			v,
			"the value can not be converted for an unknown reason",
		)
	}

	conv := got.Convert(wantT)
	if conv.Equal(p.want) {
		return satisfied(p, v, "the values are equal after conversion")
	}

	return violated(p, v, "the values are not equal after conversion")
}

func (p convertibleTo) Simplify() (Predicate, bool) {
	return p, false
}
