package routes

import (
	"github.com/vocationnations/api/handler"
	"github.com/vocationnations/api/handler/onet/jobfamily-handler"
	"github.com/vocationnations/api/handler/onet/jobsearch-handler"
	"github.com/vocationnations/api/handler/vns/category-handler"
	"github.com/vocationnations/api/handler/vns/skill-handler"
	"github.com/vocationnations/api/handler/vns/statement-handler"
	"github.com/vocationnations/api/handler/vns/user-handler"
	"github.com/vocationnations/api/handler/vns/userentry-handler"
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
	GET  = "GET"
	POST = "POST"
)

// Routes is where main routes for the API are stored
type Routes []Route

var AllRoutes = Routes{

	Route{"GetUsers", GET, "/get_users", user_handler.GetUsers},
	Route{"GetUser", GET, "/get_user/{id}", user_handler.GetUser},

	Route{"GetCategoryStatements", GET, "/get_category_statements", statement_handler.GetStatements},
	Route{"GetCategoryStatement", GET, "/get_category_statement/{id}", statement_handler.GetStatement},

	Route{"GetCultureCategories", GET, "/get_culture_categories", category_handler.GetCultureCategories},
	Route{"GetCultureCategory", GET, "/get_culture_category/{id}", category_handler.GetCultureCategory},

	Route{"GetSkills", GET, "/get_skills", skill_handler.GetSkills},
	Route{"GetSkill", GET, "/get_skill/{id}", skill_handler.GetSkill},

	Route{"GetUserEntries", GET, "/get_user_entries", userentry_handler.GetUserEntries},
	Route{"GetUserEntry", GET, "/get_user_entry/{id}", userentry_handler.GetUserEntry},

	Route{"CreateUser", POST, "/create_user", user_handler.CreateUser},

	//// 2. Statement
	Route{"CreateCategoryStatement", POST, "/create_category_statement", statement_handler.CreateStatement},
	//
	//// 3. CultureCategory
	Route{"CreateCultureCategory", POST, "/create_culture_category", category_handler.CreateCultureCategory},
	//
	//// 4. Skill
	Route{"CreateSkill", POST, "/create_skill", skill_handler.CreateSkill},
	//
	//// 5. User Entry
	Route{"CreateUserEntry", POST, "/create_user_entry", userentry_handler.CreateUserEntry},

	//// 5. Job Families
	Route{"GetJobFamilies", GET, "/get_job_families", jobfamily_handler.GetJobFamilies},

	//// 6. GetOccupationsFromFamily
	Route{"GetOccupationsFromFamily", GET, "/get_occupations_from_family/{id}", jobfamily_handler.GetOccupationsFromFamily},

	//// 7.GetTasksByOccupation
	Route{"GetTasksByOccupation", GET, "/get_tasks_by_occupation/{id}", jobfamily_handler.GetTasksByOccupation},

	//// 8.GetSkillsByOccupation
	Route{"GetSkillsByOccupation", GET, "/get_skills_by_occupation/{id}", jobfamily_handler.GetSkillsByOccupation},

	//// 9. GetKnowledgeSummaryByOccupation
	Route{"GetKnowledgeSummaryByOccupation", GET, "/get_knowledge_summary_by_occupation/{id}", jobfamily_handler.GetKnowledgeSummaryByOccupation},

	////10.GetEducationOccupation
	Route{"GetEducationByOccupation", GET, "/get_education_by_occupation/{id}", jobfamily_handler.GetEducationOccupation},

	////11.GetJobResults
	Route{"GetJobResults", GET, "/get_job_by_keyword/{keyword}", jobsearch_handler.GetJobResults},

}
