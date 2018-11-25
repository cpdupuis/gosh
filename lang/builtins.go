package lang

import (
	"errors"
	"fmt"
)

func CreateBuiltin(paramNames []string, builtinFunc func(*Scope,[]*Symbol)(Value,error)) *Lambda{
	syms := make([]*Symbol, len(paramNames))
	for i,paramName := range(paramNames) {
		syms[i] = &Symbol{Sym:paramName}
	}
	return &Lambda{ParamSyms: syms, Body: Nil, BuiltinFunc: builtinFunc}
}

func BuiltinPlus(scope *Scope, paramSyms []*Symbol) (Value,error) {
	var res int64
	for _, paramSym := range(paramSyms) {
		val := scope.Resolve(paramSym)
		num,ok := val.(*Int)
		if ok {
			res += num.Number
		} else {
			return Nil, errors.New(fmt.Sprintf("Not a number: %v", val))
		}
	}
	retval := &Int{Number:res}
	return retval,nil
}
