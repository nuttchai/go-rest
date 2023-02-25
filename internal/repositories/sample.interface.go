package repositories

import (
	"database/sql"

	dto "github.com/nuttchai/go-rest/internal/dto/sample"
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/types"
)

type ISampleRepository interface {
	Test() string
	GetSample(id string, filters ...*types.QueryFilter) (*models.Sample, error)
	CreateSample(sample *dto.CreateSampleDTO) (*models.Sample, error)
	UpdateSample(sample *dto.UpdateSampleDTO) (*models.Sample, error)
	DeleteSample(id string) (sql.Result, error)
}
