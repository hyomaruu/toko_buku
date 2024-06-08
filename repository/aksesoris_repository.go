package repository

import (
	"context"
	"database/sql"
	"toko-buku/model/domain"
)

type AksesorisRepsitory interface {
	Save(ctx context.Context,tx *sql.Tx, aksesoris domain.Aksesoris) domain.Aksesoris
	Update(ctx context.Context,tx *sql.Tx,aksesoris domain.Aksesoris) domain.Aksesoris
	Delete(ctx context.Context, tx *sql.Tx,aksesoris domain.Aksesoris)
	FindById(ctx context.Context,tx *sql.Tx,aksesorisId string)(domain.Aksesoris, error)
	FindAll(ctx context.Context,tx *sql.Tx) []domain.Aksesoris
}