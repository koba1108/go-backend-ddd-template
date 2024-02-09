package main

import (
	"github.com/gin-gonic/gin"
	"github.com/koba1108/go-backend-ddd-template/internals/client"
	"github.com/koba1108/go-backend-ddd-template/internals/config"
	"github.com/koba1108/go-backend-ddd-template/internals/handler"
	"github.com/koba1108/go-backend-ddd-template/internals/infrastructure/datasource"
	"github.com/koba1108/go-backend-ddd-template/internals/usecase"
)

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	if err := r.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	dbConf := config.NewDatabaseConfig()
	db, err := client.NewGorm(dbConf)
	if err != nil {
		panic(err)
	}
	sampleRepo := datasource.NewSampleRepository(db)
	sampleUsecase := usecase.NewSampleUsecase(sampleRepo)
	sampleHandler := handler.NewSampleHandler(sampleUsecase)

	api2 := r.Group("api2")
	{
		api2.GET("/sample", sampleHandler.Get)
	}
}
