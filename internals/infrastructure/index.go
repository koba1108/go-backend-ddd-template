package infrastructure

import (
	"github.com/koba1108/go-backend-ddd-template/internals/infrastructure/composite"
	"github.com/koba1108/go-backend-ddd-template/internals/infrastructure/datasource"
	"github.com/koba1108/go-backend-ddd-template/internals/infrastructure/externals"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(composite.NewSampleService),
	fx.Provide(datasource.NewSampleRepository),
	fx.Provide(datasource.NewSomeRepository),
	fx.Provide(datasource.NewUserRepository),
	fx.Provide(datasource.NewUserBankRepository),
	fx.Provide(externals.NewFirebaseAuthRepository),
)
