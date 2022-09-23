package client

import (
	"github.com/koba1108/go-backend-ddd-template/internals/config"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(config.NewDatabaseConfigConfig, NewGorm),
	fx.Provide(config.NewFirebaseConfig, NewFirebase),
)
