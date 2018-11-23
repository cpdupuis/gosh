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

// Return the next complete parseable chunk.
func read(interp *Interpreter) (string, CommandStatus) {
	fmt.Printf("$ ")
	val, _ := interp.Reader.ReadString('\n')
	chunk := strings.TrimRight(val, "\r\n")
	return chunk,Complete
}

func eval(interp *Interpreter, cmd string) (Value, CommandStatus) {
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
		cmd, stat := read(interp)
		if stat == Quit {
			return
		}
		v, stat := eval(interp, cmd)
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
