package repositories

import (
	"database/sql"

	sampledto "github.com/nuttchai/go-rest/internal/dto/sample"
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/types"
	"github.com/nuttchai/go-rest/internal/utils/context"
	"github.com/nuttchai/go-rest/internal/utils/db"
)

func (m *DBModel) GetSample(id string, filters ...*types.QueryFilter) (*models.Sample, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	baseQuery := "select * from samples where id = $1"
	baseArgs := []interface{}{id}

	query, args := db.BuildQueryWithFilter(baseQuery, baseArgs, filters...)
	row := m.SqlDB.QueryRowContext(ctx, query, args...)

	var sample models.Sample
	err := row.Scan(
		&sample.Id,
		&sample.Name,
	)

	return &sample, err
}

func (m *DBModel) CreateSample(s *sampledto.CreateSampleDTO) (*models.Sample, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		insert into samples (name, desc)
		values ($1, $2)
		returning *
	`
	row := m.SqlDB.QueryRowContext(ctx, query, s.Name, s.Desc)

	var sample models.Sample
	if err := row.Scan(
		&sample.Id,
		&sample.Name,
		&sample.Desc,
	); err != nil {
		return nil, err
	}

	return &sample, nil
}

func (m *DBModel) UpdateSample(s *sampledto.UpdateSampleDTO) (*models.Sample, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		update samples set name = $1, desc = $2
		where id = $3
		returning *
	`
	row := m.SqlDB.QueryRowContext(ctx, query, s.Name, s.Desc, s.Id)

	var sample models.Sample
	if err := row.Scan(
		&sample.Id,
		&sample.Name,
		&sample.Desc,
	); err != nil {
		return nil, err
	}

	return &sample, nil
}

func (m *DBModel) DeleteSample(id string) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		delete from samples 
		where id = $1
	`
	result, err := m.SqlDB.ExecContext(ctx, query, id)

	return result, err
}
