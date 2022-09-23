package usecase

import (
	"context"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/repository"
)

type AuthUsecase interface {
	GetUserByID(ctx context.Context, id string) (*model.AuthUser, error)
}

type authUsecase struct {
	repository repository.AuthRepository
}

func NewAuthUsecase(r repository.AuthRepository) AuthUsecase {
	return &authUsecase{
		repository: r,
	}
}

func (u *authUsecase) GetUserByID(ctx context.Context, id string) (*model.AuthUser, error) {
	return u.repository.GetUserByID(ctx, id)
}
