package controller

import (
	"net/http"
	"toko-buku/helper"
	"toko-buku/model/web"
	"toko-buku/service"

	"github.com/julienschmidt/httprouter"
)

type BukuControllerImpl struct {
	BukuService service.BukuService
}

func NewBukuController(bukuService service.BukuService) BukuController{
	return &BukuControllerImpl{
		BukuService: bukuService,
	}
}

func(controller *BukuControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	BukuCreateRequest := web.BukuCreateRequest{}
	helper.ReadFromRequestBody(request, &BukuCreateRequest)

	bukuResponse := controller.BukuService.Create(request.Context(), BukuCreateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: bukuResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func(controller *BukuControllerImpl) Update(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	bukuUpdateRequest := web.BukuUpdateRequest{}
	helper.ReadFromRequestBody(request, &bukuUpdateRequest)
	bukuId := params.ByName("bukuId")
	
	bukuUpdateRequest.Id = bukuId

	bukuResponse := controller.BukuService.Update(request.Context(),bukuUpdateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: bukuResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func(controller *BukuControllerImpl) Delete(writer http.ResponseWriter, request *http.Request,params httprouter.Params){
	bukuId := params.ByName("bukuId")

	controller.BukuService.Delete(request.Context(),bukuId)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func(controller *BukuControllerImpl) FindById(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	bukuId := params.ByName("bukuId")

	bukuResponse := controller.BukuService.FindById(request.Context(), bukuId)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: bukuResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func(controller *BukuControllerImpl) FindAll(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	bukuResponses := controller.BukuService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: bukuResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}