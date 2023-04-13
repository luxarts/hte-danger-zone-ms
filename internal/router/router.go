package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s",
			os.Getenv(defines.EnvMongoHost),
		)))
	if err != nil {
		log.Fatalln(err)
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Error ping MongoDB %+v\n", err)
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
	repo := repository.NewDangerZoneRepository(mongoClient, os.Getenv(defines.EnvMongoDB), os.Getenv(defines.EnvDangerZonesCollection))
	dzeRepo := repository.NewDangerZoneEventRepository(redisClient, os.Getenv(defines.EnvRedisChannelCreateDangerZone), os.Getenv(defines.EnvRedisChannelDeleteDangerZone))

	// Init services
	svc := service.NewDangerZoneService(repo, dzeRepo)

	// Init controllers
	ctrl := controller.NewDangerZoneController(svc)

	// Routes
	r.POST(defines.EndpointCreateDangerZone, ctrl.Create)
	r.DELETE(defines.EndpointDeleteDangerZone, ctrl.Delete)
}
