package controller

import (
	"mysql/constant/share"
	"mysql/request"
	"mysql/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	service service.AuthorService
}

func NewAuthorController() AuthorController {
	return AuthorController{
		service: service.NewAuthorService(),
	}
}

func (cr *AuthorController) CreateAuthor(c *gin.Context) {
	var input request.AuthorRequestCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		share.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := cr.service.CreateAuthor(c, input); err != nil {
		share.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	share.ResponseSuccess(c, http.StatusOK, share.Created)
}
