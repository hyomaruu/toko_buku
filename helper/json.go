package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestFormData(request *http.Request, result interface{}){

	// ReadFromRequestBody


	// decoder := json.NewDecoder(request.Body)
	// err := decoder.Decode(result)
	// PanicIfError(err)
}


func ReadFromRequestBody(request *http.Request, result interface{}){
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}){
	writer.Header().Add("Content-Type", "Application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}