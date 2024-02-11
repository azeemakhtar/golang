package main

import (
	controller "/Users/azeemakhtar/golang/src/TODO-API/control/rout.go"
	"net/http"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe(":3000", mux)
}
