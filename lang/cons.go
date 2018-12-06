package lang

import (
	"strings"
	"errors"
	"fmt"
)

type Cons struct {
	Car Value
	Cdr List
}

func (cons *Cons) First() Value {
	return cons.Car
}

func (cons *Cons) Rest() List {
	return cons.Cdr
}

func (cons *Cons) Type() Type {
	return ConsType
}

func (cons *Cons) String() string {
	strarray := make([]string, 0)
	strarray = append(strarray, "(")
	strarray = append(strarray, cons.Car.String())
	curr := cons.Cdr

Loop:	
	for {
		switch item := curr.(type) {
		case *null:
			break Loop
		case *Cons:
			strarray = append(strarray, item.Car.String())
			curr = item.Cdr
		default:
			panic("Unexpected!")
		}
	}
	strarray = append(strarray, ")")
	return strings.Join(strarray, " ")
}

func (cons *Cons) Eval(sc *Scope, ec *EvalContext) (Value,error) {
	// Eval'ing a cons means calling the lambda in the car with the cons in the cdr.
	first := cons.Car
	firstVal,err := first.Eval(sc, ec)
	if err != nil {
		return Nil,err
	}
	if lambda,ok := firstVal.(*Lambda); ok {
		ec.Push(lambda.Form)
		retval,err := lambda.Call(sc, ec, cons.Cdr)
		ec.Pop()
		return retval,err
	} else {
		return Nil,errors.New(fmt.Sprintf("Can't eval this: %+v", cons))
	}
}

func (cons *Cons) Length() int {
	return 1 + cons.Cdr.Length()
}

