package datasource

import (
	"context"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/repository"
	"gorm.io/gorm"
	"time"
)

type UserBankRepository struct {
	db *gorm.DB
}

type UserBank struct {
	ID        int       `gorm:"id"`
	UserId    int       `gorm:"userId"`
	BankName  string    `gorm:"bankName"`
	Status    string    `gorm:"status"`
	CreatedAt time.Time `gorm:"createdAt"`
	UpdatedAt time.Time `gorm:"updatedAt"`
}

func NewUserBankRepository(db *gorm.DB) repository.UserBankRepository {
	return &UserBankRepository{
		db: db,
	}
}

func (r *UserBankRepository) GetByUserID(ctx context.Context, userId int) (*model.UserBank, error) {
	// todo: implement
	return nil, nil
}
