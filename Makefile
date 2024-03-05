all: js/main.wasm js/wasm_exec.js

js/main.wasm:
	GOOS=js GOARCH=wasm go build -ldflags "-X github.com/jimmykodes/joker/builtins.PrintDest=buffer" -o js/main.wasm

js/wasm_exec.js:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js js/wasm_exec.js
