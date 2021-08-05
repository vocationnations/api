package model

// CultureDomain
type CultureDomain struct {
	Adhocracy string `json:"adhocracy"`
	Clan      string `json:"clan"`
	Hierarchy string `json:"hierarchy"`
	Market    string `json:"market"`
}

// Rank
type Rank struct {
	Rank1 string `json:"rank1"`
	Rank2 string `json:"rank2"`
	Rank3 string `json:"rank3"`
	Bin   string `json:"bin"`
}

// UserTypes
type UserTypes struct {
	JobSeeker string `json:"jobseeker"`
	Employer  string `json:"employer"`
	Admin     string `json:"admin"`
}
