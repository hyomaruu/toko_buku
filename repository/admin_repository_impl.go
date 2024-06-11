package repository

import (
	"context"
	"database/sql"
	"errors"
	"toko-buku/helper"
	"toko-buku/model/domain"
)

type AdminRepositoryImpl struct {
}

func NewAdminRepositry() AdminRepository {
	return &AdminRepositoryImpl{}
}

func (repository *AdminRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, admin domain.Admin) domain.Admin {
	SQL := "insert into data_admin(username, password, nama_lengkap, role) values (?,?,?,?)"

	val, err := tx.ExecContext(ctx, SQL, admin.Username, admin.Password, admin.Name, admin.Role)
	helper.PanicIfError(err)

	id, err := val.LastInsertId()

	admin.Id = id

	return admin
}
func (repository *AdminRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, admin domain.Admin) domain.Admin {
	SQL := "update data_admin set username = ?, password = ?, nama_lengkap = ?, role = ? where id_admin = ?"
	_, err := tx.ExecContext(ctx, SQL, admin.Username, admin.Password, admin.Name, admin.Role, admin.Id)
	helper.PanicIfError(err)

	return admin
}
func (repository *AdminRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, admin domain.Admin) {
	SQL := "delete from data_admin where id_admin = ?"
	_, err := tx.ExecContext(ctx, SQL, admin.Id)
	helper.PanicIfError(err)
}
func (repository *AdminRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, adminId int) (domain.Admin, error) {
	SQL := "select * from data_admin where id_admin = ?"
	rows, err := tx.QueryContext(ctx, SQL, adminId)
	helper.PanicIfError(err)
	defer rows.Close()

	admin := domain.Admin{}
	if rows.Next() {
		err := rows.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.Name, &admin.Role)
		helper.PanicIfError(err)
		return admin, nil
	} else {
		return admin, errors.New("admin is not found")
	}
}
func (repository *AdminRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Admin {
	SQL := "select * from data_admin"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var admins []domain.Admin
	for rows.Next() {
		admin := domain.Admin{}
		err := rows.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.Name, &admin.Role)
		helper.PanicIfError(err)
		admins = append(admins, admin)
	}
	return admins
}
