package lang

import (
	"errors"
	"fmt"
)

func CreateBuiltin(paramSyms []*Symbol, builtinFunc func(*Scope,[]*Symbol)(Value,error), form Form) *Lambda{
	return &Lambda{ParamSyms: paramSyms, Body: Nil, BuiltinFunc: builtinFunc, Form: form}
}

func DefineQuote(scope *Scope) {
	paramSyms := []*Symbol{Gensym()}
	lambda := &Lambda{ParamSyms: paramSyms, Body: paramSyms[0], Form:DefForm}
	scope.Define(&Symbol{Sym:"quote"}, lambda)
}

var argsyms []*Symbol = []*Symbol{Gensym(), Gensym()}

func DefBinOp(scope *Scope, opname string, op func(Value,Value) (Value,error)) {
	// Define the lambda
	builtInFunc := func(sc *Scope, paramSyms []*Symbol) (Value,error) {
		arg0 := sc.Resolve(paramSyms[0])
		arg1 := sc.Resolve(paramSyms[1])
		return op(arg0, arg1)
	}
	lambda := &Lambda{ParamSyms:argsyms, Body: Nil, BuiltinFunc: builtInFunc, Form: StandardForm}
	// Register its definition with the scope
	scope.Define(&Symbol{Sym:opname}, lambda)
}

func generateIntBinOp(numfunc func(int64,int64)int64) func(Value,Value) (Value,error)  {
	return func(arg0 Value, arg1 Value) (Value,error) {
		num0, ok := arg0.(*Int)
		if !ok {
			return nil,errors.New(fmt.Sprintf("Not a number: %v", arg0))
		}
		num1, ok := arg1.(*Int)
		if !ok {
			return nil,errors.New(fmt.Sprintf("Not a number: %v", arg1))
		}
		result := numfunc(num0.Number,num1.Number)
		return &Int{Number:result},nil
	}
}

func DefineIntBinOps(sc *Scope) {
	DefBinOp(sc, "+", generateIntBinOp(func(a int64, b int64)int64 { return a + b}))
	DefBinOp(sc, "-", generateIntBinOp(func(a int64, b int64)int64 { return a - b}))
	DefBinOp(sc, "*", generateIntBinOp(func(a int64, b int64)int64 { return a * b}))
	DefBinOp(sc, "/", generateIntBinOp(func(a int64, b int64)int64 { return a / b}))
	DefBinOp(sc, "%", generateIntBinOp(func(a int64, b int64)int64 { return a % b}))	
}

func BuiltinCons(scope *Scope, paramSyms []*Symbol) (Value,error) {
	if len(paramSyms) != 2 {
		return Nil,errors.New(fmt.Sprintf("Too many args: %d", len(paramSyms)))
	}
	f := scope.Resolve(paramSyms[0])
	r := scope.Resolve(paramSyms[1])
	if first,ok := f.(Value); ok {
		if rest,ok := r.(List); ok {
			return &Cons{Car:first, Cdr:rest}, nil
		} else {
			return Nil,errors.New(fmt.Sprintf("Invalid rest: %+v", r))
		}
	} else {
		return Nil,errors.New(fmt.Sprintf("Invalid first: %+v", f))
	}
}

func BuiltinDef(scope *Scope, paramSyms []*Symbol) (Value,error) {
	if len(paramSyms) != 2 {
		return Nil,errors.New(fmt.Sprintf("Too many args: %d", len(paramSyms)))
	}
	key := scope.Resolve(paramSyms[0])
	val := scope.Resolve(paramSyms[1])
	if k,ok := key.(*Symbol); ok {
		scope.Parent.Define(k, val)
		return val,nil
	} else {
		return Nil,errors.New(fmt.Sprintf("Key is not symbol: %+v", key))
	}
}

func BuiltinQuote(scope *Scope, paramSyms []*Symbol) (Value,error) {
	if len(paramSyms) != 1 {
		return nil,errors.New(fmt.Sprintf("Wrong number of args: %d", len(paramSyms)))
	}
	val := scope.Resolve(paramSyms[0])
	return val,nil
}

func BuiltinEval(scope *Scope, paramSyms []*Symbol) (Value,error) {
	if len(paramSyms) != 1 {
		return nil,errors.New(fmt.Sprintf("Wrong number of args: %d", len(paramSyms)))
	}
	val := scope.Resolve(paramSyms[0])
	ec := &EvalContext{}
	
	once,err := val.Eval(scope, ec)
	if err != nil {
		return nil,err
	}
	return once,nil
}

func BuiltinLambda(scope *Scope, paramSyms []*Symbol) (Value, error) {
	argslist := scope.Resolve(paramSyms[0])
	if _,ok := argslist.(*Cons); ok {
		body := scope.Resolve(paramSyms[1])
		lambda := &Lambda{ParamSyms:paramSyms, Body:body, Form:LambdaForm}
		return lambda,nil		
	} else {
		return nil,errors.New(fmt.Sprint("Not a symbol slice: ",argslist))
	}
}