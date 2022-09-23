package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/koba1108/go-backend-ddd-template/internals/usecase"
)

type SampleHandler interface {
	List(c *gin.Context)
	Get(c *gin.Context)
	Post(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
}

type SampleCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type SampleUpdateRequest struct {
	Name *string `json:"name" binding:"required"`
}

func NewSampleHandler(su usecase.SampleUsecase) SampleHandler {
	return &sampleHandler{
		sampleUsecase: su,
	}
}

type sampleHandler struct {
	sampleUsecase usecase.SampleUsecase
}

// List godoc
// @router /sample [get]
// @description サンプルをリストで取得します
// @accept application/json
// @Success 200 {object} []model.Sample
// @Failure 400 {object} helper.Error
func (sh *sampleHandler) List(c *gin.Context) {

}

// Get godoc
// @router /sample/:id [get]
// @description サンプルを1件取得します
// @accept application/json
// @Param id path string true "id"
// @Success 200 {object} model.Sample
// @Failure 400 {object} helper.Error
func (sh *sampleHandler) Get(c *gin.Context) {

}

// Post godoc
// @router /sample [post]
// @description サンプルを作成します
// @accept application/json
// @Param params body SampleCreateRequest true "body"
// @Success 201 {object} model.Sample
// @Failure 400 {object} helper.Error
func (sh *sampleHandler) Post(c *gin.Context) {

}

// Put godoc
// @router /sample/:id [put]
// @description サンプルを更新します
// @accept application/json
// @Param id path string true "id"
// @Param params body SampleUpdateRequest true "body"
// @Success 200 {object} model.Sample
// @Failure 400 {object} helper.Error
func (sh *sampleHandler) Put(c *gin.Context) {

}

// Delete godoc
// @router /sample/:id [delete]
// @description サンプルを削除します
// @accept application/json
// @Param id path string true "id"
// @Success 204
// @Failure 400 {object} helper.Error
func (sh *sampleHandler) Delete(c *gin.Context) {

}
