package main

import (
	"fmt"
	"net/http"
	"os"
	"toko-buku/app"
	"toko-buku/controller"
	"toko-buku/helper"
	"toko-buku/repository"
	"toko-buku/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	helper.PanicIfError(envErr)

	host := os.Getenv("HOST")

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

	keranjangRepository := repository.NewKeranjangRepository()
	keranjangService := service.NewKeranjangService(keranjangRepository,db,validate)
	keranjangController := controller.NewKeranjangController(keranjangService, bukuService, aksesorisService)

	transaksiRepository := repository.NewTransaksiRepository()
	transaksiService := service.NewTransaksiService(transaksiRepository,db,validate)
	transaksiController := controller.NewTransaksiController(transaksiService, keranjangService, bukuService, aksesorisService)


	authRepository := repository.NewAuthRepositry()
	authService := service.NewAuthService(authRepository, db, validate)
	authController := controller.NewAuthController(authService)

	router := app.NewRouter(adminController, memberController, bukuController, aksesorisController, keranjangController, transaksiController, authController)

	server := http.Server{
		Addr: host,
		Handler: router,
	}
	fmt.Printf("Server started at %s\n", server.Addr)

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}


