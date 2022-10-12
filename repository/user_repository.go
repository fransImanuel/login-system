package repository

import (
	"context"
	"login-system/helper"
	"login-system/model/domain"
)

type UserRepository interface {
	Add(ctx context.Context, user domain.User)
	Update(ctx context.Context)
	Delete(ctx context.Context)
	Read(ctx context.Context)
}	

type UserRepositoryImpl struct {
	
}

func (repo *UserRepositoryImpl) Add(ctx context.Context, user domain.User)  {
	helper.PanicIfError(error)
	panic("Implement Me")
}

func (repo *UserRepositoryImpl) Update(ctx context.Context)  {
	panic("Implement Me")
}

func (repo *UserRepositoryImpl) Delete(ctx context.Context)  {
	panic("Implement Me")
}

func (repo *UserRepositoryImpl) Read(ctx context.Context)  {
	panic("Implement Me")
}