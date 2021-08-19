package userentry_handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vocationnations/api/helper"
	"github.com/vocationnations/api/model"
	"net/http"
	"strconv"
)

func GetUserEntries(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	var db = ctx.DB.GetDatabase()

	var uentries []model.UserEntry

	err := db.Model(&uentries).Select()
	if err != nil {
		return fmt.Errorf("cannot get users, err: %v", err)
	}

	if len(uentries) == 0 {
		_ = json.NewEncoder(w).Encode(&[]model.UserEntry{})
	} else {
		err = json.NewEncoder(w).Encode(uentries)
		if err != nil {
			return fmt.Errorf("cannot encode users, err: %v", err)
		}
	}
	return nil
}

func GetUserEntry(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	params := mux.Vars(r)

	var db = ctx.DB.GetDatabase()

	// get the id from the url
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("the id '%s' is not a valid integer", idStr)
	}

	uentry := &model.UserEntry{Id: id}

	if err := db.Model(uentry).WherePK().Select(); err != nil {
		return fmt.Errorf("user entry cannot be retrieved, err: %v", err)
	}
	if err := json.NewEncoder(w).Encode(uentry); err != nil {
		return fmt.Errorf("cannot encode user entry, err: %v", err)
	}
	return nil
}
