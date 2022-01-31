package routes

import (
	"github.com/vocationnations/api/handler"
	"github.com/vocationnations/api/handler/onet/jobfamily-handler"
	"github.com/vocationnations/api/handler/onet/jobsearch-handler"
	"github.com/vocationnations/api/handler/vns/culture_handler"
	"github.com/vocationnations/api/handler/vns/profession-handler"
	"github.com/vocationnations/api/handler/vns/skill-handler"
	"github.com/vocationnations/api/handler/vns/user-handler"
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
	POST   = "POST"
	UPDATE = "UPDATE"
)

// Routes is where main routes for the API are stored
type Routes []Route

var AllRoutes = Routes{

	////////////
	// USERS //
	//////////
	Route{"GetUsers", GET, "/get_users", user_handler.GetUsers},
	Route{"GetUser", GET, "/get_user/{id}", user_handler.GetUser},
	Route{"GetUserName", GET, "/get_username/{id}", user_handler.GetUserName},
	Route{"GetUserEmail", GET, "/get_user_email/{id}", user_handler.GetUserEmail},

	Route{"SetUser", POST, "/set_user", user_handler.SetUser},
	Route{"SetUserName", UPDATE, "/set_username_by_id", user_handler.SetUserName},
	Route{"SetUserEmail", UPDATE, "/set_user_email_by_id", user_handler.SetUserEmail},

	//////////////////
	// PROFESSIONS //
	////////////////
	Route{"GetUserProfession", GET, "/get_user_professions/{id}", profession_handler.GetUserProfessions},
	Route{"CreateUserProfession", POST, "/create_user_profession", profession_handler.CreateUserProfession},

	/////////////
	// SKILLS //
	///////////
	Route{"GetUserSkills", GET, "/get_user_skills/{id}", skill_handler.GetUserSkills},
	Route{"CreateUserSkill", POST, "/create_user_skill", skill_handler.CreateUserSkill},
	Route{"GetUserSkillsManual", GET, "/get_user_skills_manual/{id}", skill_handler.GetUserSkillsManual},
	Route{"CreateUserSkillManual", POST, "/create_user_skill_manual", skill_handler.CreateUserSkillManual},

	///////////////////
	// CULTURE ENTRY //
	//////////////////
	Route{"CreateCultureEntry", POST, "/create_culture_entry", culture_handler.CreateCultureEntry},

	/////////////////////////////
	// ONET RELATED ENDPOINTS //
	///////////////////////////

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

	////12.GetTechnologySkillsByOccupation
	Route{"GetTechnologySkillsByOccupation", GET, "/get_technology_skills_by_occupation/{id}", jobfamily_handler.GetTechnologySkillsByOccupation},

}
