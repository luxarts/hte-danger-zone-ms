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
	GetAllByCompanyID(companyID string) (*[]domain.DangerZone, error)
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
	getDz, err := svc.repo.GetByDeviceID(body.DeviceID)
	if err != nil {
		return err
	}
	if getDz != nil {
		return defines.ErrZoneExists
	}
	dzCreate := body.ToDangerZone()
	dzCreate.EndTs = time.Now().UTC().Add(time.Duration(body.TTL) * time.Second).Unix()
	err = svc.repo.Create(dzCreate)
	if err != nil {
		return err
	}
	return svc.eventRepo.Create(dzCreate)
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
func (svc *dangerZoneService) GetAllByCompanyID(companyID string) (*[]domain.DangerZone, error) {
	return svc.repo.GetAllByCompanyID(companyID)
}
func (svc *dangerZoneService) GetByDeviceID(deviceID string) (*domain.DangerZone, error) {
	return svc.repo.GetByDeviceID(deviceID)
}
