package sampledto

type GetSampleWithUserDTO struct {
	Id          int    `json:"id" validate:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
}
