package model

import "fmt"

type UsersProfessions struct {
	tableName    struct{} `pg:"usersprofessions"`
	UserId       int      `json:"user_id"`
	Title        string   `json:"title"`
	ProfessionId string   `json:"profession_id"`
}

func (u UsersProfessions) String() string {
	return fmt.Sprintf("ID: <%d> Title: <%s> ProfessionId: <%s>", u.UserId, u.Title, u.ProfessionId)
}
