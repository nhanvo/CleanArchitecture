package database

import "database/sql/driver"

type SqlHandler interface {
	Excute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
}

