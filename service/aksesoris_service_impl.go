package service

import (
	"context"
	"database/sql"
	"toko-buku/exception"
	"toko-buku/helper"
	"toko-buku/model/domain"
	"toko-buku/model/web"
	"toko-buku/repository"

	"github.com/go-playground/validator/v10"
)

type AksesorisServiceImpl struct {
	AksesorisRepository repository.AksesorisRepsitory
	DB *sql.DB
	Validate *validator.Validate
}

func NewAksesorisService(aksesorisRepository repository.AksesorisRepsitory,DB *sql.DB, validate *validator.Validate) AksesorisService{
	return &AksesorisServiceImpl{
		AksesorisRepository: aksesorisRepository,
		DB: DB,
		Validate: validate,
	}
}

func(service *AksesorisServiceImpl) Create(ctx context.Context, request web.AksesorisCreateRequest) web.AksesorisResponse{
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	aksesoris := domain.Aksesoris{
		Id: request.Id,
		Name: request.Name,
		Harga: request.Harga,
		Image: request.Image,
		Deskripsi: request.Deskripsi,
		Stok: 0,
	}

	aksesoris = service.AksesorisRepository.Save(ctx,tx,aksesoris)
	return helper.ToAksesorisResponse(aksesoris)
}

func(service *AksesorisServiceImpl) Update(ctx context.Context, request web.AksesorisUpdateRequest, aksesorisId string) web.AksesorisResponse{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	aksesoris, err := service.AksesorisRepository.FindById(ctx,tx, aksesorisId)
	if err!=nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	aksesoris.Name = request.Name
	aksesoris.Harga = request.Harga
	aksesoris.Image = request.Image
	aksesoris.Deskripsi = request.Deskripsi

	aksesoris = service.AksesorisRepository.Update(ctx,tx, aksesoris, aksesorisId)
	return helper.ToAksesorisResponse(aksesoris)
}

func(service *AksesorisServiceImpl) UpdateStok(ctx context.Context, request web.AksesorisUpdateStockRequest, aksesorisId string) web.AksesorisResponse{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	aksesoris, err := service.AksesorisRepository.FindById(ctx,tx, aksesorisId)
	if err!=nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	aksesoris.Stok = request.Stok

	aksesoris = service.AksesorisRepository.UpdateStok(ctx,tx, aksesoris, aksesorisId)
	return helper.ToAksesorisResponse(aksesoris)
}

func(service *AksesorisServiceImpl) ReduceStok(ctx context.Context, qty int, aksesorisId string) web.AksesorisResponse{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	aksesoris, err := service.AksesorisRepository.FindById(ctx,tx, aksesorisId)
	if err!=nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	aksesoris.Stok = qty

	aksesoris = service.AksesorisRepository.ReduceStok(ctx,tx, aksesoris, aksesorisId)
	return helper.ToAksesorisResponse(aksesoris)
}

func(service *AksesorisServiceImpl) Delete(ctx context.Context, aksesorisId string){
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.AksesorisRepository.FindById(ctx,tx, aksesorisId)
	if err!=nil{
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.AksesorisRepository.Delete(ctx,tx, aksesorisId)
}

func(service *AksesorisServiceImpl) FindById(ctx context.Context, aksesorisId string)web.AksesorisResponse{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	aksesoris, err := service.AksesorisRepository.FindById(ctx,tx,aksesorisId)
	if err!=nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToAksesorisResponse(aksesoris)
}

func(service *AksesorisServiceImpl) FindAll(ctx context.Context) []web.AksesorisResponse{
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	aksesoriss := service.AksesorisRepository.FindAll(ctx,tx)
	return helper.ToAksesorisResponses(aksesoriss)
}