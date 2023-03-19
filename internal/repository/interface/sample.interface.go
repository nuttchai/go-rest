package irepository

import (
	"database/sql"

	dto "github.com/nuttchai/go-rest/internal/dto/sample"
	"github.com/nuttchai/go-rest/internal/model"
	"github.com/nuttchai/go-rest/internal/types"
)

type ISampleRepository interface {
	Test() string
	RetrieveOne(id string, filters ...*types.TQueryFilter) (*model.Sample, error)
	CreateOne(sample *dto.CreateSampleDTO) (*model.Sample, error)
	UpdateOne(sample *dto.UpdateSampleDTO) (*model.Sample, error)
	DeleteOne(id string) (sql.Result, error)
}
