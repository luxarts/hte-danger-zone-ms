package repository

type DangerZoneRepository interface {
	Create()
}

type dangerZoneRepository struct {
}

func NewDangerZoneRepository() DangerZoneRepository {
	return &dangerZoneRepository{}
}

func (repo *dangerZoneRepository) Create() {

}
