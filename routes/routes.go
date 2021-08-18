package routes

import (
	"github.com/vocationnations/api/handler"
	"github.com/vocationnations/api/handler/category-handler"
	"github.com/vocationnations/api/handler/skill-handler"
	"github.com/vocationnations/api/handler/statement-handler"
	"github.com/vocationnations/api/handler/user-handler"
	"github.com/vocationnations/api/handler/userentry-handler"
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
	Handler handler.Func
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

	Route{"GetUsers", GET, "/get_users", user_handler.GetUsers},
	Route{"GetUser", GET, "/get_user/{id}", user_handler.GetUser},

	Route{"GetCategoryStatements", GET, "/get_category_statements", statement_handler.GetCategoryStatements},
	Route{"GetCategoryStatement", GET, "/get_category_statement/{id}", statement_handler.GetCategoryStatement},

	Route{"GetCultureCategories", GET, "/get_culture_categories", category_handler.GetCultureCategories},
	Route{"GetCultureCategory", GET, "/get_culture_category/{id}", category_handler.GetCultureCategory},

	Route{"GetSkills", GET, "/get_skills", skill_handler.GetSkills},
	Route{"GetSkill", GET, "/get_skill/{id}", skill_handler.GetSkill},

	Route{"GetUserEntries", GET, "/get_user_entries", userentry_handler.GetUserEntries},
	Route{"GetUserEntry", GET, "/get_user_entry/{id}", userentry_handler.GetUserEntry},

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
