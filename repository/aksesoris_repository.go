package repository

import (
	"context"
	"database/sql"
	"toko-buku/model/domain"
	"toko-buku/model/web"
)

type AksesorisRepsitory interface {
	Save(ctx context.Context,tx *sql.Tx, aksesoris domain.Aksesoris) domain.Aksesoris
	Update(ctx context.Context,tx *sql.Tx,aksesoris domain.Aksesoris, aksesorisId string) domain.Aksesoris
	UpdateStok(ctx context.Context,tx *sql.Tx,aksesoris domain.Aksesoris, aksesorisId string) domain.Aksesoris
	ReduceStok(ctx context.Context,tx *sql.Tx,aksesoris domain.Aksesoris, aksesorisId string) domain.Aksesoris
	Delete(ctx context.Context, tx *sql.Tx, aksesorisId string)
	FindById(ctx context.Context,tx *sql.Tx, aksesorisId string)(domain.Aksesoris, error)
	FindAll(ctx context.Context,tx *sql.Tx, request web.FetchRequest) []domain.Aksesoris
}