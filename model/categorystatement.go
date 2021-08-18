package model

import "fmt"

type CultureDomain string

const (
	Adhocracy CultureDomain = "Adhocracy"
	Clan      CultureDomain = "Clan"
	Hierarchy CultureDomain = "Hierarchy"
	Market    CultureDomain = "Market"
)

type CategoryStatement struct {
	tableName  struct{}      `sql:"category_statements"`
	Id         int           `json:"id"`
	Domain     CultureDomain `json:"domain"`
	Statement  string        `json:"statement"`
	QuestionId int           `json:"question_id"`
}

func (cs CategoryStatement) String() string {
	return fmt.Sprintf("ID: <%d> Domain: <%s> Statement: <%s> QuestionId: <%d>", cs.Id, cs.Domain, cs.Statement, cs.QuestionId)
}
