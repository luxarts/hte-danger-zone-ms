package service

import (
	"hte-danger-zone-ms/internal/defines"
	"hte-danger-zone-ms/internal/domain"
	"hte-danger-zone-ms/internal/repository"
	"time"
)

type DangerZoneService interface {
	Create(body *domain.DangerZoneCreateReq) error
	Delete(deviceID string) error
	GetAll() (*[]domain.DangerZone, error)
	GetByDeviceID(deviceID string) (*domain.DangerZone, error)
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
	dz, err := svc.repo.GetByDeviceID(body.DeviceID)
	if err != nil {
		return err
	}
	if dz != nil {
		return defines.ErrZoneExists
	}

	dz = body.ToDangerZone()
	dz.EndTs = time.Now().UTC().Add(time.Duration(body.TTL) * time.Second).Unix()

	err = svc.repo.Create(dz)
	if err != nil {
		return err
	}

	return svc.eventRepo.Create(dz)
}

func (svc *dangerZoneService) Delete(deviceID string) error {
	err := svc.repo.Delete(deviceID)
	if err != nil {
		return err
	}
	return svc.eventRepo.Delete(deviceID)
}

func (svc *dangerZoneService) GetAll() (*[]domain.DangerZone, error) {
	return svc.repo.GetAll()
}

func (svc *dangerZoneService) GetByDeviceID(deviceID string) (*domain.DangerZone, error) {
	return svc.repo.GetByDeviceID(deviceID)
}
