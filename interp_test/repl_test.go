package interp_test

import (
	"github.com/cpdupuis/gosh/interp"
	"github.com/cpdupuis/gosh/lang"
	"testing"
)

func TestInterpAddition(t *testing.T) {
	cmd := &lang.Cons{First:&lang.Int{Number:11}, Rest:lang.Nil}
	cmd = &lang.Cons{First:&lang.Int{Number: 88}, Rest:cmd}
	cmd = &lang.Cons{First:&lang.Symbol{Sym:"+"}, Rest: cmd}

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
	cmd1 := &lang.Cons{First:&lang.Int{Number:8675309}, Rest:lang.Nil}
	cmd1 = &lang.Cons{First:&lang.Symbol{Sym:"a"}, Rest:cmd1}
	cmd1 = &lang.Cons{First:&lang.Symbol{Sym:"def"}, Rest:cmd1}
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
	cmd2 := &lang.Cons{First:&lang.Symbol{Sym:"a"}, Rest:lang.Nil}
	cmd2 = &lang.Cons{First:&lang.Symbol{Sym:"quote"}, Rest: cmd2}
	cmd2 = &lang.Cons{First:cmd2, Rest:lang.Nil}
	cmd2 = &lang.Cons{First:&lang.Symbol{Sym:"b"}, Rest: cmd2}
	cmd2 = &lang.Cons{First: &lang.Symbol{Sym:"def"}, Rest: cmd2}
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
	cmd3 := &lang.Cons{First:&lang.Symbol{Sym:"b"}, Rest:lang.Nil}
	cmd3 = &lang.Cons{First:&lang.Symbol{Sym:"eval"}, Rest:cmd3}
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