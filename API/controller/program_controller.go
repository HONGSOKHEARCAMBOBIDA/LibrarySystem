package controller

import (
	"mysql/constant/apperror"
	"mysql/constant/share"
	"mysql/service"
	"mysql/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProgramController struct {
	service service.ProgramService
}

func NewProgramController() ProgramController {
	return ProgramController{
		service: service.NewProgramService(),
	}
}

func (cr *ProgramController) GetProgram(c *gin.Context) {
	departmentid, ok := utils.GetParamID(c)
	if !ok {
		return
	}
	data, err := cr.service.GetProgram(c, departmentid)
	if err != nil {
		status := apperror.HTTPStatus(err)
		msg := apperror.ClientMessage(err)
		share.ResponseError(c, status, msg)
		return
	}
	share.RespondDate(c, http.StatusOK, data)
}
