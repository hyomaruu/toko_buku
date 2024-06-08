package repository

import (
	"context"
	"database/sql"
	"errors"
	"toko-buku/helper"
	"toko-buku/model/domain"
)

type MemberRepositoryImpl struct {
}

func NewMemberRepository() MemberRepository {
	return &MemberRepositoryImpl{}
}

func (repository *MemberRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, member domain.Member) domain.Member {
<<<<<<< HEAD
	SQL := "insert into data_member(id_member,nama_lengkap,email,telepon,alamat) values (?,?,?,?,?)"

	_, err := tx.ExecContext(ctx, SQL, member.Id, member.Name, member.Email, member.Telepon, member.Alamat)
=======
	SQL := "insert into data_member(id_member,nama_lengkap) values (?,?)"

	_, err := tx.ExecContext(ctx, SQL, member.Id, member.Name)
>>>>>>> 7a1f8e50849a6ee819f6ea835677610a48fcb3f0
	helper.PanicIfError(err)

	return member
}
func (repository *MemberRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, member domain.Member) domain.Member {
<<<<<<< HEAD
	SQL := "update data_member set nama_lengkap = ?,email = ?, telepon = ?, alamat = ? where id_member = ?"
	_, err := tx.ExecContext(ctx, SQL, member.Name, member.Email, member.Telepon, member.Alamat, member.Id)
=======
	SQL := "update data_member set nama_lengkap = ? where id_member = ?"
	_, err := tx.ExecContext(ctx, SQL, member.Name, member.Id)
>>>>>>> 7a1f8e50849a6ee819f6ea835677610a48fcb3f0
	helper.PanicIfError(err)

	return member
}
func (repository *MemberRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, member domain.Member) {
	SQL := "delete from data_member where id_member = ?"
	_, err := tx.ExecContext(ctx, SQL, member.Id)
	helper.PanicIfError(err)
}
func (repository *MemberRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, memberId int) (domain.Member, error) {
<<<<<<< HEAD
	SQL := "select id_member, nama_lengkap,email,telepon,alamat from data_member where id_member = ?"
=======
	SQL := "select id_member, nama_lengkap from data_member where id_member = ?"
>>>>>>> 7a1f8e50849a6ee819f6ea835677610a48fcb3f0
	rows, err := tx.QueryContext(ctx, SQL, memberId)
	helper.PanicIfError(err)
	defer rows.Close()

	member := domain.Member{}
	if rows.Next() {
<<<<<<< HEAD
		err := rows.Scan(&member.Id, &member.Name, &member.Email, &member.Telepon, &member.Alamat)
=======
		err := rows.Scan(&member.Id, &member.Name)
>>>>>>> 7a1f8e50849a6ee819f6ea835677610a48fcb3f0
		helper.PanicIfError(err)
		return member, nil
	} else {
		return member, errors.New("member is not found")
	}
}
func (repository *MemberRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Member {
<<<<<<< HEAD
	SQL := "select * from data_member"
=======
	SQL := "select id_member, nama_lengkap from data_member"
>>>>>>> 7a1f8e50849a6ee819f6ea835677610a48fcb3f0
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var members []domain.Member
	for rows.Next() {
		member := domain.Member{}
<<<<<<< HEAD
		err := rows.Scan(&member.Id, &member.Name, &member.Email, &member.Telepon, &member.Alamat)
=======
		err := rows.Scan(&member.Id, &member.Name)
>>>>>>> 7a1f8e50849a6ee819f6ea835677610a48fcb3f0
		helper.PanicIfError(err)
		members = append(members, member)
	}
	return members
}
