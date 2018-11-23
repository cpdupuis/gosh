package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

type Symbol string

type Value interface{}


type Interpreter struct {
	Globals map[Symbol]Value
	Reader *bufio.Reader
}

type CommandStatus int
const (
	Complete = iota
	Incomplete
	Quit
)

func readEval(interp *Interpreter) (Value, CommandStatus) {
	val, _ := interp.Reader.ReadString('\n')
	cmd := strings.TrimRight(val, "\r\n")
	// For now, let's just return "woo hoo" if the value is "hello world", and otherwise "yup"
	if cmd == "hello world" {
		return "woo hoo", Complete
	} else if cmd == "exit" {
		return "", Quit
	} else {
		return fmt.Sprintf("yup: <%s>", cmd), Complete
	}
}

func print(interp *Interpreter, value Value) {
	fmt.Printf("%+v\n", value)
}

func repl(interp *Interpreter) {
	for {
		fmt.Printf("$ ")
		v, stat := readEval(interp)
		if stat == Quit {
			return
		}
		print(interp, v)
	}
}


func main() {
	interp := &Interpreter {
		Reader: bufio.NewReader(os.Stdin),
	}
	repl(interp)
}
