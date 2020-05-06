package user

// User defines the structure of user data
type User struct {
	Id	string `json:"id"`
	Name string `json:"name"`
	EducationHistory Education `json:"education"`
	WorkHistory Work `json:"work"`
}

// Education defines the structure of education data
type Education struct {
	Id string `json:"id"`
	DegreeTitle string `json:"degreetitle"`
	Start string `json:"start"`
	End string `json:"end"`
	Institution string `json:"institution"`
}

type Work struct {
	Id string `json:"id"`
	Position string `json:"position"`
	Organization string `json:"organization"`
	Start string `json:"start"`
	End string `json:"end"`
}