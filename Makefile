all: js/main.wasm js/wasm_exec.js

js/main.wasm:
	GOOS=js GOARCH=wasm go build -o js/main.wasm

js/wasm_exec.js:
	cp ${HOME}/go/misc/wasm/wasm_exec.js js/wasm_exec.js
