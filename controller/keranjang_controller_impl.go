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
	BookDirectoryPath string
	AccessoryDirectoryPath string

}

func NewKeranjangController(keranjangService service.KeranjangService) KeranjangController {
	return &KeranjangControllerImpl{
		KeranjangService: keranjangService,
		BookDirectoryPath: os.Getenv("BOOK_DIRECTORY_PATH"),
		AccessoryDirectoryPath: os.Getenv("ACCSESSORY_DIRECTORY_PATH"),
	}
}

func (controller *KeranjangControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	KeranjangCreateRequest := web.KeranjangCreateRequest{}
	helper.ReadFromRequestBody(request, &KeranjangCreateRequest)

	keranjangResponse := controller.KeranjangService.Create(request.Context(), KeranjangCreateRequest)

	if dataAksesoris, ok := keranjangResponse.Data.(web.AksesorisResponse); ok {
		dataAksesoris.Image = request.Host + "/"+ controller.AccessoryDirectoryPath +"/"  + dataAksesoris.Image
		keranjangResponse.Data = dataAksesoris
    }

	if dataBuku, ok := keranjangResponse.Data.(web.BukuResponse); ok {
		dataBuku.Image = request.Host + "/"+ controller.BookDirectoryPath +"/"  + dataBuku.Image
		keranjangResponse.Data = dataBuku
    }

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

	if dataAksesoris, ok := keranjangResponse.Data.(web.AksesorisResponse); ok {
		dataAksesoris.Image = request.Host + "/"+ controller.AccessoryDirectoryPath +"/"  + dataAksesoris.Image
		keranjangResponse.Data = dataAksesoris
    }

	if dataBuku, ok := keranjangResponse.Data.(web.BukuResponse); ok {
		dataBuku.Image = request.Host + "/"+ controller.BookDirectoryPath +"/"  + dataBuku.Image
		keranjangResponse.Data = dataBuku
    }

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

	if dataAksesoris, ok := keranjangResponse.Data.(web.AksesorisResponse); ok {
		dataAksesoris.Image = request.Host + "/"+ controller.AccessoryDirectoryPath +"/"  + dataAksesoris.Image
		keranjangResponse.Data = dataAksesoris
    }

	if dataBuku, ok := keranjangResponse.Data.(web.BukuResponse); ok {
		dataBuku.Image = request.Host + "/"+ controller.BookDirectoryPath +"/"  + dataBuku.Image
		keranjangResponse.Data = dataBuku
    }
	
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
		if dataAksesoris, ok := keranjangResponse.Data.(web.AksesorisResponse); ok {
			dataAksesoris.Image = request.Host + "/"+ controller.AccessoryDirectoryPath +"/"  + dataAksesoris.Image
			keranjangResponses[i].Data = dataAksesoris
		}
	
		if dataBuku, ok := keranjangResponse.Data.(web.BukuResponse); ok {
			dataBuku.Image = request.Host + "/"+ controller.BookDirectoryPath +"/"  + dataBuku.Image
			keranjangResponses[i].Data = dataBuku
		}
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   keranjangResponses,
		Message: "Successfully fetch data",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
