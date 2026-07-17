package controller

import (
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
		share.RespondServiceError(c, err)
		return
	}
	share.RespondDate(c, http.StatusOK, data)
}
