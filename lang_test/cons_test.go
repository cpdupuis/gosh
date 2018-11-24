package lang_test

import (
	"github.com/cpdupuis/gosh/lang"
	"testing"
)

func TestConsString(t *testing.T) {
	hello := &lang.String{Str: "Hello"}
	world := &lang.String{Str: "world"}
	num := &lang.Int{Number: 123}
	a := &lang.Cons{First: world, Rest: lang.Nil}
	b := &lang.Cons{First: hello, Rest: a}
	if b.String() != "( \"Hello\" \"world\" )" {
		t.Errorf("Wrong first string: %v\n", b.String())
	}
	d := &lang.Cons{First: hello, Rest: num}
	c := &lang.Cons{First: b, Rest: d}
	if c.String() != "( ( \"Hello\" \"world\" ) \"Hello\" . 123 )" {
		t.Errorf("Wrong second string: %v\n", c.String())
	}

}