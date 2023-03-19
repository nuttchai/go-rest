package iservice

import (
	sampledto "github.com/nuttchai/go-rest/internal/dto/sample"
	"github.com/nuttchai/go-rest/internal/model"
)

type ISampleService interface {
	Test() string
	GetSample(id string) (*model.Sample, error)
	CreateSample(sample *sampledto.CreateSampleDTO) (*model.Sample, error)
	UpdateSample(sample *sampledto.UpdateSampleDTO) (*model.Sample, error)
	DeleteSample(id string) error
}
