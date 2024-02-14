package main

import (
	"command-line-arguments/Users/azeemakhtar/golang/src/TODO-API/structs/structs.go"
	"encoding/json"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			json.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe("localhost:3000", mux)
}
