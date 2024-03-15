package datasource

import (
	"context"
	"gorm.io/gorm"

	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/repository"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{
		db: db,
	}
}

type User struct {
	ID         int        `gorm:"id"`
	Username   string     `gorm:"username"`
	Email      string     `gorm:"email"`
	IsWithdraw bool       `gorm:"isWithdraw"`
	CreatedAt  time.Time  `gorm:"createdAt"`
	UpdatedAt  time.Time  `gorm:"updatedAt"`
	DeletedAt  *time.Time `gorm:"deletedAt"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) toDomainModel() *model.User {
	return &model.User{
		ID:         u.ID,
		Username:   u.Username,
		Email:      u.Email,
		IsWithdraw: u.IsWithdraw,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

func (r *UserRepository) Find(ctx context.Context) ([]*model.User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	var models []*model.User
	for _, user := range users {
		models = append(models, user.toDomainModel())
	}
	return models, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*model.User, error) {
	var user User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return user.toDomainModel(), nil
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	record := r.toDataModel(user)
	if err := r.db.Create(record).Error; err != nil {
		return nil, err
	}
	return record.toDomainModel(), nil
}

func (r *UserRepository) Update(ctx context.Context, input *model.UserUpdateInput) error {
	if err := r.db.Updates(input).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Save(ctx context.Context, user *model.User) (*model.User, error) {
	record := r.toDataModel(user)
	if err := r.db.Save(record).Error; err != nil {
		return nil, err
	}
	return record.toDomainModel(), nil
}

func (r *UserRepository) Delete(ctx context.Context, userId int) error {
	if err := r.db.Delete(&User{ID: userId}).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) toDataModel(u *model.User) *User {
	return &User{
		ID:         u.ID,
		Username:   u.Username,
		Email:      u.Email,
		IsWithdraw: u.IsWithdraw,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
		DeletedAt:  u.DeletedAt,
	}
}
