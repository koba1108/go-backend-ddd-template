package repository

import (
	"context"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
)

type UserBankRepository interface {
	GetByUserID(ctx context.Context, userId string) (*model.UserBank, error)
}
