package controller

import (
	"mysql/constant/share"
	"mysql/service"
	"mysql/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FilingCabinetController struct {
	service service.FilingCabinetService
}

func NewFilingCabinetController() FilingCabinetController {
	return FilingCabinetController{
		service: service.NewFilingCabinetService(),
	}
}

func (cr *FilingCabinetController) GetFilingCabinet(c *gin.Context) {
	cabinetID, ok := utils.GetParamID(c)
	if !ok {
		return
	}
	data, err := cr.service.GetFilingCabinet(c, cabinetID)
	if err != nil {
		share.RespondServiceError(c, err)
		return
	}
	share.RespondDate(c, http.StatusOK, data)
}
