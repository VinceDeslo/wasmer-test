# WASM - TEST

This repo is a playground to mess around with compiling GO to WASM.

### Objective
It will eventually serve me as a template for spinning up quick WASM projects in GO. Be it for the browser or on the server.

### Top level Make commands:
- `make run-node` : compiles the main.go file to a WASM binary and executes it in node.
- `make run-tests` : compiles the main.go file to a WASM binary and executes go tests in node.
- `make run-browser` : compiles the main.go file to a WASM binary and serves it on localhost:9090

### Structure
```
wasm_test
├── Makefile
├── README.md
├── go.mod
└── webassembly
    ├── assets
    │   ├── ** Assets for serving WASM in browser**
    ├── cmd
    │   ├── server
    │   │   └── ** Web browser server **
    │   └── wasm
    │       └── ** Application code compiled to WASM **
    └── scripts
        ├── ** Various scripts called by the Makefile **
```
