package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/secure"
	"github.com/vocationnations/api/helpter"
	"github.com/vocationnations/api/routes"
	"log"
	"net/http"
)

// StartAPIServer creates and serves the API on the newly created server
func StartAPIServer(ctx helpter.AppContext) {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes.AllRoutes {
		var handler http.Handler
		handler = helpter.MakeHandler(ctx, route.Handler)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// security
	var isDevelopment = false
	if ctx.Env == helpter.ENV_LOCAL {
		isDevelopment = true
	}
	secureMiddleware := secure.New(secure.Options{
		IsDevelopment:      isDevelopment, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
		AllowedHosts:       []string{},    // AllowedHosts is a list of fully qualified domain names that are allowed (CORS)
		ContentTypeNosniff: true,          // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:   true,          // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
	})

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.HandlerFunc(secureMiddleware.HandlerFuncWithNext))
	n.UseHandler(router)
	log.Println("===> Starting app (v" + ctx.Version + ") on port " + ctx.Port + " in " + ctx.Env + " mode.")
	if ctx.Env == helpter.ENV_LOCAL {
		n.Run("localhost:" + ctx.Port)
	} else {
		n.Run(":" + ctx.Port)
	}
}
