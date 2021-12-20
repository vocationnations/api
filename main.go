package main

import (
	"github.com/joho/godotenv"
	"github.com/vocationnations/api/database"
	"github.com/vocationnations/api/helper"
	"github.com/vocationnations/api/server"
	"log"
	"os"
)

const OnetApiBase = "https://services.onetcenter.org/ws/online/"

func main() {

	// load the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalln("ERROR: Please create .env file with appropriate variables")
	}

	// load all the relevant variables
	var (
		DBPort = os.Getenv("DB_PORT")
		DBName = os.Getenv("DB_NAME")
		DBUser = os.Getenv("DB_USER")
		DBPass = os.Getenv("DB_PASS")
		DBHost = os.Getenv("DB_HOST")

		APIPort    = os.Getenv("API_PORT")
		ApiVersion = os.Getenv("API_VERSION")

		ONETAuthToken   = os.Getenv("ONET_AUTH_TOKEN")
		ONETAuthType    = os.Getenv("ONET_AUTH_TYPE")
		ONETContentType = os.Getenv("ONET_CONTENT_TYPE")
		ONETAcceptType  = os.Getenv("ONET_ACCEPT_TYPE")
	)

	//initiate a new database connection
	db := database.NewConnection(DBName, DBUser, DBHost, DBPort, DBPass)

	ctx := helper.AppContext{
		VNConfiguration: helper.VNConfiguration{
			DatabaseConfiguration: helper.VNDatabaseConfiguration{
				Port: DBPort,
				Name: DBName,
				User: DBUser,
				Pass: DBPass,
				Host: DBHost,
			},
			APIPort: APIPort,
			ONET: helper.ONETConfiguration{
				Base:        OnetApiBase,
				AuthToken:   ONETAuthToken,
				AuthType:    ONETAuthType,
				ContentType: ONETContentType,
				AcceptType:  ONETAcceptType,
			},
			Version: ApiVersion,
			Env:     "local",
		},
		DB: db,
	}
	server.StartAPIServer(ctx)
}
