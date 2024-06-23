package controller

import (
	"net/http"
	"strconv"
	"toko-buku/helper"
	"toko-buku/model/web"
	"toko-buku/service"

	"github.com/julienschmidt/httprouter"
)

type AdminControllerImpl struct {
	AdminService service.AdminService
}

func NewAdminController(adminService service.AdminService) AdminController {
	return &AdminControllerImpl{
		AdminService: adminService,
	}
}

func (controller *AdminControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	// writer.Header().Set("Access-Control-Allow-Methods", "POST")
	// writer.Header().Set("Access-Control-Allow-Headers", "*")

	// if request.Method == "OPTIONS" {
	// 	writer.Write([]byte("allowed"))
	// 	return
	// }

	AdminCreateRequest := web.AdminCreateRequest{}
	helper.ReadFromRequestBody(request, &AdminCreateRequest)

	adminResponse := controller.AdminService.Create(request.Context(), AdminCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   adminResponse,
		Message: "Admin successfully created",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AdminControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	adminUpdateRequest := web.AdminUpdateRequest{}
	helper.ReadFromRequestBody(request, &adminUpdateRequest)
 
	adminId := params.ByName("adminId")
	id, err := strconv.Atoi(adminId)
	helper.PanicIfError(err)

	adminResponse := controller.AdminService.Update(request.Context(), adminUpdateRequest, id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   adminResponse,
		Message: "Admin successfully updated",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AdminControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	adminId := params.ByName("adminId")
	id, err := strconv.Atoi(adminId)
	helper.PanicIfError(err)

	controller.AdminService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Message: "Admin successfully deleted",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AdminControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	adminId := params.ByName("adminId")
	id, err := strconv.Atoi(adminId)
	helper.PanicIfError(err)

	adminResponse := controller.AdminService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   adminResponse,
		Message: "Successfully fetch data",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AdminControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	adminResponses := controller.AdminService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   adminResponses,
		Message: "Successfully fetch data",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
