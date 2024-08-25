package akin

import "reflect"

func compileSlice(spec reflect.Value) Spec {
	return specFunc(func(arg reflect.Value) error {
		if err := requireKind(reflect.Slice, arg); err != nil {
			return err
		}
		// TODO
		return nil
	})
}
