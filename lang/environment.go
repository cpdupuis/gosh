package lang

type Environment struct {
	Root *Scope
}

func NewEnvironment() *Environment {
	env := &Environment{}
	sc := NewScope(nil)
	SetupRootScope(sc)
	DefineIntBinOps(sc)
	DefineQuote(sc)
	env.Root = sc
	return env
}

func SetupRootScope(sc *Scope) {
	ab := []string {"a", "b"}
	cons := CreateBuiltin(ab, BuiltinCons, StandardForm)
	sc.Define(&Symbol{Sym:"cons"}, cons)
	def := CreateBuiltin(ab, BuiltinDef, DefForm)
	sc.Define(&Symbol{Sym:"def"}, def)
}
