package usecase

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewSampleUsecase),
	fx.Provide(NewSomethingUsecase),
	fx.Provide(NewAuthUsecase),
	fx.Provide(NewUserUsecase),
)
