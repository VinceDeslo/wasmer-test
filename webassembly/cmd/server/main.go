package main

import (
	"log"
	"net/http"
)

func main() {
	address := "http://localhost"
	port := ":9090"
	fs := http.FileServer(http.Dir("../../assets"))

	log.Printf("Starting server at %s\n", address+port)

	err := http.ListenAndServe(port, fs)
	if err != nil {
		log.Println("Failed to start server", err)
		return
	}
}
