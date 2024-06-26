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

type KeranjangControllerImpl struct {
	KeranjangService service.KeranjangService
	BukuService service.BukuService
	AksesorisService service.AksesorisService
	BookDirectoryPath string
	AccessoryDirectoryPath string

}

func NewKeranjangController(keranjangService service.KeranjangService, bukuService service.BukuService, aksesorisService service.AksesorisService) KeranjangController {
	return &KeranjangControllerImpl{
		KeranjangService: keranjangService,
		BukuService: bukuService,
		AksesorisService: aksesorisService,
		BookDirectoryPath: os.Getenv("BOOK_DIRECTORY_PATH"),
		AccessoryDirectoryPath: os.Getenv("ACCSESSORY_DIRECTORY_PATH"),
	}
}

func (controller *KeranjangControllerImpl) FetchMetaData(request *http.Request, keranjang web.KeranjangResponse) web.KeranjangResponse {
	if (keranjang.TipeBarang == "AKSESORIS") {
		aksesorisResponse := controller.AksesorisService.FindById(request.Context(), keranjang.KodeBarang)
		keranjang.Barang = aksesorisResponse
	}

	if (keranjang.TipeBarang == "BUKU") {
		BukuResponse := controller.BukuService.FindById(request.Context(), keranjang.KodeBarang)
		keranjang.Barang = BukuResponse
	}

	if dataAksesoris, ok := keranjang.Barang.(web.AksesorisResponse); ok {
		dataAksesoris.Image = "http://" + request.Host + "/"+ controller.AccessoryDirectoryPath +"/"  + dataAksesoris.Image
		keranjang.Barang = dataAksesoris
    }

	if dataBuku, ok := keranjang.Barang.(web.BukuResponse); ok {
		dataBuku.Image = "http://" + request.Host + "/"+ controller.BookDirectoryPath +"/"  + dataBuku.Image
		keranjang.Barang = dataBuku
    }

	return keranjang
}

func (controller *KeranjangControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	KeranjangCreateRequest := web.KeranjangCreateRequest{}
	helper.ReadFromRequestBody(request, &KeranjangCreateRequest)
	
	if (KeranjangCreateRequest.TipeBarang == "AKSESORIS") {
		KeranjangCreateRequest.Harga = controller.AksesorisService.FindById(request.Context(), KeranjangCreateRequest.KodeBarang).Harga
	}

	if (KeranjangCreateRequest.TipeBarang == "BUKU") {
		KeranjangCreateRequest.Harga = controller.BukuService.FindById(request.Context(), KeranjangCreateRequest.KodeBarang).Harga
	}

	keranjangResponse := controller.KeranjangService.Create(request.Context(), KeranjangCreateRequest)

	keranjangResponse = controller.FetchMetaData(request, keranjangResponse)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   keranjangResponse,
		Message: "Keranjang successfully created",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KeranjangControllerImpl) UpdateTotalBarang(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	keranjangUpdateStockRequest := web.KeranjangUpdateStockRequest{}
	helper.ReadFromRequestBody(request, &keranjangUpdateStockRequest)

	keranjangId := params.ByName("keranjangId")
	id, err := strconv.Atoi(keranjangId)
	helper.PanicIfError(err)

	keranjangResponse := controller.KeranjangService.UpdateTotalBarang(request.Context(), keranjangUpdateStockRequest, id)
	
	keranjangResponse = controller.FetchMetaData(request, keranjangResponse)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   keranjangResponse,
		Message: "Keranjang successfully updated",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KeranjangControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	keranjangId := params.ByName("keranjangId")
	id, err := strconv.Atoi(keranjangId)
	helper.PanicIfError(err)

	controller.KeranjangService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Message: "Keranjang successfully deleted",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KeranjangControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	keranjangId := params.ByName("keranjangId")
	id, err := strconv.Atoi(keranjangId)
	helper.PanicIfError(err)

	keranjangResponse := controller.KeranjangService.FindById(request.Context(), id)

	keranjangResponse = controller.FetchMetaData(request, keranjangResponse)
	
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   keranjangResponse,
		Message: "Successfully fetch data",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KeranjangControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	keranjangResponses := controller.KeranjangService.FindAll(request.Context())

	for i, keranjangResponse := range keranjangResponses {
		keranjangResponses[i] = controller.FetchMetaData(request, keranjangResponse)
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   keranjangResponses,
		Message: "Successfully fetch data",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
