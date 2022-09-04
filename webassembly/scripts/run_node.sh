#!/bin/bash
EXEC_PATH="$(go env GOROOT)/misc/wasm/go_js_wasm_exec"
cd ../cmd/wasm && GOOS=js GOARCH=wasm go run -exec=$EXEC_PATH .