package controller

import (
	"net/http"
	"toko-buku/model/web"

	"github.com/julienschmidt/httprouter"
)

type TransaksiController interface {
	Create(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
	// Update(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
	// Delete(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
	FetchMetaData(r *http.Request, keranjang web.KeranjangResponse) web.KeranjangResponse
	FetchAdmin(r *http.Request, transaksi web.TransaksiResponse) web.TransaksiResponse
	FetchMember(r *http.Request, transaksi web.TransaksiResponse) web.TransaksiResponse
	ReduceStok(r *http.Request, keranjang web.KeranjangResponse)
	FindById(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
}
