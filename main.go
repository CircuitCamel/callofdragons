package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir(".")) // Serve files from "static" directory
	http.Handle("/", fs)

	fmt.Println("Server running on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
