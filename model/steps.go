package model

import "fmt"

type Steps struct {
	tableName struct{} `pg:"steps"`
	Id        int      `json:"id"`
	Title  string   `json:"title"`
	Value  int   `json:"value"`
	QuestionId  int   `json:"question_id"`
}


func (s Steps) String() string {
	return fmt.Sprintf("ID: <%d> Title: <%s> Value: <%d> QuestionId: <%d>", s.Id, s.Title, s.Value, s.QuestionId)
}
