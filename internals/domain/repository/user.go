package repository

import (
	"context"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
)

type UserRepository interface {
	Find(ctx context.Context) ([]*model.User, error)
	GetByID(ctx context.Context, userId int) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, input *model.UserUpdateInput) error
	Save(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, userId int) error
}
