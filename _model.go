package akin

// To returns a [Predicate] that matches values that are "akin to" the model
// value m.
func To(m any) Predicate {
	panic("not implemented")
	// v := reflectx.ValueOf(m)
	// return model(v)
}

// func model(v reflect.Value) Predicate {
// 	if v.Type().PkgPath() != "" {
// 		return Equal{v}
// 	}

// 	switch v.Kind() {

// 	case reflect.Interface:
// 		if v.Type() == reflect.TypeFor[any]() && v.IsNil() {
// 			return Or(IsNil, IsEqualTo(uintptr(0)))
// 		}

// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
// 		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
// 		reflect.Float32, reflect.Float64,
// 		reflect.Uintptr:
// 		if c, ok := reflectx.ToComplex128(v); ok {
// 			return or{
// 				Equivalent{v},
// 				IsEquivalentTo(c),
// 			}
// 		}

// 	case reflect.Complex64, reflect.Complex128:
// 		if f, ok := reflectx.ToFloat64(v); ok {
// 			return or{
// 				Equivalent{v},
// 				IsEquivalentTo(f),
// 			}
// 		}

// 		// case reflect.Array:
// 		// case reflect.Chan:
// 		// case reflect.Func:
// 		// case reflect.Map:
// 		// case reflect.Pointer:
// 		// case reflect.Slice:
// 		// case reflect.String:
// 		// case reflect.Struct:
// 	}

// 	return Equivalent{v}
// }
