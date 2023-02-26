package sampledto

type UpdateSampleDTO struct {
	Id          int    `json:"id" validate:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
