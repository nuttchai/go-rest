package sampledto

type UpdateSampleDTO struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}
