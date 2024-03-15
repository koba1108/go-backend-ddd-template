package config

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewDatabaseConfig),
	fx.Provide(NewFirebaseConfig),
)
