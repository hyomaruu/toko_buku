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
		Id: request.Id,
		Name: request.Name,
<<<<<<< HEAD
		Email: request.Email,
		Telepon: request.Telepon,
		Alamat: request.Alamat,
=======
>>>>>>> 7a1f8e50849a6ee819f6ea835677610a48fcb3f0
	}

	member = service.MemberRepository.Save(ctx,tx,member)
	return helper.ToMemberResponse(member)
}

func(service *MemberServiceImpl) Update(ctx context.Context, request web.MemberUpdateRequest) web.MemberResponse{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	member, err := service.MemberRepository.FindById(ctx,tx,request.Id)
	if err!= nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	member.Name = request.Name
<<<<<<< HEAD
	member.Email = request.Email
	member.Telepon = request.Telepon
	member.Alamat = request.Alamat
=======
>>>>>>> 7a1f8e50849a6ee819f6ea835677610a48fcb3f0

	member = service.MemberRepository.Update(ctx,tx,member)
	return helper.ToMemberResponse(member)
}

func (service *MemberServiceImpl) Delete(ctx context.Context, MemberId int){
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	member, err := service.MemberRepository.FindById(ctx,tx,MemberId)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.MemberRepository.Delete(ctx,tx,member)
}

func(service *MemberServiceImpl) FindById(ctx context.Context, MemberId int)web.MemberResponse{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	member, err := service.MemberRepository.FindById(ctx,tx,MemberId)
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