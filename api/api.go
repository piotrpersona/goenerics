package api

import (
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
	"io/ioutil"
)

type Controller[B, R any] interface {
	Handle(body B, request *http.Request) (response R, status int, err error)
}

func HandleJSON[B, R any](controller Controller[B, R]) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		bodyBytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, fmt.Sprintf("error reading body bytes: %s", err), http.StatusBadRequest)
			return
		}
		defer request.Body.Close()

		request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		var requestBody B
		err = json.Unmarshal(bodyBytes, &requestBody)
		if err != nil {
			http.Error(writer, fmt.Sprintf("error unmarshaling body: %s", err), http.StatusInternalServerError)
			return
		}
		
		response, status, err := controller.Handle(requestBody, request)
		if err != nil {
			http.Error(writer, fmt.Sprintf("error from controller: %s", err), http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(status)
		err = json.NewEncoder(writer).Encode(response)
		if err != nil {
			http.Error(writer, fmt.Sprintf("error encoding response controller: %s", err), http.StatusInternalServerError)
			return
		}
	}
}