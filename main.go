package main

import (
	"github.com/joho/godotenv"
	"github.com/vocationnations/api/database"
	"github.com/vocationnations/api/helper"
	"github.com/vocationnations/api/server"
	"log"
	"os"
)

func main() {

	// load the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalln("ERROR: Please create .env file with appropriate variables")
	}

	// load all the relevant variables
	var (
		DBName  = os.Getenv("DB_NAME")
		DBPort  = os.Getenv("DB_PORT")
		DBHost  = os.Getenv("DB_HOST")
		DBUser  = os.Getenv("DB_USER")
		DBPass  = os.Getenv("DB_PASS")
		APIPort = os.Getenv("API_PORT")
	)

	//initiate a new database connection
	db := database.NewConnection(DBName, DBUser, DBHost, DBPort, DBPass)

	ctx := helper.AppContext{
		DB:      db,
		Version: "v0.1",
		Env:     "local",
		Port:    APIPort,
	}
	server.StartAPIServer(ctx)
}
