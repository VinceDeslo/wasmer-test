#!/bin/bash
EXEC_PATH="$(go env GOROOT)/misc/wasm/wasm_exec.js"
echo "Copying glue code from $EXEC_PATH"
cp $EXEC_PATH ../assets
