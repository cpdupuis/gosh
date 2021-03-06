package lang

import (
	"errors"
	"fmt"
)
type Scope struct {
	ScopeTable map[string]Value
	Parent *Scope
}

func NewScope(parent *Scope) *Scope {
	scope := &Scope{Parent:parent}
	scope.ScopeTable = make(map[string]Value)
	return scope
}

func (scope *Scope) Type() Type {
	return ScopeType
}

func (scope *Scope) String() string {
	return fmt.Sprintf("ScopeTable: %+v", scope.ScopeTable)
}
func (scope *Scope) Eval(sc *Scope, ec *EvalContext) (Value,error) {
	// You can't evaluate a scope. Sorry.
	return Nil,errors.New("Can't eval a scope")
}

func (scope *Scope) Resolve(sym *Symbol) Value {
	key := sym.Sym
	for sc := scope; sc != nil; sc = sc.Parent {
		if val := sc.ScopeTable[key]; val != nil {
			return val
		}
	}
	return nil
}

func (scope *Scope) Define(sym *Symbol, val Value) {
	key := sym.Sym
	scope.ScopeTable[key] = val
}
