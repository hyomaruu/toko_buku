package repository

import (
	"context"
	"database/sql"
	"toko-buku/model/domain"
)

type BukuRepository interface {
	Save(ctx context.Context, tx *sql.Tx, buku domain.Buku) domain.Buku
	Update(ctx context.Context, tx *sql.Tx, buku domain.Buku, bukuId string) domain.Buku
	UpdateStok(ctx context.Context, tx *sql.Tx, buku domain.Buku, bukuId string) domain.Buku
	Delete(ctx context.Context, tx *sql.Tx, bukuId string)
	FindById(ctx context.Context,tx *sql.Tx, bukuId string) (domain.Buku, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Buku
}