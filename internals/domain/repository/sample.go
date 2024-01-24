package repository

import (
	"context"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
)

type SampleRepository interface {
	Find(ctx context.Context, params *model.SampleFindParams, p *model.Pagination) ([]*model.Sample, error)
	FindByID(ctx context.Context, id int) (*model.Sample, error)
	Create(ctx context.Context, sample *model.Sample) error
	Update(ctx context.Context, sample *model.Sample) error
	Delete(ctx context.Context, id int) error
}
