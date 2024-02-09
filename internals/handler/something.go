package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/koba1108/go-backend-ddd-template/internals/usecase"
)

type SomethingHandler interface {
	List(ctx *gin.Context)
}

type somethingHandler struct {
	someUsecase usecase.SomethingUsecase
}

func NewSomethingHandler(u usecase.SomethingUsecase) SomethingHandler {
	return &somethingHandler{
		someUsecase: u,
	}
}

func (h *somethingHandler) List(ctx *gin.Context) {
	somethings, err := h.someUsecase.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, somethings)
}
