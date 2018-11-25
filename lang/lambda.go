package lang

import (
	"errors"
	"fmt"
	"strings"
)

type Lambda struct {
	ParamSyms []*Symbol
	Body Value
	BuiltinFunc func(*Scope,[]*Symbol) (Value,error)
}

func (lambda *Lambda) Type() Type {
	return LambdaType
}

func (lambda *Lambda) String() string {
	sb := strings.Builder{}
	var next bool
	for _,item := range(lambda.ParamSyms) {
		if next {
			sb.WriteString(" ")
		}
		sb.WriteString(item.String())
		next = true
	}
	return fmt.Sprintf("(lambda (%v) %v)", sb.String(), lambda.Body.String())
}

func (lambda *Lambda) Eval(scope *Scope) Value {
	// Evaluating a lambda just returns the lambda, I guess. Or maybe it should call it with no args? Seems odd.
	return lambda
}

func (lambda *Lambda) Arity() int {
	return len(lambda.ParamSyms)
}

func (lambda *Lambda) Call(scope *Scope, params List) (Value,error) {
		// OK, so someone wants to call this function with some parameters. Cool. Let's make it happen.
	if params.Length() != lambda.Arity() {
		return Nil,errors.New("Inconceivable!")
	}
	sc := NewScope(scope)
	plist := params
	for _, paramSym := range(lambda.ParamSyms) {
		cons,ok := plist.(*Cons)
		if ok {
			sc.Define(paramSym, cons.First)
			plist = cons.Rest
		} else {
			return Nil,errors.New("Malformed list in call!")
		}
	}
	// OK, our scope is now populated with our values! Woo hoo! Let's eval!
	var result Value
	var err error
	if lambda.BuiltinFunc != nil {
		result,err = lambda.BuiltinFunc(sc, lambda.ParamSyms)
	} else {
		result = lambda.Body.Eval(sc)
	}
	return result,err
}
