package userentry_handler

import (
	"encoding/json"
	"fmt"
	"github.com/vocationnations/api/helper"
	"github.com/vocationnations/api/model"
	"net/http"
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
	// TODO: Implement me
	return nil
}
