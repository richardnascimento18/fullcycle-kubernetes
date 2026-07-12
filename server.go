package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hello)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello FullCycle!!!</h1>"))
}
