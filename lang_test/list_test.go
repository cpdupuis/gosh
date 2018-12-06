package lang_test

import (
	"testing"
	"github.com/cpdupuis/gosh/lang"
)

func TestLengthNull(t *testing.T) {
	length := lang.Nil.Length()
	if length != 0 {
		t.Fail()
	}
}

func TestLength1(t *testing.T) {
	sym := &lang.Symbol{Sym:"hi"}
	cons := &lang.Cons{Car:sym, Cdr:lang.Nil}
	length:= cons.Length()
	if length != 1 {
		t.Fail()
	}
}

func TestLength4(t *testing.T) {
	sym := &lang.Symbol{Sym:"hi"}
	cons4 := &lang.Cons{Car:sym, Cdr:lang.Nil}
	cons3 := &lang.Cons{Car:sym, Cdr:cons4}
	cons2 := &lang.Cons{Car:sym, Cdr:cons3}
	cons1 := &lang.Cons{Car:sym, Cdr:cons2}
	length:= cons1.Length()
	if length != 4 {
		t.Fail()
	}
}
