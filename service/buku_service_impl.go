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

type BukuServiceImpl struct {
	BukuRepository repository.BukuRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewBukuService(bukuRepository repository.BukuRepository, DB *sql.DB, validate *validator.Validate) BukuService {
	return &BukuServiceImpl{
		BukuRepository: bukuRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *BukuServiceImpl) Create(ctx context.Context, request web.BukuCreateRequest) web.BukuResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Rilis, _ := time.Parse("2006-01-02", request.Rilis)

	buku := domain.Buku{
		Id:        request.Id,
		Name:      request.Name,
		Penerbit:  request.Penerbit,
		Image:     request.Image,
		Pengarang: request.Pengarang,
		Genre:     request.Genre,
		Rilis:     Rilis,
		Harga:     request.Harga,
	}

	buku = service.BukuRepository.Save(ctx, tx, buku)
	return helper.ToBukuResponse(buku)
}

func (service *BukuServiceImpl) Update(ctx context.Context, request web.BukuUpdateRequest, bukuId string) web.BukuResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	Rilis, _ := time.Parse("2006-01-02", request.Rilis)

	buku, err := service.BukuRepository.FindById(ctx, tx, bukuId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	buku.Name = request.Name
	buku.Penerbit = request.Penerbit
	buku.Image = request.Image
	buku.Pengarang = request.Pengarang
	buku.Genre = request.Genre
	buku.Rilis = Rilis
	buku.Harga = request.Harga

	buku = service.BukuRepository.Update(ctx, tx, buku, bukuId)
	return helper.ToBukuResponse(buku)
}

func (service *BukuServiceImpl) Delete(ctx context.Context, bukuId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.BukuRepository.FindById(ctx, tx, bukuId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.BukuRepository.Delete(ctx, tx, bukuId)
}

func (service *BukuServiceImpl) FindById(ctx context.Context, bukuId string) web.BukuResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	buku, err := service.BukuRepository.FindById(ctx, tx, bukuId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToBukuResponse(buku)
}

func (service *BukuServiceImpl) FindAll(ctx context.Context) []web.BukuResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	bukus := service.BukuRepository.FindAll(ctx, tx)
	return helper.ToBukuResponses(bukus)
}
