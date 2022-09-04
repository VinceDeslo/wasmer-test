package main

import (
	"io/ioutil"
	"log"

	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
	filePath := "../../out/tiny.wasm"
	wasmBytes, err := ioutil.ReadFile(filePath)

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	module, err := wasmer.NewModule(store, wasmBytes)
	check(err, "Failed to compile the WASM module.")

	wasiEnv, err := wasmer.NewWasiStateBuilder("wasi-program").Finalize()
	check(err, "Failed to create the WASI state builder.")

	importObject, err := wasiEnv.GenerateImportObject(store, module)
	check(err, "Failed to generate the WASI import object.")

	instance, err := wasmer.NewInstance(module, importObject)
	check(err, "Failed to instantiate the WASM module.")

	msg, err := instance.Exports.GetFunction("PrintMessage")
	check(err, "Failed to fetch WASM functions.")

	msg()
}

func check(e error, msg string) {
	if e != nil {
		log.Println(e)
		log.Panic(msg)
	}
}
