package service

import (
	"context"
	"toko-buku/model/web"
)

type KeranjangService interface {
	Create(ctx context.Context, request web.KeranjangCreateRequest) web.KeranjangResponse
	UpdateTotalBarang(ctx context.Context,request web.KeranjangUpdateStockRequest, keranjangId int) web.KeranjangResponse
	UpdateKodeTransaksi(ctx context.Context, keranjangId int, kodeTransaksi int) web.KeranjangResponse
	Delete(ctx context.Context, keranjangId int)
	FindByKodeTransaksi(ctx context.Context, kodeTransaksi int) []web.KeranjangResponse
	FindById(ctx context.Context, keranjangId int) web.KeranjangResponse
	FindAll(ctx context.Context) []web.KeranjangResponse
}