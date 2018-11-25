package lang

import (
	"fmt"
)

type String struct {
	Str string
}
func (str *String) Type() Type {
	return StringType
}
func (str *String) String() string {
	return fmt.Sprintf("\"%s\"", str.Str)
}
func (str *String) Eval(sc *Scope) (Value, error) {
	// Evaluating a String returns itself
	return str, nil
}
