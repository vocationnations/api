package model

import "fmt"

type CultureCategory struct {
	tableName     struct{} `sql:"culture_categories"`
	Id            int      `json:"id"`
	CategoryTitle string   `json:"category_title"`
}

func (cc CultureCategory) String() string {
	return fmt.Sprintf("ID: <%d> Category Title: <%s>", cc.Id, cc.CategoryTitle)
}
