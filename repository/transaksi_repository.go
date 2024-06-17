package repository

import (
	"context"
	"database/sql"
	"toko-buku/model/domain"
)

type TransaksiRepository interface {
	Save(ctx context.Context, tx *sql.Tx, transaksi domain.Transaksi) domain.Transaksi

	FindById(ctx context.Context, tx *sql.Tx, transaksiId int) (domain.Transaksi, error)

	FindAll(ctx context.Context, tx *sql.Tx) []domain.Transaksi
}