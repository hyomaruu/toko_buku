package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
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

func (repository *KeranjangRepositoryImpl) UpdateKodeTransaksi(ctx context.Context, tx *sql.Tx, keranjangId int, kodeTransaksi int) domain.Keranjang {
	SQL := "update keranjang set kode_transaksi = ? where kode_keranjang = ?"

	fmt.Printf(strconv.Itoa(kodeTransaksi), keranjangId)
	val, err := tx.ExecContext(ctx, SQL, kodeTransaksi, keranjangId)
	helper.PanicIfError(err)

	fmt.Printf("val: %v\n", val)

	keranjang := domain.Keranjang{}

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

func (repository *KeranjangRepositoryImpl) FindByKodeTransaksi(ctx context.Context, tx *sql.Tx, kodeTransaksi int) []domain.Keranjang {
	SQL := "select kode_keranjang, tipe_barang, kode_barang, total_barang, harga from keranjang where kode_transaksi = ?"
	rows, err := tx.QueryContext(ctx, SQL, kodeTransaksi)
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
