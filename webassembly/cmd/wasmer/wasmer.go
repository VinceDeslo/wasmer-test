package main

const (
	PrintMessage     = "PrintMessage"
	PrintDescription = "PrintDescription"
)

const (
	PrintingWasmFile = "../../out/printing.wasm"
)

func main() {
	functionNames := []string{PrintMessage, PrintDescription}

	wasm := LoadWasm(PrintingWasmFile, functionNames)

	msg := wasm.functions[PrintMessage]
	desc := wasm.functions[PrintDescription]

	msg()
	desc()
}
