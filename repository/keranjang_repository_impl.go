package repository

import (
	"context"
	"database/sql"
	"errors"
	"toko-buku/helper"
	"toko-buku/model/domain"
)

type KeranjangRepositoryImpl struct {
}

func NewKeranjangRepository() KeranjangRepository {
	return &KeranjangRepositoryImpl{}
}

func (repository *KeranjangRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, keranjang domain.Keranjang) domain.Keranjang {
	SQL := "insert into keranjang(tipe_barang,kode_barang,total_barang,harga) values (?,?,?,?)"

	val, err := tx.ExecContext(ctx, SQL, keranjang.TipeBarang, keranjang.KodeBarang, keranjang.TotalBarang, keranjang.Harga)
	helper.PanicIfError(err)

	id, err := val.LastInsertId()

	keranjang.KodeKeranjang = id

	return keranjang
}

func (repository *KeranjangRepositoryImpl) UpdateTotalBarang(ctx context.Context, tx *sql.Tx, keranjang domain.Keranjang, keranjangId int) domain.Keranjang {
	SQL := "update keranjang set total_barang = ? where kode_keranjang = ?"
	_, err := tx.ExecContext(ctx, SQL, keranjang.TotalBarang, keranjangId)
	helper.PanicIfError(err)

	return keranjang
}

func (repository *KeranjangRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, keranjangId int) {
	SQL := "delete from keranjang where kode_keranjang = ?"
	_, err := tx.ExecContext(ctx, SQL, keranjangId)
	helper.PanicIfError(err)
}

func (repository *KeranjangRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, keranjangId int) (domain.Keranjang, error) {
	SQL := "select kode_keranjang, tipe_barang, kode_barang, total_barang, harga from keranjang where kode_keranjang = ?"
	rows, err := tx.QueryContext(ctx, SQL, keranjangId)
	helper.PanicIfError(err)
	defer rows.Close()

	keranjang := domain.Keranjang{}
	if rows.Next() {
		err := rows.Scan(&keranjang.KodeKeranjang, &keranjang.TipeBarang, &keranjang.KodeBarang, &keranjang.TotalBarang,  &keranjang.Harga)
		helper.PanicIfError(err)
		return keranjang, nil
	} else {
		return keranjang, errors.New("keranjang is not found")
	}
}

func (repository *KeranjangRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Keranjang {
	SQL := "select kode_keranjang, tipe_barang, kode_barang, total_barang, harga from keranjang"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var keranjangs []domain.Keranjang
	for rows.Next() {
		keranjang := domain.Keranjang{}
		err := rows.Scan(&keranjang.KodeKeranjang, &keranjang.TipeBarang, &keranjang.KodeBarang, &keranjang.TotalBarang,  &keranjang.Harga)
		helper.PanicIfError(err)
		keranjangs = append(keranjangs, keranjang)
	}
	return keranjangs
}


// SELECT kode_keranjang, tipe_barang, kode_barang, total_barang, t_keranjang.harga, 
// 	CASE WHEN tipe_barang = 'BUKU' THEN t_buku.nama_buku ELSE null END AS nama,
//     CASE WHEN tipe_barang = 'AKSESORIS' THEN t_aksesoris.nama_barang ELSE null END AS nama
// 	FROM keranjang AS t_keranjang 
// 	LEFT JOIN data_buku t_buku ON t_keranjang.kode_barang = t_buku.kode_buku AND t_keranjang.tipe_barang = 'BUKU'
// 	LEFT JOIN aksesoris t_aksesoris ON t_keranjang.kode_barang = t_aksesoris.kode AND t_keranjang.tipe_barang = 'AKSESORIS'

	

// 	SELECT kode_keranjang, tipe_barang, kode_barang, total_barang, t_keranjang.harga, 
// 	CASE WHEN tipe_barang = 'BUKU' THEN t_buku.nama_buku ELSE null ,
//     WHEN tipe_barang = 'AKSESORIS' THEN t_aksesoris.nama_barang ELSE null END AS nama
// 	FROM keranjang AS t_keranjang 
// 	LEFT JOIN data_buku t_buku ON t_keranjang.kode_barang = t_buku.kode_buku AND t_keranjang.tipe_barang = 'BUKU'
// 	LEFT JOIN aksesoris t_aksesoris ON t_keranjang.kode_barang = t_aksesoris.kode AND t_keranjang.tipe_barang = 'AKSESORIS'
