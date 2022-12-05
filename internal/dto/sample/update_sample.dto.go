package sample

type UpdateSampleDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type DeletedSampleDTO struct {
	Id      int  `json:"id"`
	Deleted bool `json:"status"`
}
