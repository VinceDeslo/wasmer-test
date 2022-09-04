package main

type WasmFunctions struct {
	functions map[string]func(...interface{}) (interface{}, error)
}

func NewWasmFunctions() *WasmFunctions {
	return &WasmFunctions{
		functions: make(map[string]func(...interface{}) (interface{}, error)),
	}
}
