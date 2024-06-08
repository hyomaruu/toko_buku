package controller

import (
	"net/http"
	"strconv"
	"toko-buku/helper"
	"toko-buku/model/web"
	"toko-buku/service"

	"github.com/julienschmidt/httprouter"
)

type MemberControllerImpl struct {
	MemberService service.MemberService
}

func NewMemberController(memberService service.MemberService) MemberController {
	return &MemberControllerImpl{
		MemberService: memberService,
	}
}

func (controller *MemberControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	MemberCreateRequest := web.MemberCreateRequest{}
	helper.ReadFromRequestBody(request, &MemberCreateRequest)

	memberResponse := controller.MemberService.Create(request.Context(), MemberCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   memberResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MemberControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	memberUpdateRequest := web.MemberUpdateRequest{}
	helper.ReadFromRequestBody(request, &memberUpdateRequest)

	memberId := params.ByName("memberId")
	id, err := strconv.Atoi(memberId)
	helper.PanicIfError(err)

	memberUpdateRequest.Id = id

	memberResponse := controller.MemberService.Update(request.Context(), memberUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   memberResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MemberControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	memberId := params.ByName("memberId")
	id, err := strconv.Atoi(memberId)
	helper.PanicIfError(err)

	controller.MemberService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MemberControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	memberId := params.ByName("memberId")
	id, err := strconv.Atoi(memberId)
	helper.PanicIfError(err)

	memberResponse := controller.MemberService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   memberResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MemberControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	memberResponses := controller.MemberService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   memberResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
