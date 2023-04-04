package controller

import (
	"github.com/gin-gonic/gin"
	"hte-danger-zone-ms/internal/service"
)

type DangerZoneController interface {
	Create(c *gin.Context)
}

type dangerZoneController struct {
	svc service.DangerZoneService
}

func NewDangerZoneController(svc service.DangerZoneService) DangerZoneController {
	return &dangerZoneController{svc: svc}
}

func (ctrl *dangerZoneController) Create(c *gin.Context) {

}
