package lang

import (
	"fmt"
	"strings"
)

type Lambda struct {
	ParamSyms []*Symbol
	Body Value
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

func (lambda *Lambda) Call(params *List) Value {
	// OK, so someone wants to call this function with some parameters. Cool. Let's make it happen.
	return Nil
}
