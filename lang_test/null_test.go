package lang_test

import (
	"testing"
	"github.com/cpdupuis/gosh/lang"
)

func TestNullType(t *testing.T) {
	if lang.Nil.Type() != lang.NullType {
		t.Errorf("Expected NullType, got: %v", lang.Nil.Type())
	}
}

func TestNullStr(t *testing.T) {
	if lang.Nil.String() != "Nil" {
		t.Errorf("Expected Nil, got %v", lang.Nil.String())
	}
}

