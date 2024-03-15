package datasource

import (
	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/repository"
	"gorm.io/gorm"
	"time"
)

type SomeOldRepository struct {
	db *gorm.DB
}

func NewSomeOldRepository(db *gorm.DB) repository.SomethingRepository {
	return &SomeOldRepository{
		db: db,
	}
}

type SomethingOld struct {
	ID        int       `gorm:"primary_key"`
	Name      string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (s SomethingOld) TableName() string {
	return "old_somethings"
}

func (s SomethingOld) toDomainModel() *model.Something {
	return &model.Something{
		ID:        s.ID,
		Name:      s.Name,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func (r *SomeOldRepository) FindAll() ([]*model.Something, error) {
	var somethings []*SomethingOld
	err := r.db.Find(&somethings).Error
	if err != nil {
		return nil, err
	}

	var results []*model.Something
	for _, s := range somethings {
		results = append(results, s.toDomainModel())
	}
	return results, nil
}
