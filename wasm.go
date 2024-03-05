package main

import (
	"io"
	"strings"
	"syscall/js"

	"github.com/jimmykodes/joker/builtins"
	"github.com/jimmykodes/joker/compiler"
	"github.com/jimmykodes/joker/lexer"
	"github.com/jimmykodes/joker/parser"
	"github.com/jimmykodes/joker/vm"
)

var errors []string

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("run", js.FuncOf(func(this js.Value, args []js.Value) any {
		errors = []string{}
		l := lexer.New(args[0].String())
		p := parser.New(l)
		prog := p.ParseProgram()
		if len(p.Errors()) > 0 {

			errors = make([]string, len(p.Errors()))
			for i, err := range p.Errors() {
				errors[i] = err.Error()
			}
			return "err"
		}
		comp := compiler.New()
		err := comp.Compile(prog)
		if err != nil {
			errors = []string{err.Error()}
			return "err"
		}
		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			errors = []string{err.Error()}
			return "err"
		}

		obj := machine.LastPoppedStackElem()

		if obj == nil {
			return "<nil>"
		}
		data, _ := io.ReadAll(builtins.Dest)
		return string(data) + obj.Inspect()
	}))

	js.Global().Set("errors", js.FuncOf(func(_ js.Value, _ []js.Value) any {
		return strings.Join(errors, "@@")
	}))
	<-c
}
