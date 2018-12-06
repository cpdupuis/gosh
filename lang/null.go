package lang

type null struct {
}

func (nul *null) First() Value {
	panic("Don't do that")
}

func (nul *null) Rest() List {
	return Nil
}

func (nul *null) Type() Type {
	return NullType
}
func (nul *null) String() string {
	return "Nil"
}
func (nul *null) Eval(sc *Scope, ec *EvalContext) (Value,error) {
	// evaluating null returns itself
	return Nil,nil
}

func (nul *null) Length() int {
	return 0
}

var Nil *null
