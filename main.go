package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to my Go Web Server!</h1>")
}

func main() {
	http.HandleFunc("/", handler) // Set route for the homepage

	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil) // Listen on port 8080
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
