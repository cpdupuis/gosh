package lang

type Scope struct {
	ScopeTable map[string]Value
}

func (scope *Scope) Type() Type {
	return ScopeType
}

func (scope *Scope) String() string {
	return "Scope"
}
func (scope *Scope) Eval(sc *Scope) Value {
	// Evaluating a scope always returns itself
	return scope
}
