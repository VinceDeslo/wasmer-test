package main

import (
	"log"
)

func main() {
	PrintMessage()
}

//go:export PrintMessage
func PrintMessage() {
	msg := "Welcome to Go Web Assembly"
	log.Println(msg)
}

//go:export PrintDescription
func PrintDescription() {
	msg := "This program is fetching WASM functions in a loop from specified file path and function names."
	log.Println(msg)
}
