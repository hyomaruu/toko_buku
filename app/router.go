package app

import (
	"net/http"
	"os"
	"toko-buku/controller"
	"toko-buku/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(adminController controller.AdminController,memberController controller.MemberController, bukuController controller.BukuController, aksesorisController controller.AksesorisController, keranjangController controller.KeranjangController, transaksiController controller.TransaksiController) *httprouter.Router{
	apiRest := os.Getenv("API_REST")
	bookDirectoryPath := os.Getenv("BOOK_DIRECTORY_PATH")
	accessoryDirectoryPath := os.Getenv("ACCSESSORY_DIRECTORY_PATH")
	
	router := httprouter.New()

	//admin
	router.GET(apiRest+"/admins", adminController.FindAll)
	router.GET(apiRest+"/admins/:adminId", adminController.FindById)
	router.POST(apiRest+"/admins", adminController.Create)
	router.PUT(apiRest+"/admins/:adminId", adminController.Update)
	router.DELETE(apiRest+"/admins/:adminId", adminController.Delete)
	
	//members
	router.GET(apiRest+"/members",memberController.FindAll)
	router.GET(apiRest+"/members/:memberId",memberController.FindById)
	router.POST(apiRest+"/members",memberController.Create)
	router.PUT(apiRest+"/members/:memberId",memberController.Update)
	router.DELETE(apiRest+"/members/:memberId",memberController.Delete)
	router.PanicHandler = exception.ErrorHandler

	//buku
	router.GET(apiRest+"/buku",bukuController.FindAll)
	router.GET(apiRest+"/buku/:bukuId",bukuController.FindById)
	router.POST(apiRest+"/buku",bukuController.Create)
	router.PUT(apiRest+"/buku/:bukuId",bukuController.Update)
	router.PUT(apiRest+"/buku/:bukuId/stok",bukuController.UpdateStok)
	router.DELETE(apiRest+"/buku/:bukuId",bukuController.Delete)

	//aksesoris
	router.GET(apiRest+"/aksesoris",aksesorisController.FindAll)
	router.GET(apiRest+"/aksesoris/:aksesorisId",aksesorisController.FindById)
	router.POST(apiRest+"/aksesoris",aksesorisController.Create)
	router.PUT(apiRest+"/aksesoris/:aksesorisId",aksesorisController.Update)
	router.PUT(apiRest+"/aksesoris/:aksesorisId/stok",aksesorisController.UpdateStok)
	router.DELETE(apiRest+"/aksesoris/:aksesorisId",aksesorisController.Delete)

	//keranjang
	router.GET(apiRest+"/keranjang",keranjangController.FindAll)
	router.GET(apiRest+"/keranjang/:keranjangId",keranjangController.FindById)
	router.POST(apiRest+"/keranjang",keranjangController.Create)
	router.PUT(apiRest+"/keranjang/:keranjangId/total",keranjangController.UpdateTotalBarang)
	router.DELETE(apiRest+"/keranjang/:keranjangId",keranjangController.Delete)

	//transaksi
	router.GET(apiRest+"/transaksi",transaksiController.FindAll)
	router.GET(apiRest+"/transaksi/:transaksiId",transaksiController.FindById)
	router.POST(apiRest+"/transaksi",transaksiController.Create)


	router.ServeFiles("/" + bookDirectoryPath +  "/*filepath", http.Dir(bookDirectoryPath))
	router.ServeFiles("/" + accessoryDirectoryPath +  "/*filepath", http.Dir(accessoryDirectoryPath))

	return router
}