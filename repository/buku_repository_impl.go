package repository

import (
	"context"
	"database/sql"
	"errors"
	"toko-buku/helper"
	"toko-buku/model/domain"
)

type BukuRepositoryImpl struct {
}

func NewBukuRepository() BukuRepository {
	return &BukuRepositoryImpl{}
}

func (repository *BukuRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, buku domain.Buku) domain.Buku {
	SQL := "insert into data_buku(kode_buku,nama_buku,penerbit,gambar,pengarang,genre,rilis,harga) values (?,?,?,?,?,?,?,?)"

	_, err := tx.ExecContext(ctx, SQL, buku.Id, buku.Name, buku.Penerbit, buku.Image, buku.Pengarang, buku.Genre,buku.Rilis, buku.Harga)
	helper.PanicIfError(err)

	return buku
}

func (repository *BukuRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, buku domain.Buku) domain.Buku {
	SQL := "update data_buku set nama_buku = ?,penerbit = ?,gambar = ?,pengarang = ?,genre =?,rilis = ?,harga = ? where kode_buku = ?"
	_, err := tx.ExecContext(ctx, SQL, buku.Name, buku.Penerbit, buku.Image, buku.Pengarang, buku.Genre, buku.Rilis, buku.Harga, buku.Id)
	helper.PanicIfError(err)

	return buku
}

func (repository *BukuRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, buku domain.Buku) {
	SQL := "delete from data_buku where kode_buku = ?"
	_, err := tx.ExecContext(ctx, SQL, buku.Id)
	helper.PanicIfError(err)

}

func (repository *BukuRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bukuId string) (domain.Buku, error) {
	SQL := "select kode_buku ,nama_buku ,penerbit ,gambar ,pengarang ,genre, rilis ,harga from data_buku where kode_buku = ?"
	rows, err := tx.QueryContext(ctx, SQL, bukuId)
	helper.PanicIfError(err)
	defer rows.Close()

	buku := domain.Buku{}
	if rows.Next() {
		err := rows.Scan(&buku.Id, &buku.Name, &buku.Penerbit, &buku.Image, &buku.Pengarang, &buku.Genre,&buku.Rilis, &buku.Harga)
		helper.PanicIfError(err)
		return buku, nil
	} else {
		return buku, errors.New("buku tidak ditemukan")
	}
}

func (repository *BukuRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Buku {
	SQL := "select kode_buku, nama_buku, penerbit, gambar, pengarang, genre, rilis, harga from data_buku"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	

	var bukus []domain.Buku
	for rows.Next() {
		buku := domain.Buku{}
		err := rows.Scan(&buku.Id, &buku.Name, &buku.Penerbit, &buku.Image, &buku.Pengarang, &buku.Genre,&buku.Rilis, &buku.Harga)
		helper.PanicIfError(err)
		bukus = append(bukus, buku)
	}
	return bukus
}
