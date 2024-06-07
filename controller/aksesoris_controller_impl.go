package controller

import (
	"net/http"
	"toko-buku/helper"
	"toko-buku/model/web"
	"toko-buku/service"

	"github.com/julienschmidt/httprouter"
)

type AksesorisControllerImpl struct {
	AksesorisService service.AksesorisService
}

func NewAksesorisController(aksesorisService service.AksesorisService) AksesorisController{
	return &AksesorisControllerImpl{
		AksesorisService: aksesorisService,
	}
}

func(controller *AksesorisControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	AksesorisCreateRequest := web.AksesorisCreateRequest{}
	helper.ReadFromRequestBody(request, &AksesorisCreateRequest)

	aksesorisResponse := controller.AksesorisService.Create(request.Context(), AksesorisCreateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: aksesorisResponse,
	}
	helper.WriteToResponseBody(writer,webResponse)
}

func(controller *AksesorisControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	aksesorisUpdateRequest := web.AksesorisUpdateRequest{}
	helper.ReadFromRequestBody(request, &aksesorisUpdateRequest)
	aksesorisId := params.ByName("aksesorisId")

	aksesorisUpdateRequest.Id = aksesorisId
	aksesorisResponse := controller.AksesorisService.Update(request.Context(),aksesorisUpdateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: aksesorisResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func(controller *AksesorisControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	aksesorisId := params.ByName("aksesorisId")

	controller.AksesorisService.Delete(request.Context(),aksesorisId)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func(controller *AksesorisControllerImpl) FindById(writer http.ResponseWriter, request *http.Request,params httprouter.Params){
	aksesorisId := params.ByName("aksesorisId")

	aksesorisResponse := controller.AksesorisService.FindById(request.Context(), aksesorisId)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: aksesorisResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func(controller *AksesorisControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request,params httprouter.Params){
	aksesorisResponses := controller.AksesorisService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: aksesorisResponses,
	}
	helper.WriteToResponseBody(writer,webResponse)
}