package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"hte-danger-zone-ms/internal/domain"
	"log"
)

type DangerZoneRepository interface {
	Create(body *domain.DangerZone) error
	Delete(deviceID string) error
	GetAll(filter map[string]string) (*[]domain.DangerZone, error)
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

func (repo *dangerZoneRepository) Delete(deviceID string) error {
	ctx := context.Background()
	_, err := repo.db.Collection(repo.collection).DeleteOne(ctx, bson.D{{"device_id", deviceID}})
	if err != nil {
		return err
	}
	return nil
}

func (repo *dangerZoneRepository) GetAll(filter map[string]string) (*[]domain.DangerZone, error) {
	ctx := context.Background()
	var resp []domain.DangerZone
	dgBson, err := repo.db.Collection(repo.collection).Find(ctx, filter)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	for dgBson.Next(ctx) {
		var dangerZone domain.DangerZone
		err = dgBson.Decode(&dangerZone)
		if err != nil {
			log.Println("Error unmarshal danger zone")
			continue
		}
		resp = append(resp, dangerZone)
	}
	return &resp, nil
}
