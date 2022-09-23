package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/koba1108/go-backend-ddd-template/internals/client"
	"github.com/koba1108/go-backend-ddd-template/internals/handler"
	"github.com/koba1108/go-backend-ddd-template/internals/usecase"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		client.Module,
		handler.Module,
		usecase.Module,
		fx.Invoke(NewGinEngine),
	).Run()
}

func NewGinEngine(sh handler.SampleHandler, ah handler.AuthHandler) {
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
		}
	}
	if err := r.Run(); err != nil {
		panic(err)
	}
}
