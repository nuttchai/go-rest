package service

import (
	dto "github.com/nuttchai/go-rest/internal/dto/sample"
	"github.com/nuttchai/go-rest/internal/model"
	irepository "github.com/nuttchai/go-rest/internal/repository/interface"
	iservice "github.com/nuttchai/go-rest/internal/service/interface"
	"github.com/nuttchai/go-rest/internal/util/validators"
)

type TSampleService struct {
	Repository irepository.ISampleRepository
}

var (
	SampleService iservice.ISampleService
)

func initSampleService(sampleService *TSampleService) {
	SampleService = sampleService
}

func (s *TSampleService) Test() string {
	return s.Repository.Test()
}

func (s *TSampleService) GetSample(id string) (*model.Sample, error) {
	return s.Repository.RetrieveOne(id)
}

func (s *TSampleService) CreateSample(sample *dto.CreateSampleDTO) (*model.Sample, error) {
	return s.Repository.CreateOne(sample)
}

func (s *TSampleService) UpdateSample(sample *dto.UpdateSampleDTO) (*model.Sample, error) {
	return s.Repository.UpdateOne(sample)
}

func (s *TSampleService) DeleteSample(id string) error {
	result, err := s.Repository.DeleteOne(id)
	if err != nil {
		return err
	}

	return validators.CheckRowsAffected(result)
}
