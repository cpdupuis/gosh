package lang

type Scope struct {
	ScopeTable map[string]Value
	Parent *Scope
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
func (scope *Scope) Resolve(sym *Symbol) Value {
	key := sym.Sym
	for sc := scope; sc != nil; sc = sc.Parent {
		if val := sc.ScopeTable[key]; val != nil {
			return val
		}
	}
	return Nil
}

func (scope *Scope) Define(sym *Symbol, val Value) {
	key := sym.Sym
	scope.ScopeTable[key] = val
}
