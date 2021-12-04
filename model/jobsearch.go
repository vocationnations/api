package model

import "fmt"

type KeywordWrapper struct {
	Keyword     string       `json:"keyword"`
	Start      int          `json:"start"`
	End        int          `json:"end"`
	Total      int          `json:"total"`
	OLinks     []Link       `json:"link"`
	Occupation []OccupationKeyword `json:"occupation"`
}

type OccupationKeyword struct {
	Href       string              `json:"href"`
	RelevanceScore      int          `json:"relevance_score"`
	Code       string              `json:"code"`
	Title      string              `json:"title"`
	Tags       Tag                 `json:"tags"`
}


func (key KeywordWrapper) String() string {
	return fmt.Sprintf(
		"Keyword: <%s> "+
			"Start: <%d> "+
			"End: <%d> "+
			"Total: <%d> "+
			"OLinks: <%s> "+
			"Occupation: <%s>", key.Keyword, key.Start, key.End, key.Total, key.OLinks, key.Occupation)
}

func (o OccupationKeyword) String() string {
	return fmt.Sprintf("Href: <%s> RelevanceScore: <%d> Code: <%s> Title: <%s> Tags: [BrightOutlook: <%t> Green: <%t>]", o.Href, o.RelevanceScore, o.Code, o.Title, o.Tags.BrightOutlook, o.Tags.Green)
}