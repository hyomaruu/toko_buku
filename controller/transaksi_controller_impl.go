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

type TransaksiControllerImpl struct {
	TransaksiService service.TransaksiService
	KeranjangService service.KeranjangService
	BukuService service.BukuService
	AksesorisService service.AksesorisService
	AdminService service.AdminService
	MemberService service.MemberService
	BookDirectoryPath string
	AccessoryDirectoryPath string

}

func NewTransaksiController(transaksiService service.TransaksiService, keranjangService service.KeranjangService, bukuService service.BukuService, aksesorisService service.AksesorisService, adminService service.AdminService, memberService service.MemberService) TransaksiController {
	return &TransaksiControllerImpl{
		TransaksiService: transaksiService,
		KeranjangService: keranjangService,
		BukuService: bukuService,
		AksesorisService: aksesorisService,
		AdminService: adminService,
		MemberService: memberService,
		BookDirectoryPath: os.Getenv("BOOK_DIRECTORY_PATH"),
		AccessoryDirectoryPath: os.Getenv("ACCSESSORY_DIRECTORY_PATH"),
	}
}

func (controller *TransaksiControllerImpl) FetchMetaData(request *http.Request, keranjang web.KeranjangResponse) web.KeranjangResponse {
	if (keranjang.TipeBarang == "AKSESORIS") {
		aksesorisResponse := controller.AksesorisService.FindById(request.Context(), keranjang.KodeBarang)
		keranjang.Barang = aksesorisResponse
	}

	if (keranjang.TipeBarang == "BUKU") {
		BukuResponse := controller.BukuService.FindById(request.Context(), keranjang.KodeBarang)
		keranjang.Barang = BukuResponse
	}

	if dataAksesoris, ok := keranjang.Barang.(web.AksesorisResponse); ok {
		dataAksesoris.Image = "http://" + request.Host + "/"+ controller.AccessoryDirectoryPath +"/"  + dataAksesoris.Image
		keranjang.Barang = dataAksesoris
    }

	if dataBuku, ok := keranjang.Barang.(web.BukuResponse); ok {
		dataBuku.Image = "http://" + request.Host + "/"+ controller.BookDirectoryPath +"/"  + dataBuku.Image
		keranjang.Barang = dataBuku
    }

	return keranjang
}

func (controller *TransaksiControllerImpl) FetchMember(request *http.Request, transaksi web.TransaksiResponse) web.TransaksiResponse {
	memberResponse := controller.MemberService.FindById(request.Context(), transaksi.NomorMember)
	transaksi.Member = memberResponse

	return transaksi
}

func (controller *TransaksiControllerImpl) FetchAdmin(request *http.Request, transaksi web.TransaksiResponse) web.TransaksiResponse {
	adminResponse := controller.AdminService.FindById(request.Context(), transaksi.IdAdmin)
	transaksi.Admin = adminResponse

	return transaksi
}

func (controller *TransaksiControllerImpl) ReduceStok(request *http.Request, keranjang web.KeranjangResponse) {
	if (keranjang.TipeBarang == "AKSESORIS") {
		controller.AksesorisService.ReduceStok(request.Context(), keranjang.TotalBarang, keranjang.KodeBarang)
	}

	if (keranjang.TipeBarang == "BUKU") {
		controller.BukuService.ReduceStok(request.Context(), keranjang.TotalBarang, keranjang.KodeBarang)
	}
}

func (controller *TransaksiControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	TransaksiCreateRequest := web.TransaksiCreateRequest{}
	helper.ReadFromRequestBody(request, &TransaksiCreateRequest)

	transaksiResponse := controller.TransaksiService.Create(request.Context(), TransaksiCreateRequest)

	for _, keranjangId := range TransaksiCreateRequest.ListKeranjang {
		id, _ := strconv.Atoi(keranjangId)
		controller.KeranjangService.UpdateKodeTransaksi(request.Context(), id, int(transaksiResponse.KodeTransaksi) )
	}
	
	transaksiResponse = controller.TransaksiService.FindById(request.Context(), int(transaksiResponse.KodeTransaksi))
	keranjangResponses := controller.KeranjangService.FindByKodeTransaksi(request.Context(), int(transaksiResponse.KodeTransaksi))

	for i, keranjangResponse := range keranjangResponses {
		controller.ReduceStok(request, keranjangResponse)
		keranjangResponses[i] = controller.FetchMetaData(request, keranjangResponse)
	}

	transaksiResponse.ListKeranjang = keranjangResponses
	transaksiResponse = controller.FetchAdmin(request,transaksiResponse)
	transaksiResponse = controller.FetchMember(request,transaksiResponse)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transaksiResponse,
		Message: "Transaksi successfully created",
	}

	helper.WriteToResponseBody(writer, webResponse)
}



func (controller *TransaksiControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transaksiId := params.ByName("transaksiId")
	id, err := strconv.Atoi(transaksiId)
	helper.PanicIfError(err)

	transaksiResponse := controller.TransaksiService.FindById(request.Context(), id)
	keranjangResponses := controller.KeranjangService.FindByKodeTransaksi(request.Context(), int(transaksiResponse.KodeTransaksi))

	for i, keranjangResponse := range keranjangResponses {
		keranjangResponses[i] = controller.FetchMetaData(request, keranjangResponse)
	}

	transaksiResponse.ListKeranjang = keranjangResponses
	transaksiResponse = controller.FetchAdmin(request,transaksiResponse)
	transaksiResponse = controller.FetchMember(request,transaksiResponse)
	
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transaksiResponse,
		Message: "Successfully fetch data",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransaksiControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transaksiResponses := controller.TransaksiService.FindAll(request.Context())

	for i, transaksiResponse := range transaksiResponses {
		keranjangResponses := controller.KeranjangService.FindByKodeTransaksi(request.Context(), int(transaksiResponse.KodeTransaksi))
	
		for x, keranjangResponse := range keranjangResponses {
			keranjangResponses[x] = controller.FetchMetaData(request, keranjangResponse)
		}

		transaksiResponses[i].ListKeranjang = keranjangResponses
		transaksiResponses[i] = controller.FetchAdmin(request,transaksiResponses[i])
		transaksiResponses[i] = controller.FetchMember(request,transaksiResponses[i])
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transaksiResponses,
		Message: "Successfully fetch data",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
