package akin

import (
	"reflect"
)

func compileBool(spec reflect.Value) Spec {
	specT := spec.Type()

	if !isBuiltIn(specT) {
		return specFunc(func(arg reflect.Value) error {
			if err := requireType(specT, arg); err != nil {
				return err
			}
			return requireEquality(arg, spec)
		})
	}

	return specFunc(func(arg reflect.Value) error {
		if err := requireKind(reflect.Bool, arg); err != nil {
			return err
		}

		if spec.Bool() != arg.Bool() {
			return notEqual(arg, spec)
		}

		return nil
	})
}
