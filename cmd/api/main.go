package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/koba1108/go-backend-ddd-template/internals/client"
	"github.com/koba1108/go-backend-ddd-template/internals/config"
	"github.com/koba1108/go-backend-ddd-template/internals/handler"
	"github.com/koba1108/go-backend-ddd-template/internals/infrastructure"
	"github.com/koba1108/go-backend-ddd-template/internals/usecase"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		client.Module,
		infrastructure.Module,
		usecase.Module,
		handler.Module,
		fx.Invoke(NewGinEngine),
	).Run()
}

func NewGinEngine(
	sh handler.SampleHandler,
	ah handler.AuthHandler,
	smh handler.SomethingHandler,
	uh handler.UserHandler,
) {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	if err := r.SetTrustedProxies(nil); err != nil {
		panic(fmt.Errorf("failed to set trusted proxies: %v", err))
	}

	api := r.Group("api")
	{
		apiV1 := api.Group("v1")
		{
			apiV1Auth := apiV1.Group("auth")
			{
				apiV1Auth.GET("/:id", ah.Get)
			}
			apiV1Sample := apiV1.Group("sample")
			{
				apiV1Sample.GET("", sh.List)
				apiV1Sample.GET(":id", sh.Get)
				apiV1Sample.POST("", sh.Post)
				apiV1Sample.PUT(":id", sh.Put)
				apiV1Sample.DELETE(":id", sh.Delete)
			}
			apiV1Something := apiV1.Group("something")
			{
				apiV1Something.GET("", smh.List)
			}
			apiV1User := apiV1.Group("user")
			{
				apiV1User.GET("", uh.List)
				apiV1User.GET(":id", uh.Get)
				apiV1User.POST("", uh.Post)
				apiV1User.PUT(":id", uh.Put)
				apiV1User.PATCH(":id", uh.Patch)
				apiV1User.DELETE(":id", uh.Delete)
			}
		}
	}
	if err := r.Run(); err != nil {
		panic(err)
	}
}
