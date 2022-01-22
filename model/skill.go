package model

import "fmt"

type SkillManual struct {
	tableName struct{} `pg:"user_skills_manual"`
	Id int `json:"id"`
	SkillName string `json:"skill_name"`
	UserId int `json:"user_id"`
}

type Skill struct {
	tableName     struct{} `pg:"skills"`
	Id            int      `json:"id"`
	Skill_name    string   `json:"skill_name"`
	Value         int      `json:"value"`
	User_entry_id int      `json:"user_entry_id"`
}

func (s Skill) String() string {
	return fmt.Sprintf("ID: <%d> Skill_name: <%s> Value: <%d> User_entry_id: <%d>", s.Id, s.Skill_name, s.Value, s.User_entry_id)
}
