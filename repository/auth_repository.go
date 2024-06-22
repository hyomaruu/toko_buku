package repository

import (
	"context"
	"database/sql"
	"toko-buku/model/domain"
	"toko-buku/model/web"
)

type AuthRepository interface {
	Authentication(ctx context.Context, tx *sql.Tx,  auth web.AuthRequest) (domain.Admin, error)
}