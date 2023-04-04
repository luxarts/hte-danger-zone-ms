package router

import (
	"github.com/gin-gonic/gin"
	"hte-danger-zone-ms/internal/controller"
	"hte-danger-zone-ms/internal/defines"
	"hte-danger-zone-ms/internal/repository"
	"hte-danger-zone-ms/internal/service"
)

func New() *gin.Engine {
	r := gin.Default()

	mapRoutes(r)

	return r
}

func mapRoutes(r *gin.Engine) {
	// Init clients

	// Init repositories
	repo := repository.NewDangerZoneRepository()

	// Init services
	svc := service.NewDangerZoneService(repo)

	// Init controllers
	ctrl := controller.NewDangerZoneController(svc)

	// Routes
	r.POST(defines.EndpointCreateDangerZone, ctrl.Create)
}
