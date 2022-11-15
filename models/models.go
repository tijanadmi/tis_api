package models

import (
	"database/sql"
)

// Models is the wrapper for database
type Models struct {
	DB DBModel
}

// NewModels returns models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// Zastita is the type for all signals of protective devices
type Signal struct {
	ID   int    `json:"infor_id"`
	Code string `json:"infor_code"`
	Name string `json:"infor_name"`
}

// User is the type for users
type User struct {
	ID       int
	Username    string
	Password string
}