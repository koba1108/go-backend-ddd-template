package usecase

import (
	"context"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/repository"
)

type SampleUsecase interface {
	Find(ctx context.Context, limit, offset int, sortKey string, isAsc bool) ([]*model.Sample, error)
}

type sampleUsecase struct {
	repository repository.SampleRepository
}

func NewSampleUsecase(r repository.SampleRepository) SampleUsecase {
	return &sampleUsecase{
		repository: r,
	}
}

func (u *sampleUsecase) Find(ctx context.Context, limit, offset int, sortKey string, isAsc bool) ([]*model.Sample, error) {
	pagination, err := model.NewPagination(limit, offset, sortKey, isAsc)
	if err != nil {
		return nil, err
	}
	return u.repository.Find(ctx, pagination)
}
