package repository

import (
	"context"
	"database/sql"
	"errors"
	"toko-buku/helper"
	"toko-buku/model/domain"
	"toko-buku/model/web"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepositry() AuthRepository {
	return &AuthRepositoryImpl{}
}


func (repository *AuthRepositoryImpl) Authentication(ctx context.Context, tx *sql.Tx, auth web.AuthRequest) (domain.Admin, error) {
	SQL := "select id_admin, username, nama_lengkap, role from data_admin where username = ? AND password = ?"
	rows, err := tx.QueryContext(ctx, SQL, auth.Username, auth.Password)
	helper.PanicIfError(err)
	defer rows.Close()

	admin := domain.Admin{}
	if rows.Next() {
		err := rows.Scan(&admin.Id ,&admin.Username, &admin.Name, &admin.Role)
		helper.PanicIfError(err)
		return admin, nil
	} else {
		return admin, errors.New("admin is not found")
	}
}