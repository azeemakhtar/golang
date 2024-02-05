package controller

import (
	"encoding/json"
	"net/http"
	"github.com/" // Add the import statement for the package that contains the Registered function
)

func main() {
	mux := package.Registered() // Call the Registered function from the imported package

	http.ListenAndServe(":3000", mux)
}
