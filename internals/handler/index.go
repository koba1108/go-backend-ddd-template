package handler

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewSampleHandler),
	fx.Provide(NewAuthHandler),
	fx.Provide(NewSomethingHandler),
	fx.Provide(NewUserHandler),
)
