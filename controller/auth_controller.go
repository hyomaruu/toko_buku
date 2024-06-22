package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController interface {
	Authentication(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
}
