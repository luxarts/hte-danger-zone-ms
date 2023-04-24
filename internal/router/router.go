package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"hte-danger-zone-ms/internal/controller"
	"hte-danger-zone-ms/internal/defines"
	"hte-danger-zone-ms/internal/repository"
	"hte-danger-zone-ms/internal/service"
	"log"
	"os"
)

func New() *gin.Engine {
	r := gin.Default()

	mapRoutes(r)

	return r
}

func mapRoutes(r *gin.Engine) {
	// Init clients
	ctx := context.Background()
	postgresURI := fmt.Sprintf("postgres://%s:%s@%s/postgres?sslmode=disable",
		os.Getenv(defines.EnvPostgresUser),
		os.Getenv(defines.EnvPostgresPassword),
		os.Getenv(defines.EnvPostgresHost))
	db, err := sqlx.Open("postgres", postgresURI)
	if err != nil {
		log.Println(postgresURI)
		log.Panicln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error ping Postgres: %+v\n", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(defines.EnvRedisHost),
		Password: os.Getenv(defines.EnvRedisPassword),
	})

	err = redisClient.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("Error ping Redis: %+v\n", err)
	}

	// Init repositories
	repo := repository.NewDangerZoneRepository(db, os.Getenv(defines.EnvPostgresSchema), os.Getenv(defines.EnvPostgresDangerZonesTable))
	dzeRepo := repository.NewDangerZoneEventRepository(redisClient, os.Getenv(defines.EnvRedisChannelCreateDangerZone), os.Getenv(defines.EnvRedisChannelDeleteDangerZone))

	// Init services
	svc := service.NewDangerZoneService(repo, dzeRepo)

	// Init controllers
	ctrl := controller.NewDangerZoneController(svc)

	// Routes
	r.GET(defines.EndpointGetAllDangerZone, ctrl.GetAll)
	r.GET(defines.EndpointGetAllDangerZoneByCompanyID, ctrl.GetAll)
	r.POST(defines.EndpointCreateDangerZone, ctrl.Create)
	r.DELETE(defines.EndpointDeleteDangerZone, ctrl.Delete)
}
