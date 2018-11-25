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
	cons := &lang.Cons{First:sym, Rest:lang.Nil}
	length:= cons.Length()
	if length != 1 {
		t.Fail()
	}
}

func TestLength4(t *testing.T) {
	sym := &lang.Symbol{Sym:"hi"}
	cons4 := &lang.Cons{First:sym, Rest:lang.Nil}
	cons3 := &lang.Cons{First:sym, Rest:cons4}
	cons2 := &lang.Cons{First:sym, Rest:cons3}
	cons1 := &lang.Cons{First:sym, Rest:cons2}
	length:= cons1.Length()
	if length != 4 {
		t.Fail()
	}
}

func TestLengthDot(t *testing.T) {
	sym := &lang.Symbol{Sym:"hi"}
	cons := &lang.Cons{First:sym, Rest:sym}
	length := cons.Length()
	if length != 2 {
		t.Fail()
	}
}
