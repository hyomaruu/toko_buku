package main

import (
	"fmt"
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

	stockRepository := repository.NewStokRepository()
	stockService := service.NewStockService(stockRepository,db,validate)
	stockController := controller.NewStockController(stockService)

	router := app.NewRouter(adminController,memberController, bukuController,aksesorisController,stockController)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}
	address := "localhost:3000"
	fmt.Printf("Server started at %s\n", address)

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}


