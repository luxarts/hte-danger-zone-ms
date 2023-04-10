package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"hte-danger-zone-ms/internal/domain"
)

type DangerZoneRepository interface {
	Create(body *domain.DangerZoneCreateReq) error
}

type dangerZoneRepository struct {
	db         *mongo.Database
	collection string
}

func NewDangerZoneRepository(mc *mongo.Client, database string, collection string) DangerZoneRepository {
	return &dangerZoneRepository{
		db:         mc.Database(database),
		collection: collection,
	}
}

func (repo *dangerZoneRepository) Create(body *domain.DangerZoneCreateReq) error {
	bsonBody, err := bson.Marshal(body)
	if err != nil {
		return err
	}

	ctx := context.Background()
	_, err = repo.db.Collection(repo.collection).InsertOne(ctx, bsonBody)
	return err
}
