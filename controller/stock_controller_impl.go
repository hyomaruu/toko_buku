package controller

import (
	"net/http"
	"toko-buku/helper"
	"toko-buku/model/web"
	"toko-buku/service"

	"github.com/julienschmidt/httprouter"
)

type StockControllerImpl struct {
	StockService service.StockService
}

func NewStockController(stokService service.StockService) StockController {
	return &StockControllerImpl{
		StockService: stokService,
	}
}

func (controller *StockControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	StockCreateRequest := web.StockCreateRequest{}
	helper.ReadFromRequestBody(request, &StockCreateRequest)

	stockResponse := controller.StockService.Create(request.Context(), StockCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   stockResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StockControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	stockUpdateRequest := web.StockUpdateRequest{}
	helper.ReadFromRequestBody(request, &stockUpdateRequest)

	id := params.ByName("stockId")

	stockResponse := controller.StockService.Update(request.Context(), stockUpdateRequest, id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   stockResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StockControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("stockId")

	controller.StockService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StockControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("stockId")

	stockResponse := controller.StockService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   stockResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StockControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	stockResponses := controller.StockService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   stockResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
