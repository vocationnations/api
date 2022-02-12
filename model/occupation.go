package model

import "fmt"

type OccupationWrapper struct {
	Family     string       `json:"family"`
	Sort       string       `json:"sort"`
	Start      int          `json:"start"`
	End        int          `json:"end"`
	Total      int          `json:"total"`
	OLinks     []Link       `json:"link"`
	Occupation []Occupation `json:"occupation"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type Occupation struct {
	Href       string              `json:"href"`
	Code       string              `json:"code"`
	Title      string              `json:"title"`
	Tags       Tag                 `json:"tags"`
	OJobFamily OccupationJobFamily `json:"job_family"`
}

type Tag struct {
	BrightOutlook bool `json:"bright_outlook"`
	Green         bool `json:"green"`
}

type OccupationJobFamily struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (o OccupationWrapper) String() string {
	return fmt.Sprintf(
		"Family: <%s> "+
			"Sort: <%s> "+
			"Start: <%d> "+
			"End: <%d> "+
			"Total: <%d> "+
			"OLinks: <%s> "+
			"Occupation: <%s>", o.Family, o.Sort, o.Start, o.End, o.Total, o.OLinks, o.Occupation)
}

func (l Link) String() string {
	return fmt.Sprintf("Href: <%s> Rel: <%s>", l.Href, l.Rel)
}

func (o Occupation) String() string {
	return fmt.Sprintf("Href: <%s> Code: <%s> Label: <%s> Tags: [BrightOutlook: <%t> Green: <%t>] OccupationJobFamily: <%s>", o.Href, o.Code, o.Title, o.Tags.BrightOutlook, o.Tags.Green, o.OJobFamily)
}

func (ojf OccupationJobFamily) String() string {
	return fmt.Sprintf("Code: <%s> Name: <%s>", ojf.Code, ojf.Name)
}
