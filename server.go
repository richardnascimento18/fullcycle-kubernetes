package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello, I am %s and I am %s years old.", name, age)
}
