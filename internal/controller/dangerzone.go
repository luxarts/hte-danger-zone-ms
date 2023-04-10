package controller

import (
	"github.com/gin-gonic/gin"
	"hte-danger-zone-ms/internal/domain"
	"hte-danger-zone-ms/internal/service"
	"net/http"
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

func (ctrl *dangerZoneController) Create(ctx *gin.Context) {
	var body domain.DangerZoneCreateReq

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed ShouldBindJSON",
			"error":   err,
		})
		return
	}

	err := ctrl.svc.Create(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed Create",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, body)
}
