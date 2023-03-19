package interfaces

import (
	sampledto "github.com/nuttchai/go-rest/internal/dto/sample"
	"github.com/nuttchai/go-rest/internal/models"
)

type ISampleService interface {
	Test() string
	GetSample(id string) (*models.Sample, error)
	CreateSample(sample *sampledto.CreateSampleDTO) (*models.Sample, error)
	UpdateSample(sample *sampledto.UpdateSampleDTO) (*models.Sample, error)
	DeleteSample(id string) error
}