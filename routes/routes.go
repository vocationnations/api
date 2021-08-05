package routes

import (
	"github.com/vocationnations/api/handler"
	"github.com/vocationnations/api/handler/user"
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
	// GET
	// 1. Users
	Route{"GetUsers", GET, "/get_users", user.GetUsers},
	Route{"GetUser", GET, "/get_user/{id}", user.GetUser},
	//
	//// 2. CategoryStatement
	//Route{ "GetCultureStatements", GET, "/get_culture_statements", handler.GetCultureStatements},
	//Route{ "GetCultureStatement", GET, "/get_culture_statement/{id}", handler.GetCultureStatement},
	//
	//// 3. CultureCategory
	//Route{ "GetCultureCategories", GET, "/get_culture_categories", handler.GetCultureCategories},
	//Route{ "GetCultureCategory", GET, "/get_culture_category/{id}", handler.GetCultureCategory},
	//
	//// 4. Skill
	//Route{"GetSkills", GET, "/get_skills", handler.GetSkills},
	//Route{"GetSkill", GET, "/get_skill/{id}", handler.GetSkill},
	//
	//// 5. UserEntry
	//Route{"GetUserEntries", GET, "/get_user_entries", handler.GetUserEntries},
	//Route{"GetUserEntry", GET, "/get_user_entry/{id}", handler.GetUserEntry},
	//
	//// POST
	//// 1. Users
	//Route{"AddUser", POST, "/add_user", handler.AddUser},
	//
	//// 2. CategoryStatement
	//Route{"AddCategoryStatement", POST, "/add_category_statement", handler.AddCategoryStatement},
	//
	//// 3. CultureCategory
	//Route{ "AddCultureCategory", POST, "/add_culture_category", handler.AddCultureCategory},
	//
	//// 4. Skill
	//Route{"AddSkills", POST, "/add_skills", handler.AddSkills},

}
