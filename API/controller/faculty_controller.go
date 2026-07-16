package controller

import (
	"mysql/constant/share"
	"mysql/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FacultyController struct {
	service service.FacultyService
}

func NewFacultyController() FacultyController {
	return FacultyController{
		service: service.NewFacultyService(),
	}
}

func (cr *FacultyController) GetFaculty(c *gin.Context) {
	data, err := cr.service.GetFaculty(c)
	if err != nil {
		share.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	share.RespondDate(c, http.StatusOK, data)
}
