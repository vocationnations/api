package skill_handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vocationnations/api/helper"
	"github.com/vocationnations/api/model"
	"net/http"
	"strconv"
)

func GetUserSkills(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	params := mux.Vars(r)

	var db = ctx.DB.GetDatabase()

	// get the id from the url
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	fmt.Printf("ID %d", id)
	if err != nil {
		return fmt.Errorf("the id '%s' is not a valid integer", idStr)
	}

	var userskills []model.UserSkills

	if err := db.Model(&userskills).Where("user_id = ?", id).Select(); err != nil {
		return fmt.Errorf("user skills cannot be retrieved, err: %v", err)
	}
	if err := json.NewEncoder(w).Encode(userskills); err != nil {
		return fmt.Errorf("cannot encode user skills, err: %v", err)
	}
	return nil
}

func GetUserSkillsManual(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	fmt.Print("Sup")
	params := mux.Vars(r)

	var db = ctx.DB.GetDatabase()

	//  get the id from the url
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	fmt.Printf("ID <%d>", id)
	if err != nil {
		return fmt.Errorf("the id %s is not a valid integer", idStr)
	}

	var userskillsManual []model.SkillManual

	if err := db.Model(&userskillsManual).Where("user_id = ?", id).Select(); err != nil {
		return fmt.Errorf("user skills cannot be retrieved, err: %v", err)
	}

	if err := json.NewEncoder(w).Encode(userskillsManual); err != nil {
		return fmt.Errorf("cannot encode the userskills_manual, err: %v", err)
	}

	return nil
}

func CreateUserSkill(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {

	var usrSkill model.UserSkills
	db := ctx.DB.GetDatabase()

	if err := json.NewDecoder(r.Body).Decode(&usrSkill); err != nil {
		return fmt.Errorf("cannot decode the body for user_skills, err: %v", err)
	}

	if _, err := db.Model(&usrSkill).Insert(); err != nil {
		return fmt.Errorf("cannot insert user skill to the database, err: %v", err)
	}

	if err := json.NewEncoder(w).Encode(&usrSkill); err != nil {
		return fmt.Errorf("cannot encode the skill for printing, err: %v", err)
	}
	return nil
}

func CreateUserSkillManual(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	var usrSkillManual []model.SkillManual

	db := ctx.DB.GetDatabase()

	if err := json.NewDecoder(r.Body).Decode(&usrSkillManual); err != nil {
		return fmt.Errorf("cannot decode the body for users_skills_manual, err : %v", err)
	}

	if _, err := db.Model(&usrSkillManual).Insert(); err != nil {
		return fmt.Errorf("cannot insert user_skills_manual to the database, err: %v", err)
	}

	if err := json.NewEncoder(w).Encode(&usrSkillManual); err != nil {
		return fmt.Errorf("cannot encode the skill for printing %v", err)
	}
	return nil
}
