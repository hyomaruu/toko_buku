package repository

import (
	"database/sql"
	"toko-buku/model/domain"
	"context"
)

type AdminRepository interface {
	Save(ctx context.Context, tx *sql.Tx, admin domain.Admin) domain.Admin

	Update(ctx context.Context,tx *sql.Tx, admin domain.Admin) domain.Admin

	Delete(ctx context.Context, tx *sql.Tx, admin domain.Admin)

	FindById(ctx context.Context, tx *sql.Tx,  adminId int) (domain.Admin, error)

	FindAll(ctx context.Context, tx *sql.Tx) []domain.Admin
}