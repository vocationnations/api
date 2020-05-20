package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vocationnations/api/helpter"
)

const (
	DBNAME      = "./database/onet.db" // onet database
	CURRENT_ENV = "local"
)

func main() {

	db, err := sql.Open("sqlite3", DBNAME)
	if err != nil {
		fmt.Println("Database cannot connect, err:", err)
	}

	ctx := helpter.AppContext{
		DB:      db,
		Env:     CURRENT_ENV,
		Version: "v0.1",
		Port:    "5000",
	}
	StartAPIServer(ctx)
}
