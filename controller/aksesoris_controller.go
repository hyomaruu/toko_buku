package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AksesorisController interface {
	Create(writer http.ResponseWriter, r *http.Request,params httprouter.Params)
	Update(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateStok(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, r  *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, r *http.Request,params httprouter.Params)
}