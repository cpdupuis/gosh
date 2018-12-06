package lang_test

import (
	"testing"
	"github.com/cpdupuis/gosh/lang"
)

func TestLambdaArity(t *testing.T) {
	foo := &lang.Symbol{Sym: "foo"}
	bar := &lang.Symbol{Sym: "bar"}
	baz := &lang.Symbol{Sym: "baz"}
	params := [3]*lang.Symbol{foo, bar, baz}
	lambda := &lang.Lambda{ParamSyms: params[:], Body: foo}
	if lambda.Arity() != 3 {
		t.Fail()
	}
}

func TestLambdaString(t *testing.T) {
	foo := &lang.Symbol{Sym: "foo"}
	bar := &lang.Symbol{Sym: "bar"}
	baz := &lang.Symbol{Sym: "baz"}
	params := [3]*lang.Symbol{foo, bar, baz}
	lambda := &lang.Lambda{ParamSyms: params[:], Body: foo}
	str := lambda.String()
	if str != "(lambda (foo bar baz) foo)" {
		t.Errorf("Wrong string: %v", str)
	}
}

func TestLambdaInvalidParams(t *testing.T) {
	foo := &lang.Symbol{Sym: "foo"}
	bar := &lang.Symbol{Sym: "bar"}
	baz := &lang.Symbol{Sym: "baz"}
	params := [3]*lang.Symbol{foo, bar, baz}
	lambda := &lang.Lambda{ParamSyms: params[:], Body: foo}
	scope := lang.NewScope(nil)
	ec := &lang.EvalContext{}
	ec.Push(lang.StandardForm)
	res,err := lambda.Call(scope, ec, lang.Nil)
	if res != lang.Nil {
		t.Fail()
	}
	if err.Error() != "Calling a function with the wrong number of args: 0, expected: 3" {
		t.Errorf("Wrong error: %v", err.Error())
	}
}

func TestLambdaSimpleFunc(t *testing.T) {
	foo := &lang.Symbol{Sym: "foo"}
	bar := &lang.Symbol{Sym: "bar"}
	baz := &lang.Symbol{Sym: "baz"}
	params := [3]*lang.Symbol{foo, bar, baz}
	lambda := &lang.Lambda{ParamSyms: params[:], Body: bar}
	scope := lang.NewScope(nil)
	ec := &lang.EvalContext{}
	ec.Push(lang.StandardForm)
	num1 := &lang.Int{Number: 42}
	num2 := &lang.Int{Number:8675309}
	num3 := &lang.Int{Number:909}
	cons := &lang.Cons{Car: num3, Cdr: lang.Nil}
	cons = &lang.Cons{Car:num2, Cdr: cons}
	cons = &lang.Cons{Car:num1, Cdr: cons}
	res,err := lambda.Call(scope, ec, cons)
	if err != nil {
		t.Fail()
	}
	if res != num2 {
		t.Errorf("Wrong answer: %+v", res)
	}
}

