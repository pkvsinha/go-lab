package main

import (
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you have requested: %s\n", r.URL.Path)
}

func attachRequestHandler() {
	http.HandleFunc("/", handleRequest)
}

func StartServer() {
	attachRequestHandler()
	fmt.Println("listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
