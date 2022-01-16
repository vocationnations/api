package user_handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vocationnations/api/helper"
	"github.com/vocationnations/api/model"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, _ *http.Request, ctx helper.AppContext) error {

	var db = ctx.DB.GetDatabase()

	var usrs []model.User

	err := db.Model(&usrs).Select()
	if err != nil {
		return fmt.Errorf("cannot get users, err: %v", err)
	}

	if len(usrs) == 0 {
		_ = json.NewEncoder(w).Encode(&[]model.User{})
	} else {
		err = json.NewEncoder(w).Encode(usrs)
		if err != nil {
			return fmt.Errorf("cannot encode users, err: %v", err)
		}
	}
	return nil
}

func GetUser(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	params := mux.Vars(r)

	var db = ctx.DB.GetDatabase()

	// get the id from the url
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("the id '%s' is not a valid integer", idStr)
	}

	usr := &model.User{Id: id}

	if err := db.Model(usr).WherePK().Select(); err != nil {
		return fmt.Errorf("user cannot be retrieved, err: %v", err)
	}
	if err := json.NewEncoder(w).Encode(usr); err != nil {
		return fmt.Errorf("cannot encode user, err: %v", err)
	}
	return nil
}

func SetUser(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {

	var usr model.User
	db := ctx.DB.GetDatabase()

	if err := json.NewDecoder(r.Body).Decode(&usr); err != nil {
		return fmt.Errorf("cannot decode the body for user, err: %v", err)
	}

	if _, err := db.Model(&usr).Insert(); err != nil {
		return fmt.Errorf("cannot insert user to database, err: %v", err)
	}

	if err := json.NewEncoder(w).Encode(&usr); err != nil {
		return fmt.Errorf("cannot encode the user for printing, err: %v", err)
	}
	return nil
}

func GetUserName(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	params := mux.Vars(r)

	var db = ctx.DB.GetDatabase()

	// get the id from the url
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("the id '%s' is not a valid integer", idStr)
	}

	usr := &model.User{Id: id}

	if err := db.Model(usr).WherePK().Select(); err != nil {
		return fmt.Errorf("user cannot be retrieved, err: %v", err)
	}
	if err := json.NewEncoder(w).Encode(usr.Name); err != nil {
		return fmt.Errorf("cannot encode user, err: %v", err)
	}
	return nil
}

func SetUserName(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	//params := mux.Vars(r)
	var usr model.User
	var db = ctx.DB.GetDatabase()

	// get the id from the url
	//idStr := params["id"]
	//nameStr := params["name"]
	//
	//id, err := strconv.Atoi(idStr)
	//if err != nil {
	//	return fmt.Errorf("the id '%s' is not a valid integer", idStr)
	//}
	//
	//usr := &model.User{Id: id, Name: nameStr}

	if err := json.NewDecoder(r.Body).Decode(&usr); err != nil {
		return fmt.Errorf("cannot decode the body for user, err: %v", err)
	}
	//if err := db.Model(usr).WherePK().Select(); err != nil {
	//	return fmt.Errorf("user cannot be retrieved, err: %v", err)
	//}

	res, err := db.Model(usr).WherePK().Update()
	if err != nil {
		return fmt.Errorf("user cannot be retrieved, err: %v", err)
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		return fmt.Errorf("cannot encode user, err: %v", err)
	}
	return nil
}

func GetUserEmail(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	params := mux.Vars(r)

	var db = ctx.DB.GetDatabase()

	// get the id from the url
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("the id '%s' is not a valid integer", idStr)
	}

	usr := &model.User{Id: id}

	if err := db.Model(usr).WherePK().Select(); err != nil {
		return fmt.Errorf("user cannot be retrieved, err: %v", err)
	}
	if err := json.NewEncoder(w).Encode(usr.Email); err != nil {
		return fmt.Errorf("cannot encode user, err: %v", err)
	}
	return nil
}


func SetUserEmail(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	var usr model.User
	var db = ctx.DB.GetDatabase()

	if err := json.NewDecoder(r.Body).Decode(&usr); err != nil {
		return fmt.Errorf("cannot decode the body for user, err: %v", err)
	}

	res, err := db.Model(usr).WherePK().Update()
	if err != nil {
		return fmt.Errorf("user cannot be retrieved, err: %v", err)
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		return fmt.Errorf("cannot encode user, err: %v", err)
	}
	return nil
}
