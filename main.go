package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vocationnations/api/helpter"
)

const (
	DBNAME      = "postgres" // vns-aws-rds-postgres
	DB_USER     = "postgres"
	CURRENT_ENV = "local"
	HOST        = "vns-poc.cvnvlcppwmn7.us-west-2.rds.amazonaws.com"
	DB_PORT     = 5432
	API_PORT    = "5000"
	PASSWORD    = "Vnspoc090621251"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, DB_PORT, DB_USER, PASSWORD, DBNAME)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Database cannot connect, err:", err)
	}

	ctx := helpter.AppContext{
		DB:      db,
		Env:     CURRENT_ENV,
		Version: "v0.1",
		Port:    API_PORT,
	}
	StartAPIServer(ctx)
}
