package repository

import (
	"context"
	"database/sql"
	"toko-buku/model/domain"
)

type KeranjangRepository interface {
	Create(ctx context.Context, tx *sql.Tx, keranjang domain.Keranjang) domain.Keranjang

	UpdateTotalBarang(ctx context.Context, tx *sql.Tx, keranjang domain.Keranjang, keranjangId int) domain.Keranjang

	Delete(ctx context.Context, tx *sql.Tx, keranjangId int)

	FindById(ctx context.Context, tx *sql.Tx, keranjangId int) (domain.Keranjang, error)

	FindAll(ctx context.Context, tx *sql.Tx) []domain.Keranjang
}