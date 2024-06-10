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

type AdminServiceImpl struct {
	AdminRepository repository.AdminRepository
	DB 		*sql.DB
	Validate *validator.Validate
}

func NewAdminService(adminRepository repository.AdminRepository, DB *sql.DB, validate *validator.Validate) AdminService{
	return &AdminServiceImpl{
		AdminRepository:adminRepository,
		DB : DB,
		Validate : validate,
	}
}

func (service *AdminServiceImpl) Create(ctx context.Context, request web.AdminCreateRequest) web.AdminResponse{
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	admin := domain.Admin{
		Username: request.Username,
		Password: request.Password,
		Name: request.Name,
		Role: request.Role,
	}

	admin = service.AdminRepository.Save(ctx,tx, admin)
	return helper.ToAdminResponse(admin)

}

func (service *AdminServiceImpl) Update(ctx context.Context, request web.AdminUpdateRequest, adminId int) web.AdminResponse{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	admin, err := service.AdminRepository.FindById(ctx,tx, adminId)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	admin.Username = request.Username
	admin.Password =request.Password
	admin.Name = request.Name
	admin.Role = request.Role

	admin = service.AdminRepository.Update(ctx,tx, admin)
	return helper.ToAdminResponse(admin)
}

func (service *AdminServiceImpl) Delete(ctx context.Context, adminId int) {
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	admin, err := service.AdminRepository.FindById(ctx,tx, adminId)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.AdminRepository.Delete(ctx,tx, admin)
}

func (service *AdminServiceImpl) FindById(ctx context.Context, adminId int) web.AdminResponse{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	admin, err := service.AdminRepository.FindById(ctx,tx, adminId)
	if err != nil{ 
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToAdminResponse(admin)
}

func (service *AdminServiceImpl) FindAll(ctx context.Context) []web.AdminResponse{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	admins := service.AdminRepository.FindAll(ctx,tx)
	return helper.ToAdminResponses(admins)
}