package repository

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"hte-danger-zone-ms/internal/domain"
)

type DangerZoneEventRepository interface {
	Create(body *domain.DangerZone) error
}

type dangerZoneEventRepository struct {
	channel string
	rc      *redis.Client
}

func NewDangerZoneEventRepository(rc *redis.Client, channel string) DangerZoneEventRepository {
	return &dangerZoneEventRepository{rc: rc, channel: channel}
}

func (repo *dangerZoneEventRepository) Create(body *domain.DangerZone) error {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	ctx := context.Background()
	return repo.rc.Publish(ctx, repo.channel, bodyBytes).Err()
}
