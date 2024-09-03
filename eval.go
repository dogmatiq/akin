package akin

// Eval evaluates 𝑷❨𝒙❩.
func Eval(p Predicate, x any) (Truth, Rationale) {
	e := &evaluator{
		P: p,
		X: x,
	}

	p.VisitP(e)

	return e.PX, PofX(*e)
}

type evaluator PofX
