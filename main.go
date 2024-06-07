package main

import (
	"net/http"
	"toko-buku/app"
	"toko-buku/controller"
	"toko-buku/helper"
	"toko-buku/repository"
	"toko-buku/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {

	validate := validator.New()
	db := app.NewDB()
	adminRepository := repository.NewAdminRepositry()
	adminService := service.NewAdminService(adminRepository, db, validate)
	adminController := controller.NewAdminController(adminService)
	
	memberRepository := repository.NewMemberRepository()
	memberService := service.NewMemberService(memberRepository,db,validate)
	memberController := controller.NewMemberController(memberService)

	bukuRepository := repository.NewBukuRepository()
	bukuService := service.NewBukuService(bukuRepository,db,validate)
	bukuController := controller.NewBukuController(bukuService)

	aksesorisRepository := repository.NewAksesorisRepository()
	aksesorisService := service.NewAksesorisService(aksesorisRepository,db,validate)
	aksesorisController := controller.NewAksesorisController(aksesorisService)

	router := app.NewRouter(adminController,memberController, bukuController,aksesorisController)
	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
