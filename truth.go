package akin

// Truth represents a three-valued logic type.
//
// A [Truth] has three possible values; [True], [False] and [Indeterminate],
// denoted 𝓽, 𝓯 and 𝓾, respectively (mathematical bold script small letters).
type Truth struct{ truth int }

var (
	// Indeterminate (or 𝓾) is a [Truth] that is neither [True] nor [False].
	Indeterminate = Truth{}

	// True (or 𝓽) is the [Truth] that is equivalent to the boolean true.
	True = Truth{1}

	// False (or 𝓯) is the [Truth] that is equivalent to the boolean false.
	False = Truth{-1}
)

// truth returns the [Truth] that is equivalent to the boolean value v.
func truth[T ~bool](v T) Truth {
	if v {
		return True
	}
	return False
}

func (t Truth) String() string {
	switch t {
	case True:
		return "𝓽"
	case False:
		return "𝓯"
	default:
		return "𝓾"
	}
}
