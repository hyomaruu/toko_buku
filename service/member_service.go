package service

import (
	"context"
	"toko-buku/model/web"
)

type MemberService interface {
	Create(ctx context.Context, request web.MemberCreateRequest) web.MemberResponse
	Update(ctx context.Context, request web.MemberUpdateRequest) web.MemberResponse
	Delete(ctx context.Context, memberId int)
	FindById(ctx context.Context, memberId int) web.MemberResponse
	FindAll(ctx context.Context) []web.MemberResponse
}