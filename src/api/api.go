package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		w.Write([]byte("Hello World! GetGoing"))
	})
	http.ListenAndServe("localhost:3000", mux)
}
