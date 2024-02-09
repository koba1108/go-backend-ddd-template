package repository

import "github.com/koba1108/go-backend-ddd-template/internals/domain/model"

type SomethingRepository interface {
	FindAll() ([]*model.Something, error)
}
