package lang_test

import (
	"github.com/cpdupuis/gosh/lang"
	"testing"
)

func TestParseTree(t *testing.T) {
	inCh := make(chan string,256)
	go func() {
		inCh <- "("
		inCh <- "testing"
		inCh <- "123"
		inCh <- "("
		inCh <- "foo"
		inCh <- "bar"
		inCh <- ")"
		inCh <- "\"last one\""
		inCh <- ")"
		inCh <- "Nope" // Expect this to be ignored since it's not in the current s-expression
	}()
	val, status := lang.ParseSExp(inCh)
	if status != lang.OK {
		t.Fail()
	}
	valStr := val.String()
	if valStr != "( testing 123 ( foo bar ) \"last one\" )" {
		t.Errorf("Wrong string: %s", valStr)
	}

}