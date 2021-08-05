package routes

import (
	"github.com/vocationnations/api/handler"
	"github.com/vocationnations/api/handler/category"
	"github.com/vocationnations/api/handler/skill"
	"github.com/vocationnations/api/handler/statement"
	"github.com/vocationnations/api/handler/user"
	"github.com/vocationnations/api/handler/userentry"
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
	Route{ "GetCultureStatements", GET, "/get_culture_statements", statement.GetCultureStatements},
	Route{ "GetCultureStatement", GET, "/get_culture_statement/{id}", statement.GetCultureStatement},
	//
	//// 3. CultureCategory
	Route{ "GetCultureCategories", GET, "/get_culture_categories", category.GetCultureCategories},
	Route{ "GetCultureCategory", GET, "/get_culture_category/{id}", category.GetCultureCategory},
	//
	//// 4. Skill
	Route{"GetSkills", GET, "/get_skills", skill.GetSkills},
	Route{"GetSkill", GET, "/get_skill/{id}", skill.GetSkill},
	//
	//// 5. UserEntry
	Route{"GetUserEntries", GET, "/get_user_entries", userentry.GetUserEntries},
	Route{"GetUserEntry", GET, "/get_user_entry/{id}", userentry.GetUserEntry},
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
