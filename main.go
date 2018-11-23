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

func read(interp *Interpreter) string {
	val, _ := interp.Reader.ReadString('\n')
	return strings.TrimRight(val, "\r\n")
}

func eval(interp *Interpreter, cmd string) Value {
	// For now, let's just return "woo hoo" if the value is "hello world", and otherwise "yup"
	if cmd == "hello world" {
		return "woo hoo"
	} else {
		return fmt.Sprintf("yup: <%s>", cmd)
	}
}

func print(interp *Interpreter, value Value) {
	fmt.Printf("%+v\n", value)
}

func repl(interp *Interpreter) {
	for {
		fmt.Printf("$ ")
		s := read(interp)
		v := eval(interp, s)
		print(interp, v)
	}
}


func main() {
	interp := &Interpreter {
		Reader: bufio.NewReader(os.Stdin),
	}
	repl(interp)
}
