package akin

import "reflect"

func compileMap(spec reflect.Value) Spec {
	return specFunc(func(arg reflect.Value) error {
		if err := requireKind(reflect.Map, arg); err != nil {
			return err
		}
		// MAP
		return nil
	})
}
