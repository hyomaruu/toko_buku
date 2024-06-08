package repository

import (
	"context"
	"database/sql"
	"toko-buku/model/domain"
)

type BukuRepository interface {
	Save(ctx context.Context, tx *sql.Tx, buku domain.Buku) domain.Buku
	Update(ctx context.Context, tx *sql.Tx, buku domain.Buku) domain.Buku
	Delete(ctx context.Context, tx *sql.Tx, buku domain.Buku)
	FindById(ctx context.Context,tx *sql.Tx,bukuId string) (domain.Buku, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Buku
}