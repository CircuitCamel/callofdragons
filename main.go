package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir(".")) // Serve files from "static" directory
	http.Handle("/", fs)

	fmt.Println("Server running on :5098...")
	err := http.ListenAndServeTLS(":5098", "/etc/letsencrypt/live/el7ossen.uk/fullchain.pem", "/etc/letsencrypt/live/el7ossen.uk/privkey.pem", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
