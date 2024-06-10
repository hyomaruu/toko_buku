package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MemberController interface {
	Create(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
}



// package controller

// import (
// 	"fmt"
// 	"net/http"
// 	"toko-buku/helper"
// 	"toko-buku/model/web"
// 	"toko-buku/service"

// 	"github.com/julienschmidt/httprouter"
// )

// type BukuControllerImpl struct {
// 	BukuService service.BukuService
// }

// func NewBukuController(bukuService service.BukuService) BukuController{
// 	return &BukuControllerImpl{
// 		BukuService: bukuService,
// 	}
// }

// func(controller *BukuControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
// 	bukuCreateRequest := web.BukuCreateRequest{}

// 	helper.ReadFromRequestBody(request, &bukuCreateRequest)

// 	fmt.Println(request)

// 	// val := helper.HandleUploadImage(request, "gambar")

// 	// bukuCreateRequest.Image = val

// 	bukuResponse := controller.BukuService.Create(request.Context(), bukuCreateRequest)
// 	webResponse := web.WebResponse{
// 		Code: 200,
// 		Status: "OK",
// 		Data: bukuResponse,
// 	}

// 	helper.WriteToResponseBody(writer, webResponse)
// }

// func(controller *BukuControllerImpl) Update(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
// 	bukuUpdateRequest := web.BukuUpdateRequest{}
	
// 	helper.ReadFromRequestBody(request, &bukuUpdateRequest)

// 	bukuId := params.ByName("bukuId")

// 	val := helper.HandleUploadImage(request, "gambar")

// 	bukuUpdateRequest.Image = val
	
// 	bukuResponse := controller.BukuService.Update(request.Context(), bukuUpdateRequest, bukuId)
// 	webResponse := web.WebResponse{
// 		Code: 200,
// 		Status: "OK",
// 		Data: bukuResponse,
// 	}

// 	helper.WriteToResponseBody(writer, webResponse)
// }

// func(controller *BukuControllerImpl) Delete(writer http.ResponseWriter, request *http.Request,params httprouter.Params){
// 	bukuId := params.ByName("bukuId")

// 	controller.BukuService.Delete(request.Context(), bukuId)
// 	webResponse := web.WebResponse{
// 		Code: 200,
// 		Status: "OK",
// 	}
// 	helper.WriteToResponseBody(writer, webResponse)
// }

// func(controller *BukuControllerImpl) FindById(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
// 	bukuId := params.ByName("bukuId")

// 	bukuResponse := controller.BukuService.FindById(request.Context(), bukuId)
// 	webResponse := web.WebResponse{
// 		Code: 200,
// 		Status: "OK",
// 		Data: bukuResponse,
// 	}
// 	helper.WriteToResponseBody(writer, webResponse)
// }

// func(controller *BukuControllerImpl) FindAll(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
// 	bukuResponses := controller.BukuService.FindAll(request.Context())
// 	webResponse := web.WebResponse{
// 		Code: 200,
// 		Status: "OK",
// 		Data: bukuResponses,
// 	}
// 	helper.WriteToResponseBody(writer, webResponse)
// }