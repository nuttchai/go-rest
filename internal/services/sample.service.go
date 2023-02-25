package services

import (
	dto "github.com/nuttchai/go-rest/internal/dto/sample"
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/repositories"
	"github.com/nuttchai/go-rest/internal/utils/validators"
)

type sampleService struct {
	repo *Repository
}

var (
	SampleService ISampleService
)

func init() {
	SampleService = &sampleService{
		repo: &repo,
	}
}

func (s *sampleService) Test() string {
	return repositories.SampleRepository.Test()
}

func (s *sampleService) GetSample(id string) (*models.Sample, error) {
	return repositories.SampleRepository.GetSample(id)
}

func (s *sampleService) CreateSample(sample *dto.CreateSampleDTO) (*models.Sample, error) {
	return repositories.SampleRepository.CreateSample(sample)
}

func (s *sampleService) UpdateSample(sample *dto.UpdateSampleDTO) (*models.Sample, error) {
	return repositories.SampleRepository.UpdateSample(sample)
}

func (s *sampleService) DeleteSample(id string) error {
	result, err := repositories.SampleRepository.DeleteSample(id)
	if err != nil {
		return err
	}

	return validators.CheckRowsAffected(result)
}
