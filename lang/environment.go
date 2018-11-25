package lang

type Environment struct {
	Root *Scope
}

func NewEnvironment() *Environment {
	env := &Environment{}
	sc := &Scope{}
	env.Root = sc
	return env
}

// Here are all the functions in the default environment
func Def(sc *Scope, sym *Symbol, val Value) {
	sc.Define(sym, val)
}
