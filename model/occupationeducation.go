package model

import "fmt"

type OccupationEducation struct {
	Code	string  `json:"code"`
	Report     string  `json:"report"`
	OLevelRequired   LevelRequired  `json:"level_required"`
}

type LevelRequired struct {
	OCategory []Category `json:"category"`
}

type Category struct {
	Name       string      `json:"name"`
	OScore    Score      `json:"score"`
}

type Score struct {
	Scale    string      `json:"scale"`
	Value       int      `json:"value"`
}

func (e OccupationEducation) String() string {
	return fmt.Sprintf("Code: <%s> Report: <%s> Level_Required: <%s>", e.Code, e.Report, e.OLevelRequired)
}

func (r LevelRequired) String() string {
	return fmt.Sprintf("Category: <%s>", r.OCategory)
}

func (c Category) String() string {
	return fmt.Sprintf("Name: <%s> Score: <%s>", c.Name, c.OScore)
}

func (s Score) String() string {
	return fmt.Sprintf("Scale: <%s> Value: <%d>", s.Scale, s.Value)
}