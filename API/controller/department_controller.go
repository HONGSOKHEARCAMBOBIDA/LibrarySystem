package controller

import (
	"mysql/constant/apperror"
	"mysql/constant/share"
	"mysql/service"
	"mysql/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DepartmentController struct {
	service service.DepartmentService
}

func NewDepartmentController() DepartmentController {
	return DepartmentController{
		service: service.NewDepartmentService(),
	}
}

func (cr *DepartmentController) GetDepartment(c *gin.Context) {
	facultyid, ok := utils.GetParamID(c)
	if !ok {
		return
	}
	data, err := cr.service.GetDepartment(c, facultyid)
	if err != nil {
		status := apperror.HTTPStatus(err)
		msg := apperror.ClientMessage(err)
		share.ResponseError(c, status, msg)
		return
	}
	share.RespondDate(c, http.StatusOK, data)
}
