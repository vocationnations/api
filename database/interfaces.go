package database

import "github.com/go-pg/pg/v10"

type Database interface {
	GetDatabase() *pg.DB
}
