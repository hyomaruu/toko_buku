package service

import (
	"context"
	"toko-buku/model/web"
)

type TransaksiService interface {
	Create(ctx context.Context, request web.TransaksiCreateRequest) web.TransaksiResponse
	FindById(ctx context.Context, keranjangId int) web.TransaksiResponse
	FindAll(ctx context.Context) []web.TransaksiResponse
}