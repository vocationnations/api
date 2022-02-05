package climate_handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vocationnations/api/helper"
	"github.com/vocationnations/api/model"
	"net/http"
	"strconv"
)

func CreateClimateQuestion(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {

	var climateQuestion model.ClimateQuestions
	db := ctx.DB.GetDatabase()

	if err := json.NewDecoder(r.Body).Decode(&climateQuestion); err != nil {
		return fmt.Errorf("cannot decode the body for climate_questions, err: %v", err)
	}

	if _, err := db.Model(&climateQuestion).Insert(); err != nil {
		return fmt.Errorf("cannot insert the climate question to the database, err: %v", err)
	}

	if err := json.NewEncoder(w).Encode(&climateQuestion); err != nil {
		return fmt.Errorf("cannot encode the climate question for printing, err: %v", err)
	}
	return nil
}

func GetClimateQuestions(w http.ResponseWriter, _ *http.Request, ctx helper.AppContext) error {

	var db = ctx.DB.GetDatabase()

	var climateQuestions []model.ClimateQuestions

	err := db.Model(&climateQuestions).Select()
	if err != nil {
		return fmt.Errorf("cannot get climate questions, err: %v", err)
	}

	if len(climateQuestions) == 0 {
		_ = json.NewEncoder(w).Encode(&[]model.ClimateQuestions{})
	} else {
		err = json.NewEncoder(w).Encode(climateQuestions)
		if err != nil {
			return fmt.Errorf("cannot encode climate questions, err: %v", err)
		}
	}
	return nil
}


func CreateStep(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {

	var step model.Steps
	db := ctx.DB.GetDatabase()

	if err := json.NewDecoder(r.Body).Decode(&step); err != nil {
		return fmt.Errorf("cannot decode the body for steps, err: %v", err)
	}

	if _, err := db.Model(&step).Insert(); err != nil {
		return fmt.Errorf("cannot insert the step to the database, err: %v", err)
	}

	if err := json.NewEncoder(w).Encode(&step); err != nil {
		return fmt.Errorf("cannot encode the step for printing, err: %v", err)
	}
	return nil
}

func GetStepsByQuestionId(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	params := mux.Vars(r)

	var db = ctx.DB.GetDatabase()

	// get the id from the url
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	fmt.Printf("ID %d", id)
	if err != nil {
		return fmt.Errorf("the id '%s' is not a valid integer", idStr)
	}

	var steps []model.Steps

	if err := db.Model(&steps).Where("question_id = ?", id).Select(); err != nil {
		return fmt.Errorf("steps cannot be retrieved, err: %v", err)
	}
	if err := json.NewEncoder(w).Encode(steps); err != nil {
		return fmt.Errorf("cannot encode steps, err: %v", err)
	}
	return nil
}


func CreateClimateQuestionAnswer(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {

	var cqa model.ClimateQuestionAnswers
	db := ctx.DB.GetDatabase()

	if err := json.NewDecoder(r.Body).Decode(&cqa); err != nil {
		return fmt.Errorf("cannot decode the body for Climate Question Answers, err: %v", err)
	}

	if _, err := db.Model(&cqa).Insert(); err != nil {
		return fmt.Errorf("cannot insert the ClimateQuestionAnswers to the database, err: %v", err)
	}

	if err := json.NewEncoder(w).Encode(&cqa); err != nil {
		return fmt.Errorf("cannot encode the Climate Question Asnwers for printing, err: %v", err)
	}
	return nil
}


