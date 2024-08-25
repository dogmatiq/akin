package akin

import "reflect"

func compileSlice(Value) Spec {
	return specFunc(func(arg Value) error {
		if err := requireKind(reflect.Slice, arg); err != nil {
			return err
		}
		// TODO
		return nil
	})
}
