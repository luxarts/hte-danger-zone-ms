package repository

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"hte-danger-zone-ms/internal/defines"
	"hte-danger-zone-ms/internal/domain"
)

const (
	postgresErrNotUnique = "23505"
	schema               = "core"
	table                = "danger_zones"
)

type DangerZoneRepository interface {
	Create(body *domain.DangerZone) error
	Delete(deviceID string) error
	GetAll() (*[]domain.DangerZone, error)
	GetAllByCompanyID(companyID string) (*[]domain.DangerZone, error)
	GetByDeviceID(deviceID string) (*domain.DangerZone, error)
}

type dangerZoneRepository struct {
	db         *sqlx.DB
	sqlBuilder *tableDz
}

func NewDangerZoneRepository(db *sqlx.DB) DangerZoneRepository {
	return &dangerZoneRepository{
		db: db,
		sqlBuilder: &tableDz{
			table: schema + "." + table,
		},
	}
}

func (repo *dangerZoneRepository) Create(dz *domain.DangerZone) error {
	query, args, err := repo.sqlBuilder.CreateSQL(dz)
	if err != nil {
		return err
	}
	_, err = repo.db.Exec(query, args...)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code == postgresErrNotUnique {
				return defines.ErrZoneExists
			}
		}
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
func (repo *dangerZoneRepository) GetAll() (*[]domain.DangerZone, error) {
	var dzs []domain.DangerZone
	query, args, err := repo.sqlBuilder.GetAllSQL()
	if err != nil {
		return nil, err
	}
	rows, err := repo.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var dz domain.DangerZone
		err = rows.Err()
		if err != nil {
			return nil, err
		}
		err = rows.StructScan(&dz)
		if err != nil {
			return nil, err
		}
		dzs = append(dzs, dz)
	}
	return &dzs, err
}
func (repo *dangerZoneRepository) GetAllByCompanyID(companyID string) (*[]domain.DangerZone, error) {
	var dzs []domain.DangerZone
	query, args, err := repo.sqlBuilder.GetAllByCompanyIDSQL(companyID)
	if err != nil {
		return nil, err
	}
	rows, err := repo.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var dz domain.DangerZone
		err = rows.Err()
		if err != nil {
			return nil, err
		}
		err = rows.StructScan(&dz)
		if err != nil {
			return nil, err
		}
		dzs = append(dzs, dz)
	}
	return &dzs, err
}
func (repo *dangerZoneRepository) GetByDeviceID(deviceID string) (*domain.DangerZone, error) {
	var dz domain.DangerZone
	query, args, err := repo.sqlBuilder.GetByDeviceIDSQL(deviceID)
	if err != nil {
		return nil, err
	}
	row := repo.db.QueryRowx(query, args...)
	err = row.Err()
	if err != nil {
		return nil, err
	}
	err = row.StructScan(&dz)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &dz, nil
}

type tableDz struct {
	table string
}

func (t *tableDz) CreateSQL(dz *domain.DangerZone) (string, []interface{}, error) {
	query, args, err := squirrel.Insert(t.table).
		Columns("device_id", "company_id", "longitude", "latitude", "radius", "end_ts").
		Values(dz.DeviceID, dz.CompanyID, dz.Longitude, dz.Latitude, dz.Radius, dz.EndTs).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	return query, args, err
}
func (t *tableDz) DeleteSQL(deviceID string) (string, []interface{}, error) {
	query, args, err := squirrel.Delete(t.table).
		Where(squirrel.Eq{"device_id": deviceID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	return query, args, err
}
func (t *tableDz) GetAllSQL() (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").
		From(t.table).
		ToSql()
	return query, args, err
}
func (t *tableDz) GetAllByCompanyIDSQL(companyID string) (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").
		From(t.table).
		Where(squirrel.Eq{"company_id": companyID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	return query, args, err
}
func (t *tableDz) GetByDeviceIDSQL(deviceID string) (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").
		From(t.table).
		Where(squirrel.Eq{"device_id": deviceID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	return query, args, err
}
