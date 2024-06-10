package service

import (
	"context"
	"database/sql"
	"toko-buku/exception"
	"toko-buku/helper"
	"toko-buku/model/domain"
	"toko-buku/model/web"
	"toko-buku/repository"

	"github.com/go-playground/validator/v10"
)

type MemberServiceImpl struct {
	MemberRepository repository.MemberRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewMemberService(memberRepository repository.MemberRepository, DB *sql.DB, validate *validator.Validate) MemberService {
	return &MemberServiceImpl{
		MemberRepository: memberRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func(service *MemberServiceImpl) Create(ctx context.Context, request web.MemberCreateRequest) web.MemberResponse{
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	

	member := domain.Member{
		Name: request.Name,
		Email: request.Email,
		Telepon: request.Telepon,
		Alamat: request.Alamat,
	}

	member = service.MemberRepository.Save(ctx,tx,member)
	return helper.ToMemberResponse(member)
}

func(service *MemberServiceImpl) Update(ctx context.Context, request web.MemberUpdateRequest, memberId int) web.MemberResponse{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	member, err := service.MemberRepository.FindById(ctx,tx, memberId)
	if err!= nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	member.Name = request.Name
	member.Email = request.Email
	member.Telepon = request.Telepon
	member.Alamat = request.Alamat

	member = service.MemberRepository.Update(ctx, tx, member, memberId)
	return helper.ToMemberResponse(member)
}

func (service *MemberServiceImpl) Delete(ctx context.Context, memberId int){
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.MemberRepository.FindById(ctx, tx, memberId)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.MemberRepository.Delete(ctx, tx, memberId)
}

func(service *MemberServiceImpl) FindById(ctx context.Context, memberId int)web.MemberResponse{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	member, err := service.MemberRepository.FindById(ctx,tx, memberId)
	if err!= nil{
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToMemberResponse(member)
}

func(service *MemberServiceImpl) FindAll(ctx context.Context) []web.MemberResponse{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	members := service.MemberRepository.FindAll(ctx,tx)
	return helper.ToMemberResponses(members)
}