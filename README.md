# Messing around with wasmer

This repo is a playground to mess around with compiling GO to WASM.

### Objective
Compile a set of Go functions to WASM and consume them from a separate Go program.

### Requirements
- Go
- TinyGo

### Top level Make commands:
- `make run-node` : compiles the main.go file to a WASM binary and executes it in node.
- `make run-tests` : compiles the main.go file to a WASM binary and executes go tests in node.
- `make run-wasmer` : compiles the main.go file to a WASM binary and consumes its contents in a Go program using wasmer.

### Structure
```
wasm_test
├── Makefile
├── README.md
├── go.mod
├── go.sum
└── webassembly
    ├── out
    │   ├── ** WASM binaries **
    ├── cmd
    │   ├── wasm
    │   │   └── ** Application code compiled to WASM **
    │   └── wasmer
    │       └── ** WASM binary consumer code **
    └── scripts
        ├── ** Various scripts called by the Makefile **
```
