package model

import "fmt"

type ClimateQuestionAnswers struct {
	tableName struct{} `pg:"steps"`
	Id        int      `json:"id"`
	QuestionId  int   `json:"question_id"`
	StepId  int   `json:"step_id"`
	UserId  int   `json:"user_id"`
}


func (cqa ClimateQuestionAnswers) String() string {
	return fmt.Sprintf("ID: <%d> QuestionId: <%d> StepId: <%d> UserId: <%d>", cqa.Id, cqa.QuestionId, cqa.StepId, cqa.UserId)
}
