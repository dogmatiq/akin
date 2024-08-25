package akin

func compileBool(spec Value) Spec {
	if !isBuiltIn(spec.rtype) {
		return equalitySpec{spec}
	}
	return losslessConversionSpec{spec}
}
