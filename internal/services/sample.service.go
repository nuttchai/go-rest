package services

import (
	dto "github.com/nuttchai/go-rest/internal/dto/sample"
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/repositories"
	"github.com/nuttchai/go-rest/internal/utils/validators"
)

type TSampleService struct {
	repository repositories.ISampleRepository
}

var (
	SampleService ISampleService
)

func (s *TSampleService) Test() string {
	return s.repository.Test()
}

func (s *TSampleService) GetSample(id string) (*models.Sample, error) {
	return s.repository.GetSample(id)
}

func (s *TSampleService) CreateSample(sample *dto.CreateSampleDTO) (*models.Sample, error) {
	return s.repository.CreateSample(sample)
}

func (s *TSampleService) UpdateSample(sample *dto.UpdateSampleDTO) (*models.Sample, error) {
	return s.repository.UpdateSample(sample)
}

func (s *TSampleService) DeleteSample(id string) error {
	result, err := s.repository.DeleteSample(id)
	if err != nil {
		return err
	}

	return validators.CheckRowsAffected(result)
}
