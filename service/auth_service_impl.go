package service

import (
	"context"
	"database/sql"
	"toko-buku/exception"
	"toko-buku/helper"
	"toko-buku/model/web"
	"toko-buku/repository"

	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	DB 		*sql.DB
	Validate *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, DB *sql.DB, validate *validator.Validate) AuthService{
	return &AuthServiceImpl{
		AuthRepository:authRepository,
		DB : DB,
		Validate : validate,
	}
}


func (service *AuthServiceImpl) Authentication(ctx context.Context, auth web.AuthRequest) web.AdminResponse{
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	admin, err := service.AuthRepository.Authentication(ctx,tx, auth)
	if err != nil{ 
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToAdminResponse(admin)
}
