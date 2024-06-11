package repository

import (
	"context"
	"database/sql"
	"errors"
	"toko-buku/helper"
	"toko-buku/model/domain"
)

type StockRepositoryImpl struct {
}

func NewStokRepository() StockRepository {
	return &StockRepositoryImpl{}
}
func (repository *StockRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, stok domain.Stok) domain.Stok {
	SQL := "insert into stok_barang(kode_barang,stok) values (?,?)"

	_, err := tx.ExecContext(ctx, SQL, stok.Id, stok.Stock)
	helper.PanicIfError(err)

	return stok
}

func (repository *StockRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, stok domain.Stok, stokId string) domain.Stok {
	SQL := "update stok_barang set stok = ? where kode_barang = ?"
	_, err := tx.ExecContext(ctx, SQL, stok.Stock, stok.Id)
	helper.PanicIfError(err)

	return stok
}

func (repository *StockRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, stokId string) {
	SQL := "delete from stok_barang where kode_barang = ?"
	_, err := tx.ExecContext(ctx, SQL, stokId)
	helper.PanicIfError(err)
}

func (repository *StockRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, stokId string) (domain.Stok, error) {
	SQL := "select * from stok_barang where kode_barang = ?"
	rows, err := tx.QueryContext(ctx, SQL, stokId)
	helper.PanicIfError(err)
	defer rows.Close()

	stok := domain.Stok{}
	if rows.Next() {
		err := rows.Scan(&stok.Id, &stok.Stock)
		helper.PanicIfError(err)
		return stok, nil
	} else {
		return stok, errors.New("kode barang tidak ditemukan")
	}
}

func (repository *StockRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Stok {
	SQL := "select * from stok_barang"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var stoks []domain.Stok
	for rows.Next() {
		stok := domain.Stok{}
		err := rows.Scan(&stok.Id, &stok.Stock)
		helper.PanicIfError(err)
		stoks = append(stoks, stok)
	}
	return stoks
}
