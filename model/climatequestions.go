package model

import "fmt"

type ClimateQuestions struct {
	tableName struct{} `pg:"climatequestions"`
	Id        int      `json:"id"`
	Title  string   `json:"title"`
}

func (cq ClimateQuestions) String() string {
	return fmt.Sprintf("ID: <%d> Title: <%s>", cq.Id, cq.Title)
}
