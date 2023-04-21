package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"hte-danger-zone-ms/internal/defines"
	"hte-danger-zone-ms/internal/domain"
)

type DangerZoneRepository interface {
	Create(body *domain.DangerZone) error
	Delete(deviceID string) error
	GetAll(filter map[string]string) (*[]domain.DangerZone, error)
}

type dangerZoneRepository struct {
	db         *sqlx.DB
	sqlBuilder *tableDz
}

func NewDangerZoneRepository(db *sqlx.DB) DangerZoneRepository {
	return &dangerZoneRepository{
		db:         db,
		sqlBuilder: &tableDz{table: defines.TableDangerZone},
	}
}

func (repo *dangerZoneRepository) Create(dz *domain.DangerZone) error {
	query, args, err := repo.sqlBuilder.CreateSQL(dz)
	if err != nil {
		return err
	}
	_, err = repo.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (repo *dangerZoneRepository) Delete(deviceID string) error {
	query, args, err := repo.sqlBuilder.DeleteSQL(deviceID)
	if err != nil {
		return err
	}
	_, err = repo.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (repo *dangerZoneRepository) GetAll(filter map[string]string) (*[]domain.DangerZone, error) {

	return nil, nil
}

type tableDz struct {
	table string
}

func (t *tableDz) CreateSQL(dz *domain.DangerZone) (string, []interface{}, error) {
	query, args, err := squirrel.Insert(t.table).
		Columns("device_id", "longitude", "latitude", "radius", "end_ts").
		Values(dz.DeviceID, dz.Longitude, dz.Latitude, dz.Radius, dz.EndTs).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	return query, args, err
}

func (t *tableDz) DeleteSQL(deviceID string) (string, []interface{}, error) {
	query, args, err := squirrel.Delete(t.table).
		Where(squirrel.Eq{"device_id": deviceID}).
		ToSql()
	return query, args, err
}
