package datasource

import (
	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/repository"
	"gorm.io/gorm"
	"time"
)

type SomeRepository struct {
	db *gorm.DB
}

func NewSomeRepository(db *gorm.DB) repository.SomethingRepository {
	return &SomeRepository{
		db: db,
	}
}

type Something struct {
	ID        int       `gorm:"primary_key"`
	Name      string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (s Something) TableName() string {
	return "new_somethings"
}

func (s Something) toDomainModel() *model.Something {
	return &model.Something{
		ID:        s.ID,
		Name:      s.Name,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func (r *SomeRepository) FindAll() ([]*model.Something, error) {
	var somethings []*Something
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
