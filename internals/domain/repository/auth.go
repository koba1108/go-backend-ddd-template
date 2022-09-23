package repository

import (
	"context"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
)

type AuthRepository interface {
	GetUserByID(ctx context.Context, id string) (*model.AuthUser, error)
}
