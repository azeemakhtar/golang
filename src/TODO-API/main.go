package main

import (
	"encoding/json"
	"net/http"

	views "command-line-arguments/Users/azeemakhtar/golang/src/TODO-API/view/structs.go"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			json.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe(":3000", mux)
}
