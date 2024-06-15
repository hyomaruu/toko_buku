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
	bukuRepository repository.BukuRepository
	aksesorisRepository repository.AksesorisRepsitory
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewKeranjangService(keranjangRepository repository.KeranjangRepository, DB *sql.DB, validate *validator.Validate) KeranjangService {
	return &KeranjangServiceImpl{
		KeranjangRepository: keranjangRepository,
		bukuRepository: repository.NewBukuRepository(),
		aksesorisRepository: repository.NewAksesorisRepository(),
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

	if (keranjang.TipeBarang == "AKSESORIS") {
		aksesoris, errAksesoris := service.aksesorisRepository.FindById(ctx, tx, keranjang.KodeBarang)
		if errAksesoris != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}

		keranjang.Data = helper.ToAksesorisResponse(aksesoris)
	}

	if (keranjang.TipeBarang == "BUKU") {
		buku, errBuku := service.bukuRepository.FindById(ctx, tx, keranjang.KodeBarang)
		if errBuku != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}

		keranjang.Data = helper.ToBukuResponse(buku)
	}

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

	if (keranjang.TipeBarang == "AKSESORIS") {
		aksesoris, errAksesoris := service.aksesorisRepository.FindById(ctx, tx, keranjang.KodeBarang)
		if errAksesoris != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}

		keranjang.Data = helper.ToAksesorisResponse(aksesoris)
	}

	if (keranjang.TipeBarang == "BUKU") {
		buku, errBuku := service.bukuRepository.FindById(ctx, tx, keranjang.KodeBarang)
		if errBuku != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}

		keranjang.Data = helper.ToBukuResponse(buku)
	}
	
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

	if (keranjang.TipeBarang == "AKSESORIS") {
		aksesoris, errAksesoris := service.aksesorisRepository.FindById(ctx, tx, keranjang.KodeBarang)
		if errAksesoris != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}

		keranjang.Data = helper.ToAksesorisResponse(aksesoris)

	}

	if (keranjang.TipeBarang == "BUKU") {
		buku, errBuku := service.bukuRepository.FindById(ctx, tx, keranjang.KodeBarang)
		if errBuku != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}

		keranjang.Data = helper.ToBukuResponse(buku)

	}

	return helper.ToKeranjangResponse(keranjang)
}

func (service *KeranjangServiceImpl) FindAll(ctx context.Context) []web.KeranjangResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	keranjangs := service.KeranjangRepository.FindAll(ctx, tx)

	for i, keranjang := range keranjangs {
		if (keranjang.TipeBarang == "AKSESORIS") {
			aksesoris, errAksesoris := service.aksesorisRepository.FindById(ctx, tx, keranjang.KodeBarang)
			if errAksesoris != nil {
				panic(exception.NewNotFoundError(err.Error()))
			}

			keranjangs[i].Data = helper.ToAksesorisResponse(aksesoris)
		}

		if (keranjang.TipeBarang == "BUKU") {
			buku, errBuku := service.bukuRepository.FindById(ctx, tx, keranjang.KodeBarang)
			if errBuku != nil {
				panic(exception.NewNotFoundError(err.Error()))
			}

			keranjangs[i].Data = helper.ToBukuResponse(buku)
		}
	}

	return helper.ToKeranjangResponses(keranjangs)
}
