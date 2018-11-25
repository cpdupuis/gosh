package lang

type Environment struct {
	Root *Scope
}

func NewEnvironment() *Environment {
	env := &Environment{}
	sc := NewScope(nil)
	SetupRootScope(sc)
	env.Root = sc
	return env
}

func SetupRootScope(sc *Scope) {
	ab := []string {"a", "b"}
	plus := CreateBuiltin(ab, BuiltinPlus)
	sc.Define(&Symbol{Sym:"+"}, plus)
}

// Here are all the functions in the default environment
func Def(sc *Scope, sym *Symbol, val Value) {
	sc.Define(sym, val)
}
