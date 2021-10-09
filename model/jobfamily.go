package model

import "fmt"
type JobFamilies struct {
	JobFamily []JobFamily `json:"job_family"`
}

type JobFamily struct {
	Href    string      `json:"href"`
	Code    string      `json:"code"`
	Title   string      `json:"title"`
}

func (jf JobFamily) String() string {
	return fmt.Sprintf("Href: <%s> Code: <%s> Title: <%s>", jf.Href, jf.Code, jf.Title)
}
