package lang

type null struct {
}
func (nul *null) Type() Type {
	return NullType
}
func (nul *null) String() string {
	return "Nil"
}
func (nul *null) Eval(sc *Scope) (Value,error) {
	// evaluating null returns itself
	return Nil,nil
}

func (nul *null) Length() int {
	return 0
}

var Nil *null
