package repositories

import (
	"database/sql"

	dto "github.com/nuttchai/go-rest/internal/dto/sample"
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/repositories/interfaces"
	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/types"
	"github.com/nuttchai/go-rest/internal/utils/context"
	"github.com/nuttchai/go-rest/internal/utils/db"
)

type TSampleRepository struct {
	DB *sql.DB
}

var (
	SampleRepository interfaces.ISampleRepository
)

func (m *TSampleRepository) Test() string {
	console.App.Log("Call Test Function in Repository!")
	return "test"
}

func (m *TSampleRepository) GetSample(id string, filters ...*types.TQueryFilter) (*models.Sample, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	baseQuery := "select * from sample where id = $1"
	baseArgs := []interface{}{id}

	query, args := db.BuildQueryWithFilter(baseQuery, baseArgs, filters...)
	row := m.DB.QueryRowContext(ctx, query, args...)

	var sample models.Sample
	err := row.Scan(
		&sample.Id,
		&sample.Name,
		&sample.Description,
		&sample.OwnerId,
	)

	return &sample, err
}

func (m *TSampleRepository) CreateSample(sample *dto.CreateSampleDTO) (*models.Sample, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		insert into sample (name, description, owner_id)
		values ($1, $2, $3)
		returning *
	`
	row := m.DB.QueryRowContext(ctx, query, sample.Name, sample.Description, sample.OwnerId)

	var newSample models.Sample
	if err := row.Scan(
		&newSample.Id,
		&newSample.Name,
		&newSample.Description,
		&newSample.OwnerId,
	); err != nil {
		return nil, err
	}

	return &newSample, nil
}

func (m *TSampleRepository) UpdateSample(sample *dto.UpdateSampleDTO) (*models.Sample, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		update sample set name = $1, description = $2
		where id = $3
		returning *
	`
	row := m.DB.QueryRowContext(ctx, query, sample.Name, sample.Description, sample.Id)

	var updatedSample models.Sample
	if err := row.Scan(
		&updatedSample.Id,
		&updatedSample.Name,
		&updatedSample.Description,
		&updatedSample.OwnerId,
	); err != nil {
		return nil, err
	}

	return &updatedSample, nil
}

func (m *TSampleRepository) DeleteSample(id string) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		delete from sample 
		where id = $1
	`
	result, err := m.DB.ExecContext(ctx, query, id)

	return result, err
}
