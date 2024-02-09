package usecase

import (
	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/repository"
)

type SomethingUsecase interface {
	FindAll() ([]*model.Something, error)
}

type somethingUsecase struct {
	repository repository.SomethingRepository
}

func NewSomethingUsecase(r repository.SomethingRepository) SomethingUsecase {
	return &somethingUsecase{
		repository: r,
	}
}

func (u *somethingUsecase) FindAll() ([]*model.Something, error) {
	return u.repository.FindAll()
}
