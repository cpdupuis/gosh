package interp

import (
	"os"
	"fmt"
	"github.com/cpdupuis/gosh/lang"
)

func repl(treeCh chan lang.Value) {
	env := lang.NewEnvironment()
	for {
		tree := <-treeCh
		ctx := &lang.EvalContext{}
		ctx.Push(lang.StandardForm)
		fmt.Printf("Tree: <%+v>\n", tree)
		res,err := tree.Eval(env.Root, ctx)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
		} else {
			fmt.Printf("Result: %+v\n", res)
		}
	}
}


func treeize(inCh chan string, outCh chan lang.Value) lang.Value {
	for {
		sexp,_ := lang.ParseSExp(inCh)
		outCh <- sexp
	}
}

func Interp() {
	tokenCh := make(chan string, 256)
	treeCh := make(chan lang.Value, 256)
	go lang.Tokenize(os.Stdin, tokenCh)
	go treeize(tokenCh, treeCh)
	repl(treeCh)
}
