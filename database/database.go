package database

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"log"
	//"log"
)

type structDatabase struct {
	DBName   string
	DBUser   string
	DBHost   string
	DBPort   string
	DBPass   string
	Database *pg.DB
}

func NewConnection(name string, user string, host string, port string, pass string) Database {

	// make the address
	var addr = fmt.Sprintf("%s:%s", host, port)

	//set up a database connection
	db := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Password: pass,
		Database: name,
	})

	// check if the connection was successful
	if err := db.Ping(context.Background()); err != nil {
		log.Fatalf("ERROR: Failed to conect to the database, err: %v", err)
	}

	return &structDatabase{
		DBName:   name,
		DBUser:   user,
		DBHost:   host,
		DBPort:   port,
		DBPass:   pass,
		Database: db,
	}
}

func (d *structDatabase) GetDatabase() *pg.DB {
	return d.Database
}
