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
//	cmd2 := &lang.Cons{First:&lang.Symbol{Sym:"a"}, Rest:lang.Nil}

}