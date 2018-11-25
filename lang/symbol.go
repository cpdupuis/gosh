package lang

type Symbol struct {
	Sym string
}

func (sym *Symbol) Type() Type {
	return SymbolType
}

func (sym *Symbol) String() string {
	return sym.Sym
}

func (sym *Symbol) Eval(sc *Scope) (Value, error) {
	// Evaluating a symbol returns its value in the scope
	return sc.Resolve(sym), nil
}
