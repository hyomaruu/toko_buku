package service

import (
	"context"
	"database/sql"
	"time"
	"toko-buku/exception"
	"toko-buku/helper"
	"toko-buku/model/domain"
	"toko-buku/model/web"
	"toko-buku/repository"

	"github.com/go-playground/validator/v10"
)

type TransaksiServiceImpl struct {
	TransaksiRepository repository.TransaksiRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewTransaksiService(transaksiRepository repository.TransaksiRepository, DB *sql.DB, validate *validator.Validate) TransaksiService {
	return &TransaksiServiceImpl{
		TransaksiRepository: transaksiRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *TransaksiServiceImpl) Create(ctx context.Context, request web.TransaksiCreateRequest) web.TransaksiResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Tanggal, _ := time.Parse("2006-01-02", request.Tanggal)

	transaksi := domain.Transaksi{
		NomorMember: request.NomorMember,
		IdAdmin:     request.IdAdmin,
		Tanggal:     Tanggal,
		TotalHarga:  request.TotalHarga,
	}

	transaksi = service.TransaksiRepository.Save(ctx, tx, transaksi)

	return helper.ToTransaksiResponse(transaksi)
}


func (service *TransaksiServiceImpl) FindById(ctx context.Context, transaksiId int) web.TransaksiResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaksi, err := service.TransaksiRepository.FindById(ctx, tx, transaksiId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToTransaksiResponse(transaksi)
}

func (service *TransaksiServiceImpl) FindAll(ctx context.Context) []web.TransaksiResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaksi := service.TransaksiRepository.FindAll(ctx, tx)

	return helper.ToTransaksiResponses(transaksi)
}
