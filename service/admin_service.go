package service

import (
	"context"
	"toko-buku/model/web"
)

type AdminService interface {
	Create(ctx context.Context, request web.AdminCreateRequest) web.AdminResponse
	Update(ctx context.Context, request web.AdminUpdateRequest) web.AdminResponse
	Delete(ctx context.Context, adminId int)
	FindById(ctx context.Context, adminId int) web.AdminResponse
	FindAll(ctx context.Context) []web.AdminResponse
}