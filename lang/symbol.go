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

func (sym *Symbol) Eval(sc *Scope, ec *EvalContext) (Value, error) {
	// Evaluating a symbol returns its value in the scope
	val := sc.Resolve(sym)
	if val == Nil {
		return sym,nil
	} else {
		return val,nil
	}
}
