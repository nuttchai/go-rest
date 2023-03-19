package service

import (
	dto "github.com/nuttchai/go-rest/internal/dto/sample"
	"github.com/nuttchai/go-rest/internal/model"
	irepository "github.com/nuttchai/go-rest/internal/repository/interface"
	iservice "github.com/nuttchai/go-rest/internal/service/interface"
	"github.com/nuttchai/go-rest/internal/util/validators"
)

type TSampleService struct {
	repository irepository.ISampleRepository
}

var (
	SampleService iservice.ISampleService
)

func (s *TSampleService) Test() string {
	return s.repository.Test()
}

func (s *TSampleService) GetSample(id string) (*model.Sample, error) {
	return s.repository.RetrieveOne(id)
}

func (s *TSampleService) CreateSample(sample *dto.CreateSampleDTO) (*model.Sample, error) {
	return s.repository.CreateOne(sample)
}

func (s *TSampleService) UpdateSample(sample *dto.UpdateSampleDTO) (*model.Sample, error) {
	return s.repository.UpdateOne(sample)
}

func (s *TSampleService) DeleteSample(id string) error {
	result, err := s.repository.DeleteOne(id)
	if err != nil {
		return err
	}

	return validators.CheckRowsAffected(result)
}
