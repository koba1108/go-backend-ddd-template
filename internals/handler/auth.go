package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/koba1108/go-backend-ddd-template/internals/usecase"
)

type AuthHandler interface {
	Get(c *gin.Context)
}

type authHandler struct {
	au usecase.AuthUsecase
}

func NewAuthHandler(au usecase.AuthUsecase) AuthHandler {
	return &authHandler{au: au}
}

// Get godoc
// @router /auth/:id [get]
// @description 認証を行います
// @accept application/json
// @Param id path string true "id"
// @Success 200 {object} model.AuthUser
// @Failure 400 {object} helper.Error
func (ah *authHandler) Get(c *gin.Context) {
	authUser, err := ah.au.GetUserByID(c, c.Param("id"))
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, authUser)
}
