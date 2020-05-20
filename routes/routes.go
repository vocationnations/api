package routes

import (
	"github.com/vocationnations/api/handler"
)

// Route is the model for the route setup
type Route struct {
	// Name of the route
	Name string
	// Method for the HTTP route (e.g., GET, PUT, POST, PATCH, DELETE)
	Method string
	// Pattern for the route/endpoint URL
	Pattern string
	// Handler function for the route
	Handler handler.HandlerFunc
}

// constants for the HTTP request method type
const (
	GET    = "GET"
	PUT    = "PUT"
	POST   = "POST"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

// Routes is where main routes for the API are stored
type Routes []Route

var AllRoutes = Routes{
	Route{"Main", GET, "/", handler.GetMain},
	Route{"OccupationsMain", GET, "/occupations", handler.GetOccupations},
	Route{"OccupationByCode", GET, "/occupation/code/{onetcode}", handler.GetOccupationByCode},
	Route{"OccupationByTitle", GET, "/occupation/title/{title}", handler.GetOccupationByTitle},
	Route{"UserById", GET, "/user/{id}", handler.GetUserById},
}
