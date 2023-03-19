package model

import "database/sql"

type DBModel struct {
	DB *sql.DB
}

type Models struct {
	DBModel DBModel
}

func Init(DB *sql.DB) Models {
	return Models{
		DBModel: DBModel{
			DB: DB,
		},
	}
}
