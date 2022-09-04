package main

import (
	"io/ioutil"

	"github.com/VinceDeslo/wasmer_test/webassembly/internal/errors"
	"github.com/wasmerio/wasmer-go/wasmer"
)

func LoadWasm(filePath string, functionNames []string) *WasmFunctions {
	var wasmFns = NewWasmFunctions()

	wasmBytes, err := ioutil.ReadFile(filePath)

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	module, err := wasmer.NewModule(store, wasmBytes)
	errors.Check(err, "Failed to compile the WASM module.")

	wasiEnv, err := wasmer.NewWasiStateBuilder("wasi-program").Finalize()
	errors.Check(err, "Failed to create the WASI state builder.")

	importObject, err := wasiEnv.GenerateImportObject(store, module)
	errors.Check(err, "Failed to generate the WASI import object.")

	instance, err := wasmer.NewInstance(module, importObject)
	errors.Check(err, "Failed to instantiate the WASM module.")

	for _, functionName := range functionNames {
		fn, err := instance.Exports.GetFunction(functionName)
		errors.Check(err, "Failed to fetch WASM function.")
		wasmFns.functions[functionName] = fn
	}

	return wasmFns
}
