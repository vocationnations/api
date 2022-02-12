package model

import "fmt"

type Steps struct {
	tableName  struct{} `pg:"steps"`
	Id         int      `json:"id"`
	Label      string   `json:"label"`
	Number     int      `json:"number"`
	QuestionId int      `json:"question_id"`
}

func (s Steps) String() string {
	return fmt.Sprintf("ID: <%d> Label: <%s> Number: <%d> QuestionId: <%d>", s.Id, s.Label, s.Number, s.QuestionId)
}
