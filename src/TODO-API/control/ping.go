package controller

import (
	views "command-line-arguments/Users/azeemakhtar/golang/src/TODO-API/view/structs.go"
	"encoding/json"
	"net/http"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			json.NewEncoder(w).Encode((data))
		}
	}
}
