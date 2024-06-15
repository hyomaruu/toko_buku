package service

import (
	"context"
	"toko-buku/model/web"
)

type BukuService interface {
	Create(ctx context.Context, request web.BukuCreateRequest) web.BukuResponse
	Update(ctx context.Context,request web.BukuUpdateRequest, bukuId string) web.BukuResponse
	UpdateStok(ctx context.Context,request web.BukuUpdateStockRequest, bukuId string) web.BukuResponse
	Delete(ctx context.Context, bukuId string)
	FindById(ctx context.Context, bukuId string) web.BukuResponse
	FindAll(ctx context.Context) []web.BukuResponse
}