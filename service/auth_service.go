package service

import (
	"context"
	"toko-buku/model/web"
)

type AuthService interface {
	Authentication(ctx context.Context, request web.AuthRequest) web.AdminResponse
}