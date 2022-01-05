package controllers

import (
	"github.com/piotrpersona/goenerics/examples/http/api"
	"net/http"
	"strings"
	"fmt"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

type Customer struct{}

func NewCustomer() *Customer {
	return &Customer{}
}

func NewCustomerHandler() http.HandlerFunc {
	return api.HandleJSON[Request, Response](NewCustomer())
}

func (c Customer) Handle(body Request, request *http.Request) (response Response, status int, err error) {
	response = Response{
		Message: fmt.Sprintf("customer %s has been created", strings.ToTitle(body.Name)),
	}
	status = http.StatusCreated
	return
}
