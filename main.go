package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir(".")) // Serve files from "static" directory
	http.Handle("/", fs)

	fmt.Println("Server running on :5098...")
	err := http.ListenAndServe(":5098", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
