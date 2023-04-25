package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"hte-danger-zone-ms/internal/defines"
	"hte-danger-zone-ms/internal/domain"
	"hte-danger-zone-ms/internal/service"
	"hte-danger-zone-ms/metrics"
	"net/http"
	"strconv"
)

type DangerZoneController interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	GetAll(c *gin.Context)
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
		metrics.HTTPResponseCounter.
			With(prometheus.Labels{
				metrics.LabelStatusCode: strconv.Itoa(http.StatusBadRequest),
				metrics.LabelMethod:     "POST",
				metrics.LabelEndpoint:   defines.EndpointCreateDangerZone,
			}).
			Inc()
		return
	}

	if !body.IsValid() {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid body",
		})
		metrics.HTTPResponseCounter.
			With(prometheus.Labels{
				metrics.LabelStatusCode: strconv.Itoa(http.StatusBadRequest),
				metrics.LabelMethod:     "POST",
				metrics.LabelEndpoint:   defines.EndpointCreateDangerZone,
			}).
			Inc()
		return
	}

	err := ctrl.svc.Create(&body)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed Create",
			"error":   err.Error(),
		})
		metrics.HTTPResponseCounter.
			With(prometheus.Labels{
				metrics.LabelStatusCode: strconv.Itoa(http.StatusInternalServerError),
				metrics.LabelMethod:     "POST",
				metrics.LabelEndpoint:   defines.EndpointCreateDangerZone,
			}).
			Inc()
		return
	}

	ctx.JSON(http.StatusCreated, body)
	metrics.HTTPResponseCounter.
		With(prometheus.Labels{
			metrics.LabelStatusCode: strconv.Itoa(http.StatusCreated),
			metrics.LabelMethod:     "POST",
			metrics.LabelEndpoint:   defines.EndpointCreateDangerZone,
		}).
		Inc()
}

func (ctrl *dangerZoneController) Delete(ctx *gin.Context) {
	deviceID := ctx.Query(defines.QueryParamDeviceID)
	if deviceID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid device_id",
		})
		metrics.HTTPResponseCounter.
			With(prometheus.Labels{
				metrics.LabelStatusCode: strconv.Itoa(http.StatusBadRequest),
				metrics.LabelMethod:     "DELETE",
				metrics.LabelEndpoint:   defines.EndpointDeleteDangerZone,
			}).
			Inc()
		return
	}
	err := ctrl.svc.Delete(deviceID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Error",
		})
		metrics.HTTPResponseCounter.
			With(prometheus.Labels{
				metrics.LabelStatusCode: strconv.Itoa(http.StatusInternalServerError),
				metrics.LabelMethod:     "DELETE",
				metrics.LabelEndpoint:   defines.EndpointDeleteDangerZone,
			}).
			Inc()
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Dangerzone eliminated",
	})
	metrics.HTTPResponseCounter.
		With(prometheus.Labels{
			metrics.LabelStatusCode: strconv.Itoa(http.StatusNoContent),
			metrics.LabelMethod:     "DELETE",
			metrics.LabelEndpoint:   defines.EndpointDeleteDangerZone,
		}).
		Inc()
}

func (ctrl *dangerZoneController) GetAll(ctx *gin.Context) {
	deviceID := ctx.Query(defines.QueryParamDeviceID)
	if deviceID != "" {
		dangerZones, err := ctrl.svc.GetByDeviceID(deviceID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Error"})
			metrics.HTTPResponseCounter.
				With(prometheus.Labels{
					metrics.LabelStatusCode: strconv.Itoa(http.StatusInternalServerError),
					metrics.LabelMethod:     "GET",
					metrics.LabelEndpoint:   defines.EndpointGetAllDangerZone,
				}).
				Inc()
			return
		}
		if dangerZones == nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "danger zone not found for the given device_id"})
			metrics.HTTPResponseCounter.
				With(prometheus.Labels{
					metrics.LabelStatusCode: strconv.Itoa(http.StatusNotFound),
					metrics.LabelMethod:     "GET",
					metrics.LabelEndpoint:   defines.EndpointGetAllDangerZone,
				}).
				Inc()
			return
		}
		ctx.JSON(http.StatusOK, dangerZones)
		metrics.HTTPResponseCounter.
			With(prometheus.Labels{
				metrics.LabelStatusCode: strconv.Itoa(http.StatusOK),
				metrics.LabelMethod:     "GET",
				metrics.LabelEndpoint:   defines.EndpointGetAllDangerZone,
			}).
			Inc()
		return
	}

	companyID := ctx.Query(defines.QueryParamCompanyID)
	if companyID != "" {
		dangerZones, err := ctrl.svc.GetAllByCompanyID(companyID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Error"})
			metrics.HTTPResponseCounter.
				With(prometheus.Labels{
					metrics.LabelStatusCode: strconv.Itoa(http.StatusInternalServerError),
					metrics.LabelMethod:     "GET",
					metrics.LabelEndpoint:   defines.EndpointGetAllDangerZone,
				}).
				Inc()
			return
		}
		ctx.JSON(http.StatusOK, dangerZones)
		metrics.HTTPResponseCounter.
			With(prometheus.Labels{
				metrics.LabelStatusCode: strconv.Itoa(http.StatusOK),
				metrics.LabelMethod:     "GET",
				metrics.LabelEndpoint:   defines.EndpointGetAllDangerZone,
			}).
			Inc()
		return
	}

	dangerZones, err := ctrl.svc.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Error"})
		metrics.HTTPResponseCounter.
			With(prometheus.Labels{
				metrics.LabelStatusCode: strconv.Itoa(http.StatusInternalServerError),
				metrics.LabelMethod:     "GET",
				metrics.LabelEndpoint:   defines.EndpointGetAllDangerZone,
			}).
			Inc()
		return
	}
	ctx.JSON(http.StatusOK, dangerZones)
	metrics.HTTPResponseCounter.
		With(prometheus.Labels{
			metrics.LabelStatusCode: strconv.Itoa(http.StatusOK),
			metrics.LabelMethod:     "GET",
			metrics.LabelEndpoint:   defines.EndpointGetAllDangerZone,
		}).
		Inc()
}
