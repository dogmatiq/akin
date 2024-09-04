package akin

// Eval evaluates 𝑷❨𝒙❩.
func Eval(p Predicate, x any) (Truth, Rationale) {
	e := &evaluator{
		P: p,
		X: valueOf(x),
	}

	p.visitP(e)

	return e.Px, Px(*e)
}

type evaluator Px
