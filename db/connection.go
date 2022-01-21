package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Connection *sql.DB

func init() {
}
