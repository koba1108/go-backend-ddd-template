package usecase

import (
	"context"
	"errors"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/repository"
)

type UserUsecase interface {
	DeleteUser(ctx context.Context, userId string) error
}

type userUsecase struct {
	userRepository     repository.UserRepository
	userBankRepository repository.UserBankRepository
}

func NewUserUsecase(ur repository.UserRepository, ubr repository.UserBankRepository) UserUsecase {
	return &userUsecase{
		userRepository:     ur,
		userBankRepository: ubr,
	}
}

func (u *userUsecase) DeleteUser(ctx context.Context, userId string) error {
	user, err := u.userRepository.GetByID(ctx, userId)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	userBank, err := u.userBankRepository.GetByUserID(ctx, userId)
	if err != nil {
		return err
	}
	if userBank != nil {
		return errors.New("user bank account exists")
	}
	if err = user.IsDeletable(userBank); err != nil {
		return err
	}
	if err = u.userRepository.Delete(ctx, user); err != nil {
		return err
	}
	return nil
}
