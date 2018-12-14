package interp

import (
	"os"
	"fmt"
	"github.com/cpdupuis/gosh/lang"
)


func ReplOne(tree lang.Value, env *lang.Environment) (lang.Value,error) {
	ctx := &lang.EvalContext{}
	ctx.Push(lang.StandardForm)
	fmt.Printf("Tree: <%+v>\n", tree)
	res,err := tree.Eval(env.Root, ctx)
	ctx.Pop()
	return res,err
}

func repl(treeCh chan lang.Value) {
	env := lang.NewEnvironment()
	for {
		tree := <-treeCh
		res,err := ReplOne(tree, env)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
		} else {
			fmt.Printf("Result: %+v\n", res)
		}
		}
}


func treeize(inCh chan string, outCh chan lang.Value) lang.Value {
	for {
		sexp,_ := ParseSExp(inCh)
		outCh <- sexp
	}
}

func Interp() {
	tokenCh := make(chan string, 256)
	treeCh := make(chan lang.Value, 256)
	go Tokenize(os.Stdin, tokenCh)
	go treeize(tokenCh, treeCh)
	repl(treeCh)
}
