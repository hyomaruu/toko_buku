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

type AksesorisControllerImpl struct {
	AksesorisService service.AksesorisService
	AccessoryDirectoryPath string
}

func NewAksesorisController(aksesorisService service.AksesorisService) AksesorisController{
	return &AksesorisControllerImpl{
		AksesorisService: aksesorisService,
		AccessoryDirectoryPath: os.Getenv("ACCSESSORY_DIRECTORY_PATH"),
	}
}

func(controller *AksesorisControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	request.ParseMultipartForm(0)

	val := helper.HandleUploadImage(request, controller.AccessoryDirectoryPath, "gambar")
	hargaFormValue := request.FormValue("harga")
	harga, _ := strconv.Atoi(hargaFormValue)
	
	aksesorisCreateRequest := web.AksesorisCreateRequest{
		Id: request.FormValue("kode"),
		Name: request.FormValue("nama_barang"),
		Image: val,
		Deskripsi: request.FormValue("deskripsi"),
		Harga: harga,
	}

	aksesorisResponse := controller.AksesorisService.Create(request.Context(), aksesorisCreateRequest)
	aksesorisResponse.Image = "http://" + request.Host + "/"+ controller.AccessoryDirectoryPath +"/" + aksesorisResponse.Image
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: aksesorisResponse,
		Message: "Accessory successfully created",
	}
	helper.WriteToResponseBody(writer,webResponse)
}

func(controller *AksesorisControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	request.ParseMultipartForm(0)

	val := helper.HandleUploadImage(request, controller.AccessoryDirectoryPath, "gambar")
	hargaFormValue := request.FormValue("harga")
	harga, _ := strconv.Atoi(hargaFormValue)

	id := params.ByName("aksesorisId")
	
	aksesorisUpdateRequest := web.AksesorisUpdateRequest{
		Id: request.FormValue("kode"),
		Name: request.FormValue("nama_barang"),
		Image: val,
		Deskripsi: request.FormValue("deskripsi"),
		Harga: harga,
	}

	aksesorisResponse := controller.AksesorisService.Update(request.Context(), aksesorisUpdateRequest, id)
	aksesorisResponse.Image = "http://" + request.Host + "/"+ controller.AccessoryDirectoryPath +"/"  + aksesorisResponse.Image

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: aksesorisResponse,
		Message: "Accessory successfully updated",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func(controller *AksesorisControllerImpl) UpdateStok(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	aksesorisUpdateStockRequest := web.AksesorisUpdateStockRequest{}
	helper.ReadFromRequestBody(request, &aksesorisUpdateStockRequest)

	id := params.ByName("aksesorisId")

	aksesorisResponse := controller.AksesorisService.UpdateStok(request.Context(), aksesorisUpdateStockRequest, id)
	aksesorisResponse.Image = "http://" + request.Host + "/"+ controller.AccessoryDirectoryPath +"/"  + aksesorisResponse.Image

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: aksesorisResponse,
		Message: "Accessory stock successfully updated",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func(controller *AksesorisControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	id := params.ByName("aksesorisId")

	controller.AksesorisService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Message: "Accessory successfully deleted",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func(controller *AksesorisControllerImpl) FindById(writer http.ResponseWriter, request *http.Request,params httprouter.Params){
	id := params.ByName("aksesorisId")

	aksesorisResponse := controller.AksesorisService.FindById(request.Context(), id)
	aksesorisResponse.Image = "http://" + request.Host +  "/"+ controller.AccessoryDirectoryPath +"/"  + aksesorisResponse.Image
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: aksesorisResponse,
		Message: "Successfully fetch data",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func(controller *AksesorisControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request,params httprouter.Params){
	FetchRequest := web.FetchRequest{}
	helper.ReadQueryParams(request, &FetchRequest)

	aksesorisResponses := controller.AksesorisService.FindAll(request.Context(),FetchRequest)

	for i, aksesoris := range aksesorisResponses {
		aksesorisResponses[i].Image = "http://" + request.Host +  "/" + controller.AccessoryDirectoryPath + "/"  + aksesoris.Image
	}

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: aksesorisResponses,
		Message: "Successfully fetch data",
	}
	helper.WriteToResponseBody(writer,webResponse)
}