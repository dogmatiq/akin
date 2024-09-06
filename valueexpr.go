package akin

import "fmt"

// ValueExpr is an expression that describes how a specific [Value] was
// obtained.
type ValueExpr interface {
	visit(v ValueExprVisitor)
}

// ValueExprVisitor is an algorithm with logic specific to each [ValueExpr]
// type.
type ValueExprVisitor interface {
	VarExpr(VarExpr)
	LitExpr(LitExpr)
	IndexExpr(IndexExpr)
	KeyExpr(KeyExpr)
	FieldExpr(FieldExpr)
	DerefExpr(DerefExpr)
}

type (
	// VarExpr is a [ValueExpr] describing a named variable, such as 𝒙.
	VarExpr struct {
		Name string
	}

	// LitExpr is a [ValueExpr] describing a literal value.
	LitExpr struct {
		Value Value
	}

	// IndexExpr is a [ValueExpr] describing the an element within a slice or
	// array.
	IndexExpr struct {
		Indexable ValueExpr
		Index     int
	}

	// KeyExpr is a [ValueExpr] describing the [Value] associated with a
	// specific key within a map.
	KeyExpr struct {
		Map ValueExpr
		Key Value
	}

	// FieldExpr is a [ValueExpr] describing a specific field within a struct.
	FieldExpr struct {
		Receiver ValueExpr
		Field    string
	}

	// DerefExpr is a [ValueExpr] describing the dereferenced value of a
	// pointer.
	DerefExpr struct {
		Pointer ValueExpr
	}
)

func (e VarExpr) visit(v ValueExprVisitor)   { v.VarExpr(e) }
func (e LitExpr) visit(v ValueExprVisitor)   { v.LitExpr(e) }
func (e IndexExpr) visit(v ValueExprVisitor) { v.IndexExpr(e) }
func (e KeyExpr) visit(v ValueExprVisitor)   { v.KeyExpr(e) }
func (e FieldExpr) visit(v ValueExprVisitor) { v.FieldExpr(e) }
func (e DerefExpr) visit(v ValueExprVisitor) { v.DerefExpr(e) }

func (e VarExpr) String() string {
	return e.Name
}

func (e LitExpr) String() string {
	return e.Value.String()
}

func (e IndexExpr) String() string {
	return fmt.Sprintf("%s[%d]", e.Indexable, e.Index)
}

func (e KeyExpr) String() string {
	return fmt.Sprintf("%s[%s]", e.Map, e.Key)
}

func (e FieldExpr) String() string {
	r := e.Receiver

	if d, ok := e.Receiver.(DerefExpr); ok {
		// Render x.Field instead of (*x).Field, since the explicit dereference
		// is not necessary in Go's syntax.
		r = d.Pointer
	}

	return fmt.Sprintf("%s.%s", r, e.Field)
}

func (e DerefExpr) String() string {
	return fmt.Sprintf("(*%s)", e.Pointer)
}

var varX = VarExpr{"𝒙"}
