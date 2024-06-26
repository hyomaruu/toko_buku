package repository

import (
	"context"
	"database/sql"
	"errors"
	"toko-buku/helper"
	"toko-buku/model/domain"
	"toko-buku/model/web"
)

type AdminRepositoryImpl struct {
}

func NewAdminRepositry() AdminRepository {
	return &AdminRepositoryImpl{}
}

func (repository *AdminRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, admin domain.Admin) domain.Admin {
	SQL := "insert into data_admin(username, password, nama_lengkap, role, created_at, updated_at) values (?,?,?,?, NOW(), NOW())"

	val, err := tx.ExecContext(ctx, SQL, admin.Username, admin.Password, admin.Name, admin.Role)
	helper.PanicIfError(err)

	id, err := val.LastInsertId()

	admin.Id = id

	return admin
}

func (repository *AdminRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, admin domain.Admin) domain.Admin {
	SQL := "update data_admin set username = ?, password = ?, nama_lengkap = ?, role = ?, updated_at = NOW() where id_admin = ?"
	_, err := tx.ExecContext(ctx, SQL, admin.Username, admin.Password, admin.Name, admin.Role, admin.Id)
	helper.PanicIfError(err)

	return admin
}
func (repository *AdminRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, admin domain.Admin) {
	SQL := "update data_admin set deleted_at = NOW() where id_admin = ?"
	_, err := tx.ExecContext(ctx, SQL, admin.Id)
	helper.PanicIfError(err)
}

func (repository *AdminRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, adminId int) (domain.Admin, error) {
	SQL := "select id_admin, username, nama_lengkap, role from data_admin where id_admin = ? and deleted_at is null"
	rows, err := tx.QueryContext(ctx, SQL, adminId)
	helper.PanicIfError(err)
	defer rows.Close()

	admin := domain.Admin{}
	if rows.Next() {
		err := rows.Scan(&admin.Id, &admin.Username, &admin.Name, &admin.Role)
		helper.PanicIfError(err)
		return admin, nil
	} else {
		return admin, errors.New("admin is not found")
	}
}
func (repository *AdminRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, request web.FetchRequest) []domain.Admin {
	SQL := "select id_admin, username, nama_lengkap, role from data_admin";

	SQL += " where deleted_at is null"

	if(request.Keyword != "") {
		SQL += " and username like '%"+request.Keyword+"%' or nama_lengkap like '%"+request.Keyword+"%'"
	}

	if(request.Sort != "") {
		SQL += " order by username "+request.Sort+" , nama_lengkap "+request.Sort+""
	}
	
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var admins []domain.Admin
	for rows.Next() {
		admin := domain.Admin{}
		err := rows.Scan(&admin.Id, &admin.Username, &admin.Name, &admin.Role)
		helper.PanicIfError(err)
		admins = append(admins, admin)
	}
	return admins
}

func (repository *AdminRepositoryImpl) RoleCounter(ctx context.Context, tx *sql.Tx) []domain.AdminCounter {
	SQL := "SELECT role, COUNT(*) as counter FROM data_admin GROUP BY role"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var counters []domain.AdminCounter

	for rows.Next() {
		counter := domain.AdminCounter{}
		err := rows.Scan(&counter.Role, &counter.Counter,)
		helper.PanicIfError(err)
		counters = append(counters, counter)
	}
	return counters
}
