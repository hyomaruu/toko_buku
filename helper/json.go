package helper

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
)

func ReadFromRequestBody(request *http.Request, result interface{}){
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func ReadQueryParams(request *http.Request, result interface{}){
	decoder := schema.NewDecoder()
	err := decoder.Decode(result, request.URL.Query())
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}){
	writer.Header().Add("Content-Type", "Application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}