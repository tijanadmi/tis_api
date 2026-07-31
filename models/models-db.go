package models

import (
	"database/sql"
)

type OracleDBRepo struct {
	DB *sql.DB
}
