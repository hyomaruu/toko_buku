package repository

import (
	"context"
	"database/sql"
	"errors"
	"toko-buku/helper"
	"toko-buku/model/domain"
	"toko-buku/model/web"
)

type BukuRepositoryImpl struct {
}

func NewBukuRepository() BukuRepository {
	return &BukuRepositoryImpl{}
}

func (repository *BukuRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, buku domain.Buku) domain.Buku {
	SQL := "insert into data_buku(kode_buku,nama_buku,penerbit,gambar,deskripsi,pengarang,genre,rilis,harga,created_at,updated_at) values (?,?,?,?,?,?,?,?,?, NOW(), NOW())"

	_, err := tx.ExecContext(ctx, SQL, buku.Id, buku.Name, buku.Penerbit, buku.Image, buku.Deskripsi, buku.Pengarang, buku.Genre, buku.Rilis, buku.Harga)
	helper.PanicIfError(err)

	return buku
}

func (repository *BukuRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, buku domain.Buku, bukuId string) domain.Buku {
	SQL := "update data_buku set nama_buku = ?,penerbit = ?,gambar = ?,pengarang = ?, deskripsi = ?,genre =?,rilis = ?,harga = ?, updated_at = NOW() where kode_buku = ?"
	_, err := tx.ExecContext(ctx, SQL, buku.Name, buku.Penerbit, buku.Image, buku.Pengarang, buku.Deskripsi, buku.Genre, buku.Rilis, buku.Harga, bukuId)
	helper.PanicIfError(err)

	return buku
}

func (repository *BukuRepositoryImpl) UpdateStok(ctx context.Context, tx *sql.Tx, buku domain.Buku, bukuId string) domain.Buku {
	SQL := "update data_buku set stok = ?, updated_at = NOW() where kode_buku = ?"
	_, err := tx.ExecContext(ctx, SQL, buku.Stok, bukuId)
	helper.PanicIfError(err)

	return buku
}

func (repository *BukuRepositoryImpl) ReduceStok(ctx context.Context, tx *sql.Tx, buku domain.Buku, bukuId string) domain.Buku {
	SQL := "update data_buku set stok = stok - ?, updated_at = NOW() where kode_buku = ?"
	_, err := tx.ExecContext(ctx, SQL, buku.Stok, bukuId)
	helper.PanicIfError(err)

	return buku
}

func (repository *BukuRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, bukuId string) {
	SQL := "update data_buku set deleted_at = NOW() where kode_buku = ?"
	_, err := tx.ExecContext(ctx, SQL, bukuId)
	helper.PanicIfError(err)

}

func (repository *BukuRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bukuId string) (domain.Buku, error) {
	SQL := "select kode_buku ,nama_buku, penerbit, gambar, deskripsi, pengarang, genre, rilis, harga, stok from data_buku where kode_buku = ?"
	rows, err := tx.QueryContext(ctx, SQL, bukuId)
	helper.PanicIfError(err)
	defer rows.Close()

	buku := domain.Buku{}
	if rows.Next() {
		err := rows.Scan(&buku.Id, &buku.Name, &buku.Penerbit, &buku.Image, &buku.Deskripsi, &buku.Pengarang, &buku.Genre,&buku.Rilis, &buku.Harga,  &buku.Stok)
		helper.PanicIfError(err)
		return buku, nil
	} else {
		return buku, errors.New("buku tidak ditemukan")
	}
}

func (repository *BukuRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, request web.FetchRequest) []domain.Buku {
	SQL := "select kode_buku, nama_buku, penerbit, gambar, deskripsi, pengarang, genre, rilis, harga, stok from data_buku";

	SQL += " where deleted_at is null"

	if(request.Keyword != "") {
		SQL += " and nama_buku like '%"+request.Keyword+"%' or penerbit like '%"+request.Keyword+"%' or deskripsi like '%"+request.Keyword+"%' or pengarang like '%"+request.Keyword+"%'"
	}

	if(request.Sort != "") {
		SQL += " order by nama_buku "+request.Sort+", penerbit "+request.Sort+", deskripsi "+request.Sort+", pengarang "+request.Sort+""
	}
	
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var bukus []domain.Buku
	for rows.Next() {
		buku := domain.Buku{}
		err := rows.Scan(&buku.Id, &buku.Name, &buku.Penerbit, &buku.Image, &buku.Deskripsi, &buku.Pengarang, &buku.Genre,&buku.Rilis, &buku.Harga, &buku.Stok)
		helper.PanicIfError(err)
		bukus = append(bukus, buku)
	}
	return bukus
}
