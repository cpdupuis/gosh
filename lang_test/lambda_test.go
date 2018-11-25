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
	lambda := &lang.Lambda{params[:], foo}
	if lambda.Arity() != 3 {
		t.Fail()
	}
}

func TestLambdaString(t *testing.T) {
	foo := &lang.Symbol{Sym: "foo"}
	bar := &lang.Symbol{Sym: "bar"}
	baz := &lang.Symbol{Sym: "baz"}
	params := [3]*lang.Symbol{foo, bar, baz}
	lambda := &lang.Lambda{params[:], foo}
	str := lambda.String()
	if str != "(lambda (foo bar baz) foo)" {
		t.Errorf("Wrong string: %v", str)
	}
}