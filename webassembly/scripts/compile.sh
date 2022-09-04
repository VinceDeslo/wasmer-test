#!/bin/bash
cd ../cmd/wasm
echo "Compiling TinyGo Printing WASI module..."
tinygo build -wasm-abi=generic -target=wasi -o ../../out/printing.wasm ./printing.go
echo "Compilation finished."