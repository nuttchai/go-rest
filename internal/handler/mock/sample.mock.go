package mhandler

import (
	sampledto "github.com/nuttchai/go-rest/internal/dto/sample"
	"github.com/nuttchai/go-rest/internal/model"
)

type TSampleServiceMock struct{}
type TUserServiceMock struct{}

var (
	TestMock         func() string
	GetSampleMock    func(id string) (*model.Sample, error)
	CreateSampleMock func(sample *sampledto.CreateSampleDTO) (*model.Sample, error)
	UpdateSampleMock func(sample *sampledto.UpdateSampleDTO) (*model.Sample, error)
	DeleteSampleMock func(id string) error
	GetUserMock      func(id string) (*model.User, error)
)

func (*TSampleServiceMock) Test() string {
	return TestMock()
}

func (*TSampleServiceMock) GetSample(id string) (*model.Sample, error) {
	return GetSampleMock(id)
}

func (*TSampleServiceMock) CreateSample(sample *sampledto.CreateSampleDTO) (*model.Sample, error) {
	return CreateSampleMock(sample)
}

func (*TSampleServiceMock) UpdateSample(sample *sampledto.UpdateSampleDTO) (*model.Sample, error) {
	return UpdateSampleMock(sample)
}

func (*TSampleServiceMock) DeleteSample(id string) error {
	return DeleteSampleMock(id)
}

func (*TUserServiceMock) GetUser(id string) (*model.User, error) {
	return GetUserMock(id)
}
