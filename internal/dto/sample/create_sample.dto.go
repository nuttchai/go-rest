package sampledto

type CreateSampleDTO struct {
	Name string `json:"name" validate:"required"`
	Desc string `json:"desc"`
}
