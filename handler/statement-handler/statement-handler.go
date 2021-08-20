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

func GetStatements(w http.ResponseWriter, _ *http.Request, ctx helper.AppContext) error {
	db := ctx.DB.GetDatabase()

	var stmts []model.Statement

	if err := db.Model(&stmts).Select(); err != nil {
		return fmt.Errorf("cannot get users, err: %v", err)
	}

	if len(stmts) == 0 {
		_ = json.NewEncoder(w).Encode(&[]model.Statement{})
	} else {
		if err := json.NewEncoder(w).Encode(stmts); err != nil {
			return fmt.Errorf("cannot encode users, err: %v", err)
		}
	}

	return nil
}

func GetStatement(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {

	params := mux.Vars(r)

	db := ctx.DB.GetDatabase()

	// get the id from the url
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("the id '%s' is not a valid integer", idStr)
	}

	stmt := &model.Statement{Id: id}

	if err := db.Model(stmt).WherePK().Select(); err != nil {
		return fmt.Errorf("statement cannot be retrieved, err: %v", err)
	}
	if err := json.NewEncoder(w).Encode(stmt); err != nil {
		return fmt.Errorf("cannot encode statement, err: %v", err)
	}

	return nil
}

func CreateStatement(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {

	var stmt model.Statement
	db := ctx.DB.GetDatabase()

	if err := json.NewDecoder(r.Body).Decode(&stmt); err != nil {
		return fmt.Errorf("cannot decode the body for statement, err: %v", err)
	}

	if _, err := db.Model(&stmt).Insert(); err != nil {
		return fmt.Errorf("cannot insert statement to database, err: %v", err)
	}

	if err := json.NewEncoder(w).Encode(&stmt); err != nil {
		return fmt.Errorf("cannot encode the statement for printing, err: %v", err)
	}
	return nil
}