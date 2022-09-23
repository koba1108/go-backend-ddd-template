package usecase

import (
	"github.com/koba1108/go-backend-ddd-template/internals/infrastructure/datasource"
	"github.com/koba1108/go-backend-ddd-template/internals/infrastructure/externals"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewSampleUsecase, datasource.NewSampleRepository),
	fx.Provide(NewAuthUsecase, externals.NewFirebaseAuthRepository),
)
