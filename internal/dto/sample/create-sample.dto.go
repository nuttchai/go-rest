package sampledto

type CreateSampleDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	OwnerId     int    `json:"owner_id"`
}
