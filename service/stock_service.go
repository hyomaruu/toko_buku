package service

import (
	"context"
	"toko-buku/model/web"
)

type StockService interface {
	Create(ctx context.Context, request web.StockCreateRequest) web.StockResponse
	Update(ctx context.Context, request web.StockUpdateRequest, stokId string) web.StockResponse
	Delete(ctx context.Context, stokId string)
	FindById(ctx context.Context, stokId string) web.StockResponse
	FindAll(ctx context.Context) []web.StockResponse
}
