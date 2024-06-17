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

type KeranjangServiceImpl struct {
	KeranjangRepository repository.KeranjangRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewKeranjangService(keranjangRepository repository.KeranjangRepository, DB *sql.DB, validate *validator.Validate) KeranjangService {
	return &KeranjangServiceImpl{
		KeranjangRepository: keranjangRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *KeranjangServiceImpl) Create(ctx context.Context, request web.KeranjangCreateRequest) web.KeranjangResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	keranjang := domain.Keranjang{
		KodeBarang:      request.KodeBarang,
		TipeBarang:      request.TipeBarang,
		TotalBarang:     request.TotalBarang,
		Harga:      request.Harga,
	}

	keranjang = service.KeranjangRepository.Create(ctx, tx, keranjang)
	return helper.ToKeranjangResponse(keranjang)
}

func (service *KeranjangServiceImpl) UpdateTotalBarang(ctx context.Context, request web.KeranjangUpdateStockRequest, keranjangId int) web.KeranjangResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	keranjang, err := service.KeranjangRepository.FindById(ctx, tx, keranjangId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	keranjang.TotalBarang = request.TotalBarang
	keranjang = service.KeranjangRepository.UpdateTotalBarang(ctx, tx, keranjang, keranjangId)

	return helper.ToKeranjangResponse(keranjang)
}

func (service *KeranjangServiceImpl) UpdateKodeTransaksi(ctx context.Context, keranjangId int, kodeTransaksi int) web.KeranjangResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	keranjang, err := service.KeranjangRepository.FindById(ctx, tx, keranjangId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	keranjang = service.KeranjangRepository.UpdateKodeTransaksi(ctx, tx, keranjangId, kodeTransaksi)

	return helper.ToKeranjangResponse(keranjang)
}


func (service *KeranjangServiceImpl) Delete(ctx context.Context, keranjangId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.KeranjangRepository.FindById(ctx, tx, keranjangId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.KeranjangRepository.Delete(ctx, tx, keranjangId)
}

func (service *KeranjangServiceImpl) FindById(ctx context.Context, keranjangId int) web.KeranjangResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	keranjang, err := service.KeranjangRepository.FindById(ctx, tx, keranjangId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToKeranjangResponse(keranjang)
}

func (service *KeranjangServiceImpl) FindAll(ctx context.Context) []web.KeranjangResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	keranjangs := service.KeranjangRepository.FindAll(ctx, tx)
	return helper.ToKeranjangResponses(keranjangs)
}

func (service *KeranjangServiceImpl) FindByKodeTransaksi(ctx context.Context, kodeTransaksi int) []web.KeranjangResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	keranjangs := service.KeranjangRepository.FindByKodeTransaksi(ctx, tx, kodeTransaksi)
	return helper.ToKeranjangResponses(keranjangs)
}
