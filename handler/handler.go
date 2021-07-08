package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vocationnations/api/helpter"
	"github.com/vocationnations/api/model"
	"log"
	"net/http"
)

// GetMain serves the main route with pattern "/"
func GetMain(w http.ResponseWriter, _ *http.Request, ctx helpter.AppContext) error {
	query := `SELECT * FROM Users`
	rows, err := ctx.DB.Query(query)
	if err != nil {
		fmt.Errorf("cannot run query: %s : %v", query, err)
	}
	for rows.Next() {
		var name string
		var usertype string
		var id string
		if err := rows.Scan(&id, &name, &usertype); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s and %s %s\n", id, name, usertype)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	if _, err := fmt.Fprintf(w, "Welcome to VocationNation API"); err != nil {
		return fmt.Errorf("cannot load the main route, err: %v", err)
	}
	return nil
}

// GetOccupations gets all the occupations in the database. Serves "/occupations"
func GetOccupations(w http.ResponseWriter, r *http.Request, ctx helpter.AppContext) error {
	var AllOccupations []model.Occupation
	rows, err := ctx.DB.Query(`SELECT * FROM occupation_data`)
	if err != nil {
		return fmt.Errorf("cannot run query: `SELECT * FROM occupation_data`, err: %v", err)
	}

	for rows.Next() {
		var _occupation model.Occupation
		if err := rows.Scan(&_occupation.Onetcode, &_occupation.Title, &_occupation.Description); err != nil {
			return fmt.Errorf("cannot scan appropriate rows for GetOccupations, err: %v", err)
		}
		AllOccupations = append(AllOccupations, _occupation)
	}

	err = json.NewEncoder(w).Encode(AllOccupations)
	if err != nil {
		return fmt.Errorf("JSON encoding failed, err: %v", err)
	}
	return nil
}

// GetOccupationByCode gets all the occupation by provided onetcode
func GetOccupationByCode(w http.ResponseWriter, r *http.Request, ctx helpter.AppContext) error {
	params := mux.Vars(r)
	onetcode := params["code"]

	var occup model.Occupation

	rows, err := ctx.DB.Query(`SELECT * FROM occupation_data WHERE onetcode = ?`, onetcode)
	if err != nil {
		return fmt.Errorf("Cannot select occupations, err: %v", err)
	}

	for rows.Next() {
		var _occupation model.Occupation

		err = rows.Scan(&_occupation.Onetcode, &_occupation.Title, &_occupation.Description)
		if err != nil {
			return fmt.Errorf("Cannot scan appropriate rows, err: %v", err)
		}
		occup = _occupation
		break
	}

	err = json.NewEncoder(w).Encode(occup)
	if err != nil {
		return fmt.Errorf("JSON encdoing failed for GetOccupationByCode, err: %v", err)
	}
	return nil
}

// GetOccupationByTitle gets all the occupations pertaining to provided title
func GetOccupationByTitle(w http.ResponseWriter, r *http.Request, ctx helpter.AppContext) error {
	params := mux.Vars(r)
	title := params["title"]

	var occupations []model.Occupation

	search_title := "%" + title + "%"
	rows, err := ctx.DB.Query(`SELECT * FROM occupation_data WHERE title LIKE ?`, search_title)
	if err != nil {
		return fmt.Errorf("Cannot select occupations with title, err: %v", err)
	}

	for rows.Next() {
		var _occupation model.Occupation

		err = rows.Scan(&_occupation.Onetcode, &_occupation.Title, &_occupation.Description)
		if err != nil {
			return fmt.Errorf("Cannot scan appropriate rows, err: %v", err)
		}
		occupations = append(occupations, _occupation)
	}

	err = json.NewEncoder(w).Encode(occupations)
	if err != nil {
		return fmt.Errorf("JSON encoding for GetOccupationByTtile failed, err: %v", err)
	}
	return nil
}

// GetUserById gets the user data pertaining to provided id
func GetUserById(w http.ResponseWriter, r *http.Request, ctx helpter.AppContext) error {
	params := mux.Vars(r)
	id := params["id"]

	var usr model.User

	rows, err := ctx.DB.Query(`
		SELECT 
			u.id AS userid, 
			u.name AS name, 
			e.id AS educationid, 
			e.degreetitle, 
			e.start AS education_start, 
			e.end AS education_end, 
			e.institution, 
			w.id AS workid, 
			w.position, 
			w.organization, 
			w.start AS work_start, 
			w.end AS work_end
		FROM users as u
		INNER JOIN education AS e ON e.userid = u.id
		INNER JOIN work AS w ON w.userid = u.id
		WHERE u.id = ?;
	`, id)
	if err != nil {
		return fmt.Errorf("Cannot select user data, err: %v", err)
	}

	for rows.Next() {
		var _user model.User

		err = rows.Scan(&_user.Id,
			&_user.Name,
			&_user.EducationHistory.Id,
			&_user.EducationHistory.DegreeTitle,
			&_user.EducationHistory.Start,
			&_user.EducationHistory.End,
			&_user.EducationHistory.Institution,
			&_user.WorkHistory.Id,
			&_user.WorkHistory.Position,
			&_user.WorkHistory.Organization,
			&_user.WorkHistory.Start,
			&_user.WorkHistory.End)
		if err != nil {
			return fmt.Errorf("Cannot scan appropriate rows, err: %v", err)
		}
		usr = _user
		break
	}

	err = json.NewEncoder(w).Encode(usr)
	if err != nil {
		return fmt.Errorf("JSON encoding for GetUserById failed, err: %v", err)
	}
	return nil
}
