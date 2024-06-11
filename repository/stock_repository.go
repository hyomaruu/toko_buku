package repository

import (
	"context"
	"database/sql"
	"toko-buku/model/domain"
)

type StockRepository interface {
	Save(ctx context.Context, tx *sql.Tx, stok domain.Stok) domain.Stok
	Update(ctx context.Context, tx *sql.Tx, stok domain.Stok,stokId string) domain.Stok
	Delete(ctx context.Context, tx *sql.Tx, stokId string)
	FindById(ctx context.Context, tx *sql.Tx, stokId string) (domain.Stok, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Stok
}
