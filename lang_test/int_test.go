package lang_test

import (
	"github.com/cpdupuis/gosh/lang"
	"testing"
)

func TestIntString(t *testing.T) {
	n := lang.Int{Number: -1234567890123456789}
	s := n.String()
	if s != "-1234567890123456789" {
		t.Fail()
	}
}