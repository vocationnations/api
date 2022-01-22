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

func CreateUserSkill(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {

	var usr_skill model.UserSkills
	db := ctx.DB.GetDatabase()

	if err := json.NewDecoder(r.Body).Decode(&usr_skill); err != nil {
		return fmt.Errorf("cannot decode the body for users_professions, err: %v", err)
	}

	if _, err := db.Model(&usr_skill).Insert(); err != nil {
		return fmt.Errorf("cannot insert user professions to the database, err: %v", err)
	}

	if err := json.NewEncoder(w).Encode(&usr_skill); err != nil {
		return fmt.Errorf("cannot encode the statement for printing, err: %v", err)
	}
	return nil
}

func GetSkills(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	var db = ctx.DB.GetDatabase()

	var skills []model.Skill

	err := db.Model(&skills).Select()
	if err != nil {
		return fmt.Errorf("cannot get skills, err: %v", err)
	}

	if len(skills) == 0 {
		_ = json.NewEncoder(w).Encode(&[]model.Skill{})
	} else {
		err = json.NewEncoder(w).Encode(skills)
		if err != nil {
			return fmt.Errorf("cannot encode skills, err: %v", err)
		}
	}
	return nil
}

func GetSkill(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	params := mux.Vars(r)

	var db = ctx.DB.GetDatabase()

	// get the id from the url
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("the id '%s' is not a valid integer", idStr)
	}

	skill := &model.Skill{Id: id}

	if err := db.Model(skill).WherePK().Select(); err != nil {
		return fmt.Errorf("skill cannot be retrieved, err: %v", err)
	}
	fmt.Println(skill)
	if err := json.NewEncoder(w).Encode(skill); err != nil {
		return fmt.Errorf("cannot encode skill, err: %v", err)
	}
	return nil
}


func CreateSkill(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {

	var skill model.Skill
	db := ctx.DB.GetDatabase()

	if err := json.NewDecoder(r.Body).Decode(&skill); err != nil {
		return fmt.Errorf("cannot decode the body for skill, err: %v", err)
	}

	if _, err := db.Model(&skill).Insert(); err != nil {
		return fmt.Errorf("cannot insert skill to database, err: %v", err)
	}

	if err := json.NewEncoder(w).Encode(&skill); err != nil {
		return fmt.Errorf("cannot encode the skill for printing, err: %v", err)
	}
	return nil
}