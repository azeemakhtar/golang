package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received")
		w.Write([]byte("Hello, World!"))
	})
	http.ListenAndServe(":3000", mux)
}
