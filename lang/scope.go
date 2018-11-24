package lang

type Scope struct {
	ScopeTable map[string]Value
}

func (scope *Scope) Type() Type {
	return ScopeType
}

func (scope *Scope) String() string {
	return "Scope"
}

