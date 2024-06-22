package repository

import (
	"context"
	"database/sql"
	"errors"
	"toko-buku/helper"
	"toko-buku/model/domain"
)

type AksesorisRepsitoryImpl struct {
}

func NewAksesorisRepository() AksesorisRepsitory {
	return &AksesorisRepsitoryImpl{}
}

func (repository *AksesorisRepsitoryImpl) Save(ctx context.Context, tx *sql.Tx, aksesoris domain.Aksesoris) domain.Aksesoris {
	SQL := "insert into aksesoris(kode, nama_barang, deskripsi, gambar, harga) values(?,?,?,?,?)"

	_, err := tx.ExecContext(ctx, SQL, aksesoris.Id, aksesoris.Name, aksesoris.Deskripsi, aksesoris.Image ,aksesoris.Harga)
	helper.PanicIfError(err)

	return aksesoris
}

func (repository *AksesorisRepsitoryImpl) Update(ctx context.Context, tx *sql.Tx, aksesoris domain.Aksesoris, aksesorisId string) domain.Aksesoris {
	SQL := "update aksesoris set nama_barang = ?, deskripsi = ?, harga = ?, gambar = ? where kode = ?"
	_, err := tx.ExecContext(ctx, SQL, aksesoris.Name, aksesoris.Deskripsi ,aksesoris.Harga, aksesoris.Image, aksesorisId)
	helper.PanicIfError(err)

	return aksesoris
}

func (repository *AksesorisRepsitoryImpl) UpdateStok(ctx context.Context, tx *sql.Tx, aksesoris domain.Aksesoris, aksesorisId string) domain.Aksesoris {
	SQL := "update aksesoris set stok = ? where kode = ?"
	_, err := tx.ExecContext(ctx, SQL, aksesoris.Stok, aksesorisId)
	helper.PanicIfError(err)

	return aksesoris
}

func (repository *AksesorisRepsitoryImpl) ReduceStok(ctx context.Context, tx *sql.Tx, aksesoris domain.Aksesoris, aksesorisId string) domain.Aksesoris {
	SQL := "update aksesoris set stok = - ? where kode = ?"
	_, err := tx.ExecContext(ctx, SQL, aksesoris.Stok, aksesorisId)
	helper.PanicIfError(err)

	return aksesoris
}

func (repository *AksesorisRepsitoryImpl) Delete(ctx context.Context, tx *sql.Tx, aksesorisId string) {
	SQL := "delete from aksesoris where kode = ?"
	_, err := tx.ExecContext(ctx, SQL, aksesorisId)
	helper.PanicIfError(err)
}

func (repository *AksesorisRepsitoryImpl) FindById(ctx context.Context, tx *sql.Tx, aksesorisId string) (domain.Aksesoris, error) {
	SQL := "select kode, nama_barang, deskripsi, harga, gambar, stok from aksesoris where kode = ?"
	rows, err := tx.QueryContext(ctx, SQL, aksesorisId)
	helper.PanicIfError(err)
	defer rows.Close()

	aksesoris := domain.Aksesoris{}
	if rows.Next() {
		err := rows.Scan(&aksesoris.Id, &aksesoris.Name, &aksesoris.Deskripsi, &aksesoris.Harga, &aksesoris.Image, &aksesoris.Stok)
		helper.PanicIfError(err)
		return aksesoris, nil
	} else {
		return aksesoris, errors.New("aksesoris tidak ditemukan")
	}
}

func (repository *AksesorisRepsitoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Aksesoris {
	SQL := "select kode, nama_barang, deskripsi, harga, gambar, stok from aksesoris"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var aksesoriss []domain.Aksesoris
	for rows.Next() {
		aksesoris := domain.Aksesoris{}
		err := rows.Scan(&aksesoris.Id, &aksesoris.Name, &aksesoris.Deskripsi, &aksesoris.Harga,  &aksesoris.Image, &aksesoris.Stok)
		helper.PanicIfError(err)
		aksesoriss = append(aksesoriss, aksesoris)
	}
	return aksesoriss
}
