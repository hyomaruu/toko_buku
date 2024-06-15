package service

import (
	"context"
	"toko-buku/model/web"
)

type KeranjangService interface {
	Create(ctx context.Context, request web.KeranjangCreateRequest) web.KeranjangResponse
	UpdateTotalBarang(ctx context.Context,request web.KeranjangUpdateStockRequest, keranjangId int) web.KeranjangResponse
	Delete(ctx context.Context, keranjangId int)
	FindById(ctx context.Context, keranjangId int) web.KeranjangResponse
	FindAll(ctx context.Context) []web.KeranjangResponse
}