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

type StockServiceImpl struct {
	StockRepository repository.StockRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewStockService(stokRepository repository.StockRepository, DB *sql.DB, validate *validator.Validate) StockService {
	return &StockServiceImpl{
		StockRepository: stokRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *StockServiceImpl) Create(ctx context.Context, request web.StockCreateRequest) web.StockResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	stok := domain.Stok{
		Id:    request.Id,
		Stock: request.Stock,
	}
	stok = service.StockRepository.Save(ctx, tx, stok)
	return helper.ToStockResponse(stok)
}

func (service *StockServiceImpl) Update(ctx context.Context, request web.StockUpdateRequest, stokId string) web.StockResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	stok, err := service.StockRepository.FindById(ctx, tx, stokId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	stok.Stock = request.Stock

	stok = service.StockRepository.Update(ctx, tx, stok, stokId)
	return helper.ToStockResponse(stok)
}

func (service *StockServiceImpl) Delete(ctx context.Context,stokId string) {
	tx,err :=service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_,err = service.StockRepository.FindById(ctx,tx, stokId)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.StockRepository.Delete(ctx,tx,stokId)
}

func(service *StockServiceImpl) FindById(ctx context.Context,stokId string)web.StockResponse{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	stok,err :=service.StockRepository.FindById(ctx,tx,stokId)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToStockResponse(stok)
}

func(service *StockServiceImpl) FindAll(ctx context.Context)[]web.StockResponse{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	stoks := service.StockRepository.FindAll(ctx,tx)
	return helper.ToStockResponses(stoks)
}