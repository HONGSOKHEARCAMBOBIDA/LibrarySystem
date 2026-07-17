package controller

import (
	"mysql/constant/share"
	"mysql/request"
	"mysql/service"
	"mysql/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service service.CategoryService
}

func NewCategoryController() CategoryController {
	return CategoryController{
		service: service.NewCategoryService(),
	}
}

func (cr *CategoryController) GetCategory(c *gin.Context) {
	data, err := cr.service.GetCategory(c.Request.Context())
	if err != nil {
		share.RespondServiceError(c, err)
		return
	}
	share.RespondDate(c, http.StatusOK, data)
}

func (cr *CategoryController) CreateCategory(c *gin.Context) {
	var input request.CategoryRequestCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		share.ResponseError(c, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	if err := cr.service.CreateCategory(c.Request.Context(), input); err != nil {
		share.RespondServiceError(c, err)
		return
	}

	share.RespondDate(c, http.StatusCreated, share.Created)
}

func (cr *CategoryController) UpdateCategory(c *gin.Context) {
	id, ok := utils.GetParamID(c)
	if !ok {
		return
	}

	var input request.CategoryRequestUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		share.ResponseError(c, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	if err := cr.service.UpdateCategory(c.Request.Context(), id, input); err != nil {
		share.RespondServiceError(c, err)
		return
	}

	share.RespondDate(c, http.StatusOK, share.Updated)
}

func (cr *CategoryController) ToggleStatusCategory(c *gin.Context) {
	id, ok := utils.GetParamID(c)
	if !ok {
		share.ResponseError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	if err := cr.service.ToggleStatusCategory(c, id); err != nil {
		share.RespondServiceError(c, err)
		return
	}
}
