package repository

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"hte-danger-zone-ms/internal/domain"
)

type DangerZoneEventRepository interface {
	Create(body *domain.DangerZone) error
	Delete(deviceID string) error
}

type dangerZoneEventRepository struct {
	channelCreateDZ string
	channelDeleteDZ string
	rc              *redis.Client
}

func NewDangerZoneEventRepository(rc *redis.Client, channelCreateDZ string, channelDeleteDZ string) DangerZoneEventRepository {
	return &dangerZoneEventRepository{rc: rc, channelCreateDZ: channelCreateDZ, channelDeleteDZ: channelDeleteDZ}
}

func (repo *dangerZoneEventRepository) Create(body *domain.DangerZone) error {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	ctx := context.Background()
	return repo.rc.Publish(ctx, repo.channelCreateDZ, bodyBytes).Err()
}

func (repo *dangerZoneEventRepository) Delete(deviceID string) error {
	ctx := context.Background()
	return repo.rc.Publish(ctx, repo.channelDeleteDZ, deviceID).Err()
}
