package main

import (
	"net/http"

	"github.com/piotrpersona/goenerics/examples/http/controllers"
)

func main() {
	http.HandleFunc("/", controllers.NewCustomerHandler())
	http.ListenAndServe(":8080", nil)
}
