package repository

import (
	"context"
	"database/sql"
	"errors"
	"toko-buku/helper"
	"toko-buku/model/domain"
)

type TransaksiRepositoryImpl struct {
}

func NewTransaksiRepository() TransaksiRepository {
	return &TransaksiRepositoryImpl{}
}

func (repository *TransaksiRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, transaksi domain.Transaksi) domain.Transaksi {
	SQL := "insert into transaksi(kode_transaksi,nomor_member,id_admin,tanggal,total_harga,created_at,updated_at) values (?,?,?,?,?,NOW(),NOW())"

	val, err := tx.ExecContext(ctx, SQL, transaksi.KodeTransaksi, transaksi.NomorMember, transaksi.IdAdmin, transaksi.Tanggal, transaksi.TotalHarga)
	helper.PanicIfError(err)
	id, err := val.LastInsertId()

	transaksi.KodeTransaksi = id

	return transaksi
}

func (repository *TransaksiRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, transaksiId int) (domain.Transaksi, error) {
	SQL := "select kode_transaksi, nomor_member, id_admin, tanggal, total_harga from transaksi where kode_transaksi = ?"
	rows, err := tx.QueryContext(ctx, SQL, transaksiId)
	helper.PanicIfError(err)
	defer rows.Close()

	transaksi := domain.Transaksi{}
	if rows.Next() {
		err := rows.Scan(&transaksi.KodeTransaksi, &transaksi.NomorMember, &transaksi.IdAdmin, &transaksi.Tanggal, &transaksi.TotalHarga)
		helper.PanicIfError(err)
		return transaksi, nil
	} else {
		return transaksi, errors.New("transaksi is not found")
	}
}

func (repository *TransaksiRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Transaksi {
	SQL := "select kode_transaksi, nomor_member, id_admin, tanggal, total_harga from transaksi"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var transaksis []domain.Transaksi
	for rows.Next() {
		transaksi := domain.Transaksi{}
		err := rows.Scan(&transaksi.KodeTransaksi, &transaksi.NomorMember, &transaksi.IdAdmin, &transaksi.Tanggal, &transaksi.TotalHarga)
		helper.PanicIfError(err)
		transaksis = append(transaksis, transaksi)
	}
	return transaksis
}
