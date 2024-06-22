package controller

import (
	"net/http"
	"os"
	"strconv"
	"toko-buku/helper"
	"toko-buku/model/web"
	"toko-buku/service"

	"github.com/julienschmidt/httprouter"
)

type TransaksiControllerImpl struct {
	TransaksiService service.TransaksiService
	KeranjangService service.KeranjangService
	BukuService service.BukuService
	AksesorisService service.AksesorisService
	BookDirectoryPath string
	AccessoryDirectoryPath string

}

func NewTransaksiController(transaksiService service.TransaksiService, keranjangService service.KeranjangService, bukuService service.BukuService, aksesorisService service.AksesorisService) TransaksiController {
	return &TransaksiControllerImpl{
		TransaksiService: transaksiService,
		KeranjangService: keranjangService,
		BukuService: bukuService,
		AksesorisService: aksesorisService,
		BookDirectoryPath: os.Getenv("BOOK_DIRECTORY_PATH"),
		AccessoryDirectoryPath: os.Getenv("ACCSESSORY_DIRECTORY_PATH"),
	}
}

func (controller *TransaksiControllerImpl) FetchMetaData(request *http.Request, keranjang web.KeranjangResponse) web.KeranjangResponse {
	if (keranjang.TipeBarang == "AKSESORIS") {
		aksesorisResponse := controller.AksesorisService.FindById(request.Context(), keranjang.KodeBarang)
		keranjang.Barang = aksesorisResponse
	}

	if (keranjang.TipeBarang == "BUKU") {
		BukuResponse := controller.BukuService.FindById(request.Context(), keranjang.KodeBarang)
		keranjang.Barang = BukuResponse
	}

	if dataAksesoris, ok := keranjang.Barang.(web.AksesorisResponse); ok {
		dataAksesoris.Image = request.Host + "/"+ controller.AccessoryDirectoryPath +"/"  + dataAksesoris.Image
		keranjang.Barang = dataAksesoris
    }

	if dataBuku, ok := keranjang.Barang.(web.BukuResponse); ok {
		dataBuku.Image = request.Host + "/"+ controller.BookDirectoryPath +"/"  + dataBuku.Image
		keranjang.Barang = dataBuku
    }

	return keranjang
}

func (controller *TransaksiControllerImpl) ReduceStok(request *http.Request, keranjang web.KeranjangResponse) {
	if (keranjang.TipeBarang == "AKSESORIS") {
		controller.AksesorisService.ReduceStok(request.Context(), keranjang.TotalBarang, keranjang.KodeBarang)
	}

	if (keranjang.TipeBarang == "BUKU") {
		controller.BukuService.ReduceStok(request.Context(), keranjang.TotalBarang, keranjang.KodeBarang)
	}
}

func (controller *TransaksiControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	TransaksiCreateRequest := web.TransaksiCreateRequest{}
	helper.ReadFromRequestBody(request, &TransaksiCreateRequest)

	transaksiResponse := controller.TransaksiService.Create(request.Context(), TransaksiCreateRequest)

	for _, keranjangId := range TransaksiCreateRequest.ListKeranjang {
		controller.KeranjangService.UpdateKodeTransaksi(request.Context(), keranjangId, int(transaksiResponse.KodeTransaksi) )
	}
	
	transaksiResponse = controller.TransaksiService.FindById(request.Context(), int(transaksiResponse.KodeTransaksi))
	keranjangResponses := controller.KeranjangService.FindByKodeTransaksi(request.Context(), int(transaksiResponse.KodeTransaksi))

	for i, keranjangResponse := range keranjangResponses {
		controller.ReduceStok(request, keranjangResponse)
		keranjangResponses[i] = controller.FetchMetaData(request, keranjangResponse)
	}

	transaksiResponse.ListKeranjang = keranjangResponses

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transaksiResponse,
		Message: "Transaksi successfully created",
	}

	helper.WriteToResponseBody(writer, webResponse)
}



func (controller *TransaksiControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transaksiId := params.ByName("transaksiId")
	id, err := strconv.Atoi(transaksiId)
	helper.PanicIfError(err)

	transaksiResponse := controller.TransaksiService.FindById(request.Context(), id)
	keranjangResponses := controller.KeranjangService.FindByKodeTransaksi(request.Context(), int(transaksiResponse.KodeTransaksi))

	for i, keranjangResponse := range keranjangResponses {
		keranjangResponses[i] = controller.FetchMetaData(request, keranjangResponse)
	}

	transaksiResponse.ListKeranjang = keranjangResponses
	
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transaksiResponse,
		Message: "Successfully fetch data",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransaksiControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transaksiResponses := controller.TransaksiService.FindAll(request.Context())

	for i, transaksiResponse := range transaksiResponses {
		keranjangResponses := controller.KeranjangService.FindByKodeTransaksi(request.Context(), int(transaksiResponse.KodeTransaksi))
	
		for x, keranjangResponse := range keranjangResponses {
			keranjangResponses[x] = controller.FetchMetaData(request, keranjangResponse)
		}
	
		transaksiResponses[i].ListKeranjang = keranjangResponses
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transaksiResponses,
		Message: "Successfully fetch data",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
