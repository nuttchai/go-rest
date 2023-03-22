package repository

import (
	"database/sql"

	"github.com/nuttchai/go-rest/internal/constant"
	dto "github.com/nuttchai/go-rest/internal/dto/sample"
	"github.com/nuttchai/go-rest/internal/model"
	irepository "github.com/nuttchai/go-rest/internal/repository/interface"
	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/types"
	"github.com/nuttchai/go-rest/internal/util/context"
	"github.com/nuttchai/go-rest/internal/util/db"
)

type TSampleRepository struct {
	DB *sql.DB
}

var (
	SampleRepository irepository.ISampleRepository
)

func initSampleRepository(sampleRepository *TSampleRepository) {
	SampleRepository = sampleRepository
}

func (m *TSampleRepository) Test() string {
	console.App.Log("Call Test Function in Repository!")
	return "test"
}

func (m *TSampleRepository) RetrieveOne(id string, filters ...*types.TQueryFilter) (*model.Sample, error) {
	ctx, cancel := context.WithTimeout(constant.QueryTimeout)
	defer cancel()

	baseQuery := "select * from sample where id = $1"
	baseArgs := []interface{}{id}

	query, args := db.BuildQueryWithFilter(baseQuery, baseArgs, filters...)
	row := m.DB.QueryRowContext(ctx, query, args...)

	var sample model.Sample
	err := row.Scan(
		&sample.Id,
		&sample.Name,
		&sample.Description,
		&sample.OwnerId,
	)

	return &sample, err
}

func (m *TSampleRepository) CreateOne(sample *dto.CreateSampleDTO) (*model.Sample, error) {
	ctx, cancel := context.WithTimeout(constant.QueryTimeout)
	defer cancel()

	query := `
		insert into sample (name, description, owner_id)
		values ($1, $2, $3)
		returning *
	`
	row := m.DB.QueryRowContext(ctx, query, sample.Name, sample.Description, sample.OwnerId)

	var newSample model.Sample
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

func (m *TSampleRepository) UpdateOne(sample *dto.UpdateSampleDTO) (*model.Sample, error) {
	ctx, cancel := context.WithTimeout(constant.QueryTimeout)
	defer cancel()

	query := `
		update sample set name = $1, description = $2
		where id = $3
		returning *
	`
	row := m.DB.QueryRowContext(ctx, query, sample.Name, sample.Description, sample.Id)

	var updatedSample model.Sample
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

func (m *TSampleRepository) DeleteOne(id string) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(constant.QueryTimeout)
	defer cancel()

	query := `
		delete from sample 
		where id = $1
	`
	result, err := m.DB.ExecContext(ctx, query, id)

	return result, err
}
