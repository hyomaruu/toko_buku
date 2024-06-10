package repository

import (
	"context"
	"database/sql"
	"toko-buku/model/domain"
)

type MemberRepository interface {
	Save(ctx context.Context, tx *sql.Tx, member domain.Member) domain.Member

	Update(ctx context.Context, tx *sql.Tx, member domain.Member, memberId int) domain.Member

	Delete(ctx context.Context, tx *sql.Tx, memberId int)

	FindById(ctx context.Context, tx *sql.Tx, memberId int) (domain.Member, error)

	FindAll(ctx context.Context, tx *sql.Tx) []domain.Member
}