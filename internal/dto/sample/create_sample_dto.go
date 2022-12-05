package sample

type CreateSampleDTO struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type CreatedSampleDTO struct {
	Id int `json:"id"`
}
