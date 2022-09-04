#!/bin/bash
cd ../cmd/wasm
echo "Compiling TinyGo WASI module..."
tinygo build -wasm-abi=generic -target=wasi -o ../../out/tiny.wasm ./main.go
echo "Compilation finished."