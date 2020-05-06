package user

import (
	"database/sql"
	"fmt"
)

func GetUserById(db *sql.DB, id string) User {
	rows, err := db.Query(`
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
		fmt.Println("Cannot select user data, err: ", err)
	}

	for rows.Next(){
		var _user User

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
			fmt.Println("Cannot scan appropriate rows, err: ", err)
		}

		return _user
	}
	return User{}
}
