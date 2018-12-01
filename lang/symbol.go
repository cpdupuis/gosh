package lang

import (
	"errors"
	"fmt"
)

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
		return nil, errors.New(fmt.Sprintf("Symbol not defined: %s", sym.String()))
	} else {
		return val,nil
	}
}
