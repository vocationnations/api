package cultureEntry_handler

import(
    "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vocationnations/api/helper"
	"github.com/vocationnations/api/model"
	"net/http"
	"strconv"
)

func GetcultureEntries(w http.ResponseWriter, _ *http.Request, ctx helper.AppContext) error {

	var db = ctx.DB.GetDatabase()

	var usrs []model.User

	err := db.Model(&usrs).Select()
	if err != nil {
		return fmt.Errorf("cannot get users, err: %v", err)
	}

func GetcultureEntry(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
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