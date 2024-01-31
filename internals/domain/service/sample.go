package service

import (
	"context"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
)

type SampleService interface {
	GetSampleWithAuthUser(ctx context.Context, id int) (*model.Sample, *model.AuthUser, error)
}
