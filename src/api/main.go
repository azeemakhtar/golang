package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/getgoing/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received")
		w.Write([]byte("Hello World! GetGoing"))
	})
	mux.HandleFunc("/getgoing/go", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received")
		fmt.Println(r.Method)
		w.Write([]byte("Howdy, Going Go"))
	})

	http.ListenAndServe("localhost:3000", mux)
	http.ListenAndServe("localhost:3000", mux)
}
