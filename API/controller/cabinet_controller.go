package controller

import (
	"mysql/constant/share"
	"mysql/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CabinetController struct {
	service service.CabinetService
}

func NewCabinetController() CabinetController {
	return CabinetController{
		service: service.NewCabinetService(),
	}
}

func (cr *CabinetController) GetCabinet(c *gin.Context) {
	data, err := cr.service.GetCabinet(c.Request.Context())
	if err != nil {
		share.RespondServiceError(c, err)
	}
	share.RespondDate(c, http.StatusOK, data)
}
