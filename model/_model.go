package akin

import "reflect"

// Like returns the [Set] of values that are like the given model value.
func Like(model any) Set {
	return like(reflect.ValueOf(model))
}

func like(model reflect.Value) Set {
	// switch model.Kind() {
	// 	case reflect.Slice:
	// 		return compileSlice(spec)

	// 	case reflect.Map:
	// 		return compileMap(spec)

	// 	case reflect.Func:
	// 		return compileFunc(spec)

	// 		// 		// case reflect.Uintptr:
	// 		// 		// case reflect.Complex64:
	// 		// 		// case reflect.Complex128:
	// 		// 		// case reflect.Array:
	// 		// 		// case reflect.Chan:
	// 		// 		// case reflect.Interface:
	// 		// 		// case reflect.Pointer:
	// 		// 		// case reflect.String:
	// 		// 		// case reflect.Struct:
	// 		// 		// case reflect.UnsafePointer:
	// 		// 	}
	// 	}

	// 	if isBuiltIn(spec.dynamic) {
	// 		return Convertible{spec}
	// 	}

	// func is[T any](t reflect.Type) bool {
	// 	return t == reflect.TypeFor[T]()
	// }

	// func isBuiltIn(t reflect.Type) bool {
	// 	if t == nil {
	// 		return false
	// 	}
	// 	return t.PkgPath() == "" && t.Name() != ""
	// }

	// }
	// return equalitySpec{spec}
	panic("not implemented")
}
