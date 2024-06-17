package controller

import (
	"net/http"
	"toko-buku/model/web"

	"github.com/julienschmidt/httprouter"
)

type KeranjangController interface {
	FetchMetaData(r *http.Request, keranjang web.KeranjangResponse) web.KeranjangResponse
	Create(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateTotalBarang(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
}
