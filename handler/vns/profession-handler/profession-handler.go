package profession_handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vocationnations/api/helper"
	"github.com/vocationnations/api/model"
	"net/http"
	"strconv"
)

func GetUserProfessions(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	params := mux.Vars(r)

	var db = ctx.DB.GetDatabase()

	// get the id from the url
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	fmt.Printf("ID %d", id)
	if err != nil {
		return fmt.Errorf("the id '%s' is not a valid integer", idStr)
	}

	var userPrfsns []model.UsersProfessions

	if err := db.Model(&userPrfsns).Where("user_id = ?", id).Select(); err != nil {
		return fmt.Errorf("user professions cannot be retrieved, err: %v", err)
	}
	if err := json.NewEncoder(w).Encode(userPrfsns); err != nil {
		return fmt.Errorf("cannot encode user professions, err: %v", err)
	}
	return nil
}

func CreateUserProfession(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {

	var usr_prfsn model.UsersProfessions
	db := ctx.DB.GetDatabase()

	if err := json.NewDecoder(r.Body).Decode(&usr_prfsn); err != nil {
		return fmt.Errorf("cannot decode the body for users_professions, err: %v", err)
	}

	fmt.Printf("UserPrfsn: %s", usr_prfsn)

	if _, err := db.Model(&usr_prfsn).Insert(); err != nil {
		return fmt.Errorf("cannot insert user professions to the database, err: %v", err)
	}

	if err := json.NewEncoder(w).Encode(&usr_prfsn); err != nil {
		return fmt.Errorf("cannot encode the user professions for printing, err: %v", err)
	}
	return nil
}
