package service

import (
	"hte-danger-zone-ms/internal/domain"
	"hte-danger-zone-ms/internal/repository"
)

type DangerZoneService interface {
	Create(body *domain.DangerZoneCreateReq) error
}
type dangerZoneService struct {
	repo      repository.DangerZoneRepository
	eventRepo repository.DangerZoneEventRepository
}

func NewDangerZoneService(repo repository.DangerZoneRepository, eventRepo repository.DangerZoneEventRepository) DangerZoneService {
	return &dangerZoneService{
		repo:      repo,
		eventRepo: eventRepo,
	}
}

func (svc *dangerZoneService) Create(body *domain.DangerZoneCreateReq) error {
	err := svc.repo.Create(body)
	if err != nil {
		return err
	}

	return svc.eventRepo.Create(body.ToDangerZone())
}
