package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
	"github.com/koba1108/go-backend-ddd-template/internals/usecase"
	"net/http"
)

type UserHandler interface {
	List(c *gin.Context)
	Get(c *gin.Context)
	Post(c *gin.Context)
	Put(c *gin.Context)
	Patch(c *gin.Context)
	Delete(c *gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: uu,
	}
}

func (h *userHandler) List(c *gin.Context) {
	users, err := h.userUsecase.List(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *userHandler) Get(c *gin.Context) {
	var req struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userUsecase.Get(c, req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *userHandler) Post(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userUsecase.Create(c, &model.UserCreateInput{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *userHandler) Put(c *gin.Context) {
	var req struct {
		ID         int     `uri:"id" binding:"required"`
		Username   *string `json:"username"`
		Email      *string `json:"email"`
		IsWithdraw *bool   `json:"isWithdraw"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userUsecase.Update(c, &model.UserUpdateInput{
		ID:         req.ID,
		Username:   req.Username,
		Email:      req.Email,
		IsWithdraw: req.IsWithdraw,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *userHandler) Patch(c *gin.Context) {
	var req struct {
		ID         int     `uri:"id" binding:"required"`
		Username   *string `json:"username"`
		Email      *string `json:"email"`
		IsWithdraw *bool   `json:"isWithdraw"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userUsecase.Save(c, &model.UserUpdateInput{
		ID:         req.ID,
		Username:   req.Username,
		Email:      req.Email,
		IsWithdraw: req.IsWithdraw,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)

}

func (h *userHandler) Delete(c *gin.Context) {
	var req struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.userUsecase.Delete(c, req.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
