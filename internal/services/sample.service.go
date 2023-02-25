package services

import (
	sampledto "github.com/nuttchai/go-rest/internal/dto/sample"
	"github.com/nuttchai/go-rest/internal/models"
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

func (s *sampleService) GetSample(id string) (*models.Sample, error) {
	return s.repo.Models.DB.GetSample(id)
}

func (s *sampleService) CreateSample(sample *sampledto.CreateSampleDTO) (*models.Sample, error) {
	return s.repo.Models.DB.CreateSample(sample)
}

func (s *sampleService) UpdateSample(sample *sampledto.UpdateSampleDTO) (*models.Sample, error) {
	return s.repo.Models.DB.UpdateSample(sample)
}

func (s *sampleService) DeleteSample(id string) error {
	result, err := s.repo.Models.DB.DeleteSample(id)
	if err != nil {
		return err
	}

	return validators.CheckRowsAffected(result)
}
