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
	}

	aksesoris = service.AksesorisRepository.Save(ctx,tx,aksesoris)
	return helper.ToAksesorisResponse(aksesoris)
}

func(service *AksesorisServiceImpl) Update(ctx context.Context, request web.AksesorisUpdateRequest) web.AksesorisResponse{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	aksesoris, err := service.AksesorisRepository.FindById(ctx,tx,request.Id)
	if err!=nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	aksesoris.Name = request.Name
	aksesoris.Harga = request.Harga

	aksesoris = service.AksesorisRepository.Update(ctx,tx,aksesoris)
	return helper.ToAksesorisResponse(aksesoris)
}

func(service *AksesorisServiceImpl) Delete(ctx context.Context, AksesorisId string){
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	aksesoris, err := service.AksesorisRepository.FindById(ctx,tx,AksesorisId)
	if err!=nil{
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.AksesorisRepository.Delete(ctx,tx,aksesoris)
}

func(service *AksesorisServiceImpl) FindById(ctx context.Context, AksesorisId string)web.AksesorisResponse{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	aksesoris,err := service.AksesorisRepository.FindById(ctx,tx,AksesorisId)
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