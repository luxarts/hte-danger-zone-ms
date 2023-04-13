package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"hte-danger-zone-ms/internal/domain"
)

type DangerZoneRepository interface {
	Create(body *domain.DangerZone) error
	GetByDeviceID(deviceID string) (*domain.DangerZone, error)
	Delete(deviceID string) error
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

func (repo *dangerZoneRepository) Create(body *domain.DangerZone) error {
	bsonBody, err := bson.Marshal(body)
	if err != nil {
		return err
	}

	ctx := context.Background()
	_, err = repo.db.Collection(repo.collection).InsertOne(ctx, bsonBody)
	return err
}
func (repo *dangerZoneRepository) GetByDeviceID(deviceID string) (*domain.DangerZone, error) {
	ctx := context.Background()

	var resp domain.DangerZone

	err := repo.db.Collection(repo.collection).FindOne(ctx, bson.D{
		{"device_id", deviceID},
	}).Decode(&resp)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (repo *dangerZoneRepository) Delete(deviceID string) error {
	ctx := context.Background()
	_, err := repo.db.Collection(repo.collection).DeleteOne(ctx, bson.D{{"device_id", deviceID}})
	if err != nil {
		return err
	}
	return nil
}
