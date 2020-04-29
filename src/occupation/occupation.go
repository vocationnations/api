package occupation

import (
	"database/sql"
	"fmt"
)

func AddOccupation(db *sql.DB, onetcode string, title string, description string) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare(`INSERT INTO occupation_data (onetcode, title, description) VALUES (?,?,?)`)
	_, err := stmt.Exec(onetcode, title, description)
	if err != nil {
		fmt.Println("Cannot add Occupation, err: ", err)
	}
	tx.Commit()
}

func GetOccupations(db *sql.DB) []Occupation {

	var AllOccupations []Occupation
	rows, err := db.Query(`SELECT * FROM occupation_data`)
	if err != nil {
		fmt.Println("Cannot select occupations, err: ", err)
	}

	for rows.Next() {
		var _occupation Occupation
		err = rows.Scan(&_occupation.Onetcode, &_occupation.Title, &_occupation.Description)
		if err != nil {
			fmt.Println("Cannot scan appropriate rows, err: ", err)
		}
		AllOccupations = append(AllOccupations, _occupation)
	}
	return AllOccupations
}

func GetOccupationById(db *sql.DB, onetcode string) Occupation {
	rows, err := db.Query(`SELECT * FROM occupation_data WHERE onetcode = ?`, onetcode)
	if err != nil {
		fmt.Println("Cannot select occupations, err: ", err)
	}

	for rows.Next() {
		var _occupation Occupation

		err = rows.Scan(&_occupation.Onetcode, &_occupation.Title, &_occupation.Description)
		if err != nil {
			fmt.Println("Cannot scan appropriate rows, err: ", err)
		}

		if _occupation.Onetcode == onetcode {
			return _occupation
		}
	}
	return Occupation{}
}

func GetOccupationsByTitle(db *sql.DB, title string) []Occupation {
	search_title := "%" + title + "%"
	rows, err := db.Query(`SELECT * FROM occupation_data WHERE title LIKE ?`, search_title)
	if err != nil {
		fmt.Println("Cannot select occupations with title, err: ", err)
	}

	var occs []Occupation
	for rows.Next() {
		var _occupation Occupation

		err = rows.Scan(&_occupation.Onetcode, &_occupation.Title, &_occupation.Description)
		if err != nil {
			fmt.Println("Cannot scan appropriate rows, err: ", err)
		}
		occs = append(occs, _occupation)
	}
	return occs
}
