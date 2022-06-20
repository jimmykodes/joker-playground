//go:build wasm

package main

import (
	"bytes"
	"strings"
	"syscall/js"

	"github.com/jimmykodes/jk/evaluator"
	"github.com/jimmykodes/jk/lexer"
	"github.com/jimmykodes/jk/object"
	"github.com/jimmykodes/jk/parser"
)

var (
	errors []string
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("run", js.FuncOf(func(this js.Value, args []js.Value) any {
		l := lexer.New(args[0].String())
		p := parser.New(l)
		prog := p.ParseProgram()
		out := bytes.NewBuffer(nil)
		obj := evaluator.Eval(prog, object.NewEnvironment(object.WithOut(out)))
		errors = make([]string, len(p.Errors()))
		for i, err := range p.Errors() {
			errors[i] = err.Error()
		}
		if obj == nil {
			return "<nil>"
		}
		return out.String() + obj.Inspect()
	}))
	js.Global().Set("errors", js.FuncOf(func(_ js.Value, _ []js.Value) any {
		return strings.Join(errors, "@@")
	}))
	<-c
}
