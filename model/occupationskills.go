package model

import "fmt"

type OccupationSkills struct {
	Code	string  `json:"code"`
	Report	string  `json:"report"`
	Display	string  `json:"display"`
	Elements	[]Element `json:"element"`
}

type Element struct {
	Id string `json:"id"`
	Related  string `json:"related"`
	Name string `json:"name"`
	Description  string `json:"description"`
}

func (s OccupationSkills) String() string {
	return fmt.Sprintf("Code: <%s> Report: <%s> Display: <%d> Elements: <%d>", s.Code, s.Report, s.Display, s.Elements)
}


func (e Element) String() string {
	return fmt.Sprintf("Id: <%s> Related: <%s> Name: <%s> Description: <%s>", e.Id, e.Related, e.Name, e.Description)
}