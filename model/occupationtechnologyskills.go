package model

import "fmt"

type OccupationTechnologySkills struct {
	Code	string  `json:"code"`
	Report	string  `json:"report"`
	Display	string  `json:"display"`
	Category	[]SkillCategory `json:"category"`
}

type SkillCategory struct {
	Related  string `json:"related"`
	Title Title  `json:"title"`
	Example []Example  `json:"example"`
}

type Title struct {
	Id  int `json:"id"`
	HotTechnology int  `json:"hot_technology"`
	Name string  `json:"name"`
}

type Example struct {
	HotTechnology int  `json:"hot_technology"`
	Name string  `json:"name"`
}


func (t OccupationTechnologySkills) String() string {
	return fmt.Sprintf("Code: <%s> Report: <%s> Display: <%s> Category: <%s>", t.Code, t.Report, t.Display, t.Category)
}


func (e Title) String() string {
	return fmt.Sprintf("Id: <%d> Hot Technology: <%d> Name: <%s>", e.Id, e.HotTechnology, e.Name)
}


func (e Example) String() string {
	return fmt.Sprintf("Hot Technology: <%d> Name: <%s>", e.HotTechnology, e.Name)
}