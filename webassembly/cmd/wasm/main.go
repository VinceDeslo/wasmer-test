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
