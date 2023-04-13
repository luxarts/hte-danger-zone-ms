package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hte-danger-zone-ms/internal/domain"
	"hte-danger-zone-ms/internal/service"
	"net/http"
)

type DangerZoneController interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	GetAll(c *gin.Context)
	GetByDeviceID(c *gin.Context)
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
			"error":   err.Error(),
		})
		return
	}

	err := ctrl.svc.Create(&body)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed Create",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, body)
}

func (ctrl *dangerZoneController) Delete(ctx *gin.Context) {
	deviceID := ctx.Param("deviceID")
	if deviceID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid deviceID",
		})
		return
	}
	err := ctrl.svc.Delete(deviceID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Error",
		})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Dangerzone eliminated",
	})
}

func (ctrl *dangerZoneController) GetAll(ctx *gin.Context) {
	dangerZonesBson, err := ctrl.svc.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Error",
		})
		return
	}
	dangerZonesJson, err := json.Marshal(dangerZonesBson)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Error",
		})
		return
	}
	ctx.JSON(http.StatusOK, dangerZonesJson)
}

func (ctrl *dangerZoneController) GetByDeviceID(ctx *gin.Context) {
	deviceID, _ := ctx.GetQuery("device_id")
	if deviceID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid deviceID",
		})
		return
	}
	dangerZoneBson, err := ctrl.svc.GetByDeviceID(deviceID)
	dangerZoneJson, err := json.Marshal(dangerZoneBson)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Error",
		})
		return
	}
	ctx.JSON(http.StatusOK, dangerZoneJson)
}
