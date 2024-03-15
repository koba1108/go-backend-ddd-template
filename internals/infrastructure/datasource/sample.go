package datasource

import (
	"context"
	"fmt"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/repository"
	"gorm.io/gorm"
)

type sampleRepository struct {
	db *gorm.DB
}

func NewSampleRepository(db *gorm.DB) repository.SampleRepository {
	return &sampleRepository{db: db}
}

func (sr *sampleRepository) GetDB() *gorm.DB {
	return sr.db
}

func (sr *sampleRepository) Find(ctx context.Context, params *model.SampleFindParams, p *model.Pagination) ([]*model.Sample, error) {
	var samples []*model.Sample
	conn := sr.GetDB()
	conn = conn.Limit(p.Limit).Offset(p.Offset)
	if params != nil {
		if params.Name != "" {
			conn = conn.Where("name = ?", params.Name)
		}
		if p.SortKey == "" {
			if p.IsDesc {
				conn.Order(fmt.Sprintf("%s %s", p.SortKey, "desc"))
			} else {
				conn.Order(fmt.Sprintf("%s %s", p.SortKey, "asc"))
			}
		}
	}
	if err := conn.Find(&samples).Error; err != nil {
		return nil, err
	}
	return samples, nil
}

func (sr *sampleRepository) FindByID(ctx context.Context, id int) (*model.Sample, error) {
	var sample model.Sample
	conn := sr.GetDB()
	if err := conn.First(&sample, id).Error; err != nil {
		return nil, err
	}
	return &sample, nil
}

func (sr *sampleRepository) Create(ctx context.Context, sample *model.Sample) error {
	conn := sr.GetDB()
	if err := conn.Create(sample).Error; err != nil {
		return err
	}
	return nil
}

func (sr *sampleRepository) Update(ctx context.Context, sample *model.Sample) error {
	conn := sr.GetDB()
	if err := conn.Save(sample).Error; err != nil {
		return err
	}
	return nil
}

func (sr *sampleRepository) Delete(ctx context.Context, id int) error {
	conn := sr.GetDB()
	if err := conn.Delete(&model.Sample{}, id).Error; err != nil {
		return err
	}
	return nil
}
