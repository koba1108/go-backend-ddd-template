package repository

import (
	"context"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
	Delete(ctx context.Context, user *model.User) error
}
