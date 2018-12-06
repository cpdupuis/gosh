package interp_test

import (
	"github.com/cpdupuis/gosh/interp"
	"github.com/cpdupuis/gosh/lang"
	"testing"
)

func TestInterpAddition(t *testing.T) {
	cmd := &lang.Cons{Car:&lang.Int{Number:11}, Cdr:lang.Nil}
	cmd = &lang.Cons{Car:&lang.Int{Number: 88}, Cdr:cmd}
	cmd = &lang.Cons{Car:&lang.Symbol{Sym:"+"}, Cdr: cmd}

	env := lang.NewEnvironment()
	res,err := interp.ReplOne(cmd, env)
	if err != nil {
		t.Fail()
	} 
	if res,ok := res.(*lang.Int); ok {
		if res.Number != 99 {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}

func TestInterpDefQuoteEval(t *testing.T) {
	cmd1 := &lang.Cons{Car:&lang.Int{Number:8675309}, Cdr:lang.Nil}
	cmd1 = &lang.Cons{Car:&lang.Symbol{Sym:"a"}, Cdr:cmd1}
	cmd1 = &lang.Cons{Car:&lang.Symbol{Sym:"def"}, Cdr:cmd1}
	if cmd1.String() != "( def a 8675309 )" {
		t.Fail()
	}
	env := lang.NewEnvironment()
	res,err := interp.ReplOne(cmd1, env)
	if err != nil {
		t.Fail()
	}
	if num,ok := res.(*lang.Int); ok {
		if num.Number != 8675309 {
			t.Fail()
		}
	} else {
		t.Fail()
	}
	cmd2 := &lang.Cons{Car:&lang.Symbol{Sym:"a"}, Cdr:lang.Nil}
	cmd2 = &lang.Cons{Car:&lang.Symbol{Sym:"quote"}, Cdr: cmd2}
	cmd2 = &lang.Cons{Car:cmd2, Cdr:lang.Nil}
	cmd2 = &lang.Cons{Car:&lang.Symbol{Sym:"b"}, Cdr: cmd2}
	cmd2 = &lang.Cons{Car: &lang.Symbol{Sym:"def"}, Cdr: cmd2}
	if cmd2.String() != "( def b ( quote a ) )" {
		t.Fail()
	}
	res,err = interp.ReplOne(cmd2, env)
	if err != nil {
		t.Fail()
	}
	if resym,ok := res.(*lang.Symbol); !ok {
		t.Fail()
	} else {
		if resym.Sym != "a" {
			t.Fail()
		}
	}
	cmd3 := &lang.Cons{Car:&lang.Symbol{Sym:"b"}, Cdr:lang.Nil}
	cmd3 = &lang.Cons{Car:&lang.Symbol{Sym:"eval"}, Cdr:cmd3}
	val,err := interp.ReplOne(cmd3, env)
	if err != nil {
		t.Fail()
	}
	if anInt,ok := val.(*lang.Int); !ok {
		t.Fail()
	} else {
		if anInt.Number != 8675309 {
			t.Fail()
		}
	}
}