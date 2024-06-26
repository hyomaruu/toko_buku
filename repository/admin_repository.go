package repository

import (
	"context"
	"database/sql"
	"toko-buku/model/domain"
	"toko-buku/model/web"
)

type AdminRepository interface {
	Save(ctx context.Context, tx *sql.Tx, admin domain.Admin) domain.Admin
	Update(ctx context.Context,tx *sql.Tx, admin domain.Admin) domain.Admin
	Delete(ctx context.Context, tx *sql.Tx, admin domain.Admin)
	FindById(ctx context.Context, tx *sql.Tx,  adminId int) (domain.Admin, error)
	FindAll(ctx context.Context, tx *sql.Tx, request web.FetchRequest) []domain.Admin
	RoleCounter(ctx context.Context, tx *sql.Tx) []domain.AdminCounter

}