package helper

import (
	"database/sql"
	"net/http"
)

const (
	ENV_LOCAL = "local"
)

type AppContext struct {
	// DB is the database that is used by the application
	DB *sql.DB
	// Env will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development env.
	Env string
	// Version defines the API version
	Version string
	// Port defines the port at which the API serves
	Port string
}

// makeHandler allows us to pass an environment struct to our handlers, without resorting to global
// variables. It accepts an environment (Env) struct and our own handler function. It returns
// a function of the type http.HandlerFunc so can be passed on to the HandlerFunc.
func MakeHandler(ctx AppContext, fn func(http.ResponseWriter, *http.Request, AppContext) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, ctx)
	}
}
