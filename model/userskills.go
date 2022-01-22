package model

import "fmt"

type UserSkills struct {
	tableName     struct{} `pg:"usersskills"`
	Id            int      `json:"id"`
	User_id 	  int      `json:"user_id"`
	Skill_id	  string   `json:"skill_id"`
	Skill_name    string   `json:"skill_name"`
}


func (s UserSkills) String() string {
	return fmt.Sprintf("ID: <%d> User_id: <%d> Skill_id: <%s> Skill_name: <%s>", s.Id, s.User_id, s.Skill_id, s.Skill_name)
}
