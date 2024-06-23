package app

import (
	"net/http"
	"os"
	"toko-buku/controller"
	"toko-buku/exception"
	"toko-buku/middleware"

	"github.com/julienschmidt/httprouter"
)


func NewRouter(adminController controller.AdminController,memberController controller.MemberController, bukuController controller.BukuController, aksesorisController controller.AksesorisController, keranjangController controller.KeranjangController, transaksiController controller.TransaksiController, authController controller.AuthController) *httprouter.Router{
	apiRest := os.Getenv("API_REST")
	bookDirectoryPath := os.Getenv("BOOK_DIRECTORY_PATH")
	accessoryDirectoryPath := os.Getenv("ACCSESSORY_DIRECTORY_PATH")
	
	router := httprouter.New()

	//admin
	router.GET(apiRest+"/admins", middleware.CorsMiddleware(adminController.FindAll))
	router.GET(apiRest+"/admins/:adminId", middleware.CorsMiddleware(adminController.FindById))
	router.POST(apiRest+"/admins", middleware.CorsMiddleware(adminController.Create))
	router.PUT(apiRest+"/admins/:adminId", middleware.CorsMiddleware(adminController.Update))
	router.DELETE(apiRest+"/admins/:adminId", middleware.CorsMiddleware(adminController.Delete))
	
	//members
	router.GET(apiRest+"/members", middleware.CorsMiddleware(memberController.FindAll))
	router.GET(apiRest+"/members/:memberId", middleware.CorsMiddleware(memberController.FindById))
	router.POST(apiRest+"/members", middleware.CorsMiddleware(memberController.Create))
	router.PUT(apiRest+"/members/:memberId", middleware.CorsMiddleware(memberController.Update))
	router.DELETE(apiRest+"/members/:memberId", middleware.CorsMiddleware(memberController.Delete))
	

	//buku
	router.GET(apiRest+"/buku", middleware.CorsMiddleware(bukuController.FindAll))
	router.GET(apiRest+"/buku/:bukuId", middleware.CorsMiddleware(bukuController.FindById))
	router.POST(apiRest+"/buku", middleware.CorsMiddleware(bukuController.Create))
	router.PUT(apiRest+"/buku/:bukuId", middleware.CorsMiddleware(bukuController.Update))
	router.PUT(apiRest+"/buku/:bukuId/stok", middleware.CorsMiddleware(bukuController.UpdateStok))
	router.DELETE(apiRest+"/buku/:bukuId", middleware.CorsMiddleware(bukuController.Delete))

	//aksesoris
	router.GET(apiRest+"/aksesoris", middleware.CorsMiddleware(aksesorisController.FindAll))
	router.GET(apiRest+"/aksesoris/:aksesorisId", middleware.CorsMiddleware(aksesorisController.FindById))
	router.POST(apiRest+"/aksesoris", middleware.CorsMiddleware(aksesorisController.Create))
	router.PUT(apiRest+"/aksesoris/:aksesorisId", middleware.CorsMiddleware(aksesorisController.Update))
	router.PUT(apiRest+"/aksesoris/:aksesorisId/stok", middleware.CorsMiddleware(aksesorisController.UpdateStok))
	router.DELETE(apiRest+"/aksesoris/:aksesorisId", middleware.CorsMiddleware(aksesorisController.Delete))

	//keranjang
	router.GET(apiRest+"/keranjang", middleware.CorsMiddleware(keranjangController.FindAll))
	router.GET(apiRest+"/keranjang/:keranjangId", middleware.CorsMiddleware(keranjangController.FindById))
	router.POST(apiRest+"/keranjang", middleware.CorsMiddleware(keranjangController.Create))
	router.PUT(apiRest+"/keranjang/:keranjangId/total", middleware.CorsMiddleware(keranjangController.UpdateTotalBarang))
	router.DELETE(apiRest+"/keranjang/:keranjangId", middleware.CorsMiddleware(keranjangController.Delete))

	//transaksi
	router.GET(apiRest+"/transaksi", middleware.CorsMiddleware(transaksiController.FindAll))
	router.GET(apiRest+"/transaksi/:transaksiId", middleware.CorsMiddleware(transaksiController.FindById))
	router.POST(apiRest+"/transaksi", middleware.CorsMiddleware(transaksiController.Create))

	//auth
	router.POST(apiRest+"/authentication", middleware.CorsMiddleware(authController.Authentication))

	router.ServeFiles("/" + bookDirectoryPath +  "/*filepath", http.Dir(bookDirectoryPath))
	router.ServeFiles("/" + accessoryDirectoryPath +  "/*filepath", http.Dir(accessoryDirectoryPath))

	router.PanicHandler = exception.ErrorHandler

	return router
}