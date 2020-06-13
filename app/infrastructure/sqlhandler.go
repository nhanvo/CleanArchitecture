package infrastructure

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}
func NewSqlHandler() *SqlHandler  {

	sqlHandler := new (SqlHandler)
	return sqlHandler
}
