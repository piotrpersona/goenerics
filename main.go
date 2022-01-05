package main

import (
	// "bytes"
	"fmt"
	"net/http"
	// "encoding/json"
)

type Controller[B, R any] interface {
	Handle(body B, request *http.Request) (response R, err error)
}

func HandleJSON[B, R any](controller Controller[B, R]) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// var bodyBuffer bytes.Buffer
		// err := json.NewDecoder(&bodyBuffer).Decode(request.Body)
		// if err != nil {
		// 	http.Error(writer, fmt.Sprintf("error decoding body: %s", err), http.StatusBadRequest)
		// 	return
		// }
		// defer request.Body.Close()

		// var requestBody B
		// bodyBytes := bodyBuffer.Bytes()
		// err = json.Unmarshal(bodyBytes, &requestBody)
		// if err != nil {
		// 	http.Error(writer, fmt.Sprintf("error unmarshaling body: %s", err), http.StatusInternalServerError)
		// 	return
		// }
		
		// var requestBody B
		// response, err := controller.Handle(requestBody, request)
		// if err != nil {
		// 	http.Error(writer, fmt.Sprintf("error from controller: %s", err), http.StatusInternalServerError)
		// 	return
		// }

		// writer.WriteHeader(http.StatusOK)
		// err = json.NewEncoder(writer).Encode(response)
		// if err != nil {
		// 	http.Error(writer, fmt.Sprintf("error encoding response controller: %s", err), http.StatusInternalServerError)
		// 	return
		// }
	}
}

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

type CustomController struct {}

func (c CustomController) Handle(body Request, request *http.Request) (response Response, err error) {
	response = Response{Message: body.Name}
	return 
}

func main() {
	controller := CustomController{}
	handler := HandleJSON[Request, Response](controller.Handle)
	fmt.Println(handler)
}
