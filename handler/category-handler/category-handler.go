package category_handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vocationnations/api/helper"
	"github.com/vocationnations/api/model"
	"net/http"
	"strconv"
)

func GetCultureCategories(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {

	db := ctx.DB.GetDatabase()

	var cgres []model.CultureCategory

	if err := db.Model(&cgres).Select(); err != nil {
		return fmt.Errorf("cannot get categories, err: %v", err)
	}

	if len(cgres) == 0 {
		_ = json.NewEncoder(w).Encode(&[]model.User{})
	} else {
		if err := json.NewEncoder(w).Encode(cgres); err != nil {
			return fmt.Errorf("cannot encode users, err: %v", err)
		}
	}
	return nil
}

func GetCultureCategory(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	params := mux.Vars(r)

	var db = ctx.DB.GetDatabase()

	// get the id from the url
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("the id '%s' is not a valid integer", idStr)
	}

	ctgry := &model.CultureCategory{Id: id}

	if err := db.Model(ctgry).WherePK().Select(); err != nil {
		return fmt.Errorf("user cannot be retrieved, err: %v", err)
	}
	if err := json.NewEncoder(w).Encode(ctgry); err != nil {
		return fmt.Errorf("cannot encode user, err: %v", err)
	}
	return nil
}
