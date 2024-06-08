package service

import (
	"context"
	"toko-buku/model/web"
)

type AksesorisService interface {
	Create(ctx context.Context, request web.AksesorisCreateRequest) web.AksesorisResponse
	Update(ctx context.Context, request web.AksesorisUpdateRequest)web.AksesorisResponse
	Delete(ctx context.Context, aksesorisId string)
	FindById(ctx context.Context, aksesorisId string) web.AksesorisResponse
	FindAll(ctx context.Context) []web.AksesorisResponse
}