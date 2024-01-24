package usecase

import (
	"context"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/repository"
)

type SampleUsecase interface {
	Find(ctx context.Context, name string, limit, offset int, sortKey string, isAsc bool) ([]*model.Sample, error)
}

type sampleUsecase struct {
	repository repository.SampleRepository
}

func NewSampleUsecase(r repository.SampleRepository) SampleUsecase {
	return &sampleUsecase{
		repository: r,
	}
}

func (u *sampleUsecase) Find(ctx context.Context, name string, limit, offset int, sortKey string, isDesc bool) ([]*model.Sample, error) {
	pagination, err := model.NewPagination(limit, offset, sortKey, isDesc)
	if err != nil {
		return nil, err
	}
	params := &model.SampleFindParams{
		Name: name,
	}
	return u.repository.Find(ctx, params, pagination)
}
