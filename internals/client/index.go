package client

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewGorm),
	fx.Provide(NewFirebase),
)
