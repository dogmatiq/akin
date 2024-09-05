package akin

import (
	"fmt"
	"reflect"
)

// Is returns a [Predicate] that is satisfied when 𝒙 has the same identity as
// some expected value 𝒆.
func Is(e any) Identity {
	return Identity{E: valueOf(e)}
}

// Identity is a [Predicate] that is satisfied when 𝒙 is identical to the
// expected value 𝒆.
//
// It is notated using ≡ (identical to) and its inverse ≢ (not identical to).
// That is, given 𝑷 ≔ ❨𝒙 ≡ 𝒆❩, then 𝑷❨𝒙❩ = 𝓽 if 𝒙 and 𝒆 are identical.
//
// 𝒙 and 𝒆 can only be considered identical if they share the same [Type].
// That is, given 𝐓 ≔ 𝝉❨𝒆❩, then ❨𝒙 ≡ 𝒆❩ ⟹ 𝒙 ⦂ 𝐓.
//
// When 𝒙 and 𝒆 are the same type, the rules of identity differ based on the
// [reflect.Kind] of 𝐓, as follows:
//
// If 𝐓 is a slice type, then 𝒙 and 𝒆 are identical if they refer to the same
// elements of the same underlying array.
//
// If 𝐓 is a map type, then 𝒙 and 𝒆 are identical if modifications to 𝒙 are
// visible via 𝒆. That is, they refer to the same underlying data structure.
//
// If 𝐓 is an array type, then 𝒙 and 𝒆 are identical if all of their elements
// are identical.
//
// If 𝐓 is a struct type, then 𝒙 and 𝒆 are identical if all of their fields
// are identical.
//
// If 𝐓 is a function type, then 𝒙 and 𝒆 are identical if they point to the
// same code. As documented on [reflect.Value.UnsafePointer], this is not
// necessarily adequate to uniquely identify a single function. This behavior is
// included so that [Identity] is well-defined for all Go types.
//
// For all other kinds, 𝐓 is [comparable] and 𝒙 and 𝒆 are identical if they
// compare as equal. That is, ❨𝒙 ≡ 𝒆❩ ⟺ 𝒙 == 𝒆.
type Identity struct {
	E Value
}

func (p Identity) visit(v PVisitor)     { v.Identity(p) }
func (p Identity) String() string       { return stringP(p, affirmative) }
func (s *stringer) Identity(p Identity) { render(s, "𝒙 {≡|≢} %s", p.E) }

func (e *evaluator) Identity(p Identity) {
	if e.X.Type() != p.E.Type() {
		e.Px = False
		e.R = Ax{
			TypeEq{p.E.Type()},
			false,
		}
		return
	}

	if e.X.isNil() != p.E.isNil() {
		e.Px = False
		e.R = Ax{
			A:  ValueEq{"nil"},
			Ax: e.X.isNil(),
		}
		return
	}

	switch p.E.ref.Kind() {
	case reflect.Slice:
		if e.X.ref.UnsafePointer() != p.E.ref.UnsafePointer() {
			e.Px = False
			e.R = Ax{
				A:  ValueEq{fmt.Sprintf("TODO: ptr(%v != %v)", e.X.ref.UnsafePointer(), p.E.ref.UnsafePointer())},
				Ax: true,
			}
		} else if e.X.ref.Len() != p.E.ref.Len() {
			e.Px = False
			e.R = Ax{
				A:  ValueEq{fmt.Sprintf("TODO: len(%d != %d)", e.X.ref.Len(), p.E.ref.Len())},
				Ax: true,
			}
		} else {
			e.Px = True
			e.R = Ax{
				A:  ValueEq{p.E.String()}, // TODO
				Ax: true,
			}
		}

	case reflect.Array:
		for i := return e.X.ref.Len() {
			if e.X.ref.Index(i).UnsafePointer() != p.E.ref.Index(i).UnsafePointer() {
				e.Px = False
				e.R = Ax{
					A:  ValueEq{fmt.Sprintf("TODO: ptr(%v != %v)", e.X.ref.Index(i).UnsafePointer(), p.E.ref.Index(i).UnsafePointer())},
					Ax: true,
				}
				return
			}
		}

	// case reflect.Struct:

	case reflect.Map, reflect.Func:
		same := e.X.ref.UnsafePointer() == p.E.ref.UnsafePointer()

		e.Px = truth(same)
		e.R = Ax{
			A:  ValueEq{p.E.String()}, // TODO
			Ax: same,
		}

	default:
		same := e.X.nat == p.E.nat

		e.Px = truth(same)
		e.R = Ax{
			A:  ValueEq{p.E.String()},
			Ax: same,
		}
	}
}

func evalIs() {
	if e.X.Type() != p.E.Type() {
		e.Px = False
		e.R = Ax{
			TypeEq{p.E.Type()},
			false,
		}
		return
	}
}
