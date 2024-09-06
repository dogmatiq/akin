package akin

// Eval evaluates 𝑷❨𝒙❩, returning the [Result] and a [Rationale] for that result.
func Eval(p Assertion, x any) (Result, Rationale) {
	return eval(
		p, x,
		defaultPredicateExpr,
		defaultValueExpr,
	)
}

func eval(p Assertion, x any, px, vx Expr) (Result, Rationale) {
	e := evaluator{
		PredicateExpr: px,
		Value:         valueOf(x),
		ValueExpr:     vx,
	}

	p.acceptAssertionVisitor(&e)

	if e.Rationale == nil {
		panic("no rationale provided")
	}

	return e.Result, AssertionRationale{
		Assertion:     p,
		AssertionExpr: px,
		Value:         e.Value,
		ValueExpr:     vx,
		Result:        e.Result,
		Rationale:     e.Rationale,
	}
}

type evaluator struct {
	PredicateExpr Expr
	Value         Value
	ValueExpr     Expr
	Result        Result
	Rationale     Rationale
}

// Result is a three-valued logic type.
//
// It has three possible values; [True], [False] and [Undefined], denoted using
// the mathematical bold script small letters 𝓽, 𝓯 and 𝓾, respectively.
type Result struct{ v int }

var (
	// Undefined (or 𝓾) is a [Result] that is neither [True] nor [False].
	Undefined = Result{}

	// True (or 𝓽) is the [Result] that is equivalent to the boolean true.
	True = Result{1}

	// False (or 𝓯) is the [Result] that is equivalent to the boolean false.
	False = Result{-1}
)

func asResult[T ~bool](v T) Result {
	if v {
		return True
	}
	return False
}

func (t Result) String() string {
	switch t {
	case True:
		return "𝓽"
	case False:
		return "𝓯"
	default:
		return "𝓾"
	}
}
