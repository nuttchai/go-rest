package model

type Sample struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerId     int    `json:"owner_id"`
}
