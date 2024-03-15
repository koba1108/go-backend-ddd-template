package usecase

import (
	"context"
	"errors"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/repository"
)

type UserUsecase interface {
	List(ctx context.Context) ([]*model.User, error)
	Get(ctx context.Context, userId int) (*model.User, error)
	Create(ctx context.Context, input *model.UserCreateInput) (*model.User, error)
	Update(ctx context.Context, input *model.UserUpdateInput) (*model.User, error)
	Save(ctx context.Context, input *model.UserUpdateInput) (*model.User, error)
	Delete(ctx context.Context, userId int) error
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

func (u *userUsecase) List(ctx context.Context) ([]*model.User, error) {
	users, err := u.userRepository.Find(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUsecase) Get(ctx context.Context, userId int) (*model.User, error) {
	user, err := u.userRepository.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) Create(ctx context.Context, input *model.UserCreateInput) (*model.User, error) {
	newUser, err := model.NewUser(input.Username, input.Email)
	if err != nil {
		return nil, err
	}
	if newUser, err = u.userRepository.Create(ctx, newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *userUsecase) Update(ctx context.Context, input *model.UserUpdateInput) (*model.User, error) {
	user, err := u.userRepository.GetByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	if err = u.userRepository.Update(ctx, input); err != nil {
		return nil, err
	}
	if input.Email != nil {
		user.Email = *input.Email
	}
	if input.Username != nil {
		user.Username = *input.Username
	}
	if input.IsWithdraw != nil {
		user.IsWithdraw = *input.IsWithdraw
	}
	return user, nil
}

func (u *userUsecase) Save(ctx context.Context, input *model.UserUpdateInput) (*model.User, error) {
	user, err := u.userRepository.GetByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	if input.Email != nil {
		user.Email = *input.Email
	}
	if input.Username != nil {
		user.Username = *input.Username
	}
	if input.IsWithdraw != nil {
		user.IsWithdraw = *input.IsWithdraw
	}
	if user, err = u.userRepository.Save(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) Delete(ctx context.Context, userId int) error {
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
	if err = u.userRepository.Delete(ctx, user.ID); err != nil {
		return err
	}
	return nil
}
