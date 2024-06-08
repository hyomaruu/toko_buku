package app

import (
	"toko-buku/controller"
	"toko-buku/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(adminController controller.AdminController,memberController controller.MemberController, bukuController controller.BukuController, aksesorisController controller.AksesorisController) *httprouter.Router{
	router := httprouter.New()

	//admin
	router.GET("/api/admins", adminController.FindAll)
	router.GET("/api/admins/:adminId", adminController.FindById)
	router.POST("/api/admins", adminController.Create)
	router.PUT("/api/admins/:adminId", adminController.Update)
	router.DELETE("/api/admins/:adminId", adminController.Delete)
	
	//members
	router.GET("/api/members",memberController.FindAll)
	router.GET("/api/members/:memberId",memberController.FindById)
	router.POST("/api/members",memberController.Create)
	router.PUT("/api/members/:memberId",memberController.Update)
	router.DELETE("/api/members/:memberId",memberController.Delete)
	router.PanicHandler = exception.ErrorHandler

	//buku
	router.GET("/api/buku",bukuController.FindAll)
	router.GET("/api/buku/:bukuId",bukuController.FindById)
	router.POST("/api/buku",bukuController.Create)
	router.PUT("/api/buku/:bukuId",bukuController.Update)
	router.DELETE("/api/buku/:bukuId",bukuController.Delete)

	//aksesoris
	router.GET("/api/aksesoris",aksesorisController.FindAll)
	router.GET("/api/aksesoris/:aksesorisId",aksesorisController.FindById)
	router.POST("/api/aksesoris",aksesorisController.Create)
	router.PUT("/api/aksesoris/:aksesorisId",aksesorisController.Update)
	router.DELETE("/api/aksesoris/:aksesorisId",aksesorisController.Delete)

	
	return router
}