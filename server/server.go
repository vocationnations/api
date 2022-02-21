package server

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/unrolled/secure"
	"github.com/vocationnations/api/helper"
	"github.com/vocationnations/api/routes"
	"log"
	"net/http"
)

// StartAPIServer creates and serves the API on the newly created server
func StartAPIServer(ctx helper.AppContext) {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes.AllRoutes {
		var handler http.Handler
		handler = helper.MakeHandler(ctx, route.Handler)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	//security
	var isDevelopment = false
	if ctx.VNConfiguration.Env == helper.ENVLocal {
		isDevelopment = true
	}
	secureMiddleware := secure.New(secure.Options{
		IsDevelopment:      isDevelopment, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
		AllowedHosts:       []string{},    // AllowedHosts is a list of fully qualified domain names that are allowed (CORS)
		ContentTypeNosniff: true,          // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:   true,          // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
	})

	// enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		Debug:          true,
	})

	n := negroni.New()
	n.Use(c)
	n.Use(negroni.NewLogger())
	n.Use(negroni.HandlerFunc(secureMiddleware.HandlerFuncWithNext))
	n.UseHandler(router)
	log.Println("===> Starting app (v" + ctx.VNConfiguration.Version + ") on port " + ctx.VNConfiguration.APIPort + " in " + ctx.VNConfiguration.Env + " mode.")
	if ctx.VNConfiguration.Env == helper.ENVLocal {
		n.Run(":" + ctx.VNConfiguration.APIPort)
	} else {
		n.Run(":" + ctx.VNConfiguration.APIPort)
	}
}
