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

type BukuControllerImpl struct {
	BukuService service.BukuService
	BookDirectoryPath string
}

func NewBukuController(bukuService service.BukuService) BukuController{
	return &BukuControllerImpl{
		BukuService: bukuService,
		BookDirectoryPath: os.Getenv("BOOK_DIRECTORY_PATH"),
	}
}

func(controller *BukuControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	request.ParseMultipartForm(0)

	val := helper.HandleUploadImage(request, controller.BookDirectoryPath, "gambar")
	hargaFormValue := request.FormValue("harga")
	harga, _ := strconv.Atoi(hargaFormValue)

	bukuCreateRequest := web.BukuCreateRequest{
		Id: request.FormValue("kode_buku"),
		Name: request.FormValue("nama_buku"),
		Penerbit: request.FormValue("penerbit"),
		Image: val,
		Genre: request.FormValue("genre"),
		Rilis: request.FormValue("rilis"),
		Harga: harga,
		Deskripsi: request.FormValue("deskripsi"),
		Pengarang: request.FormValue("pengarang"),
	}

	bukuResponse := controller.BukuService.Create(request.Context(), bukuCreateRequest)
	bukuResponse.Image = request.Host + "/"+ controller.BookDirectoryPath +"/" + bukuResponse.Image
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: bukuResponse,
		Message: "Book successfully created",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func(controller *BukuControllerImpl) Update(writer http.ResponseWriter, request *http.Request,params httprouter.Params){
	request.ParseMultipartForm(0)

	val := helper.HandleUploadImage(request, controller.BookDirectoryPath, "gambar")
	hargaFormValue := request.FormValue("harga")
	harga, _ := strconv.Atoi(hargaFormValue)

	bukuId := params.ByName("bukuId")
	
	bukuUpdateRequest := web.BukuUpdateRequest{
		Name: request.FormValue("nama_buku"),
		Penerbit: request.FormValue("penerbit"),
		Image: val,
		Genre: request.FormValue("genre"),
		Rilis: request.FormValue("rilis"),
		Deskripsi: request.FormValue("deskripsi"),
		Harga: harga,
		Pengarang: request.FormValue("pengarang"),
	}
	
	bukuResponse := controller.BukuService.Update(request.Context(), bukuUpdateRequest, bukuId)
	bukuResponse.Image = request.Host + "/"+ controller.BookDirectoryPath +"/" + bukuResponse.Image
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: bukuResponse,
		Message: "Book successfully updated",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func(controller *BukuControllerImpl) UpdateStok(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	bukuUpdateStockRequest := web.BukuUpdateStockRequest{}
	helper.ReadFromRequestBody(request, &bukuUpdateStockRequest)

	id := params.ByName("bukuId")

	bukuResponse := controller.BukuService.UpdateStok(request.Context(), bukuUpdateStockRequest, id)
	bukuResponse.Image = request.Host + "/"+ controller.BookDirectoryPath +"/"  + bukuResponse.Image

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: bukuResponse,
		Message: "Book stock successfully updated",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func(controller *BukuControllerImpl) Delete(writer http.ResponseWriter, request *http.Request,params httprouter.Params){
	bukuId := params.ByName("bukuId")

	controller.BukuService.Delete(request.Context(), bukuId)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Message: "Book successfully deleted",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func(controller *BukuControllerImpl) FindById(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	bukuId := params.ByName("bukuId")

	bukuResponse := controller.BukuService.FindById(request.Context(), bukuId)
	bukuResponse.Image = request.Host +  "/"+ controller.BookDirectoryPath +"/" + bukuResponse.Image
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: bukuResponse,
		Message: "Successfully fetch data",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func(controller *BukuControllerImpl) FindAll(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	bukuResponses := controller.BukuService.FindAll(request.Context())

	for i, buku := range bukuResponses {
		bukuResponses[i].Image = request.Host +  "/"+ controller.BookDirectoryPath +"/" + buku.Image
	}

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: bukuResponses,
		Message: "Successfully fetch data",
	}
	helper.WriteToResponseBody(writer, webResponse)
}