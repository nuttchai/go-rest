package interfaces

import (
	"database/sql"

	dto "github.com/nuttchai/go-rest/internal/dto/sample"
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/types"
)

type ISampleRepository interface {
	Test() string
	RetrieveOne(id string, filters ...*types.TQueryFilter) (*models.Sample, error)
	CreateOne(sample *dto.CreateSampleDTO) (*models.Sample, error)
	UpdateOne(sample *dto.UpdateSampleDTO) (*models.Sample, error)
	DeleteOne(id string) (sql.Result, error)
}
