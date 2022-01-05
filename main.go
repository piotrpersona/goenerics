package main

import (
	"goenerics/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.NewCustomerHandler())
	http.ListenAndServe(":8080", nil)
}
