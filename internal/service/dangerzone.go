package service

import "hte-danger-zone-ms/internal/repository"

type DangerZoneService interface {
	Create()
}
type dangerZoneService struct {
	repo repository.DangerZoneRepository
}

func NewDangerZoneService(repo repository.DangerZoneRepository) DangerZoneService {
	return &dangerZoneService{repo: repo}
}

func (svc *dangerZoneService) Create() {

}
