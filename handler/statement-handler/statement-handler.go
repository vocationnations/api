package statement_handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vocationnations/api/helper"
	"github.com/vocationnations/api/model"
	"net/http"
	"strconv"
)

func GetCategoryStatements(w http.ResponseWriter, _ *http.Request, ctx helper.AppContext) error {
	db := ctx.DB.GetDatabase()

	var stmts []model.CategoryStatement

	if err := db.Model(&stmts).Select(); err != nil {
		return fmt.Errorf("cannot get users, err: %v", err)
	}

	if len(stmts) == 0 {
		_ = json.NewEncoder(w).Encode(&[]model.CategoryStatement{})
	} else {
		if err := json.NewEncoder(w).Encode(stmts); err != nil {
			return fmt.Errorf("cannot encode users, err: %v", err)
		}
	}

	return nil
}

func GetCategoryStatement(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {

	params := mux.Vars(r)

	db := ctx.DB.GetDatabase()

	// get the id from the url
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("the id '%s' is not a valid integer", idStr)
	}

	stmt := &model.CategoryStatement{Id: id}

	if err := db.Model(stmt).WherePK().Select(); err != nil {
		return fmt.Errorf("statement cannot be retrieved, err: %v", err)
	}
	if err := json.NewEncoder(w).Encode(stmt); err != nil {
		return fmt.Errorf("cannot encode statement, err: %v", err)
	}

	return nil
}
