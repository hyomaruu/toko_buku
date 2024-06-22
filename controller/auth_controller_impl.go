package controller

import (
	"net/http"
	"toko-buku/helper"
	"toko-buku/model/web"
	"toko-buku/service"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) Authentication(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	AuthRequest := web.AuthRequest{}
	helper.ReadFromRequestBody(request, &AuthRequest)

	authResponse := controller.AuthService.Authentication(request.Context(), AuthRequest)
	
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   authResponse,
		Message: "Auth successfully fetch",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
