package main

import (
	controller "command-line-arguments/Users/azeemakhtar/golang/src/TODO-API/control/rout.go"
	"net/http"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe(":3000", mux)
}
