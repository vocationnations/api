package model

// Occupation defines the structure of Occupation
type Occupation struct {
	Onetcode    string `json:"onetcode"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// User defines the structure of user data
type User struct {
	Id               string    `json:"id"`
	Name             string    `json:"name"`
	EducationHistory Education `json:"education"`
	WorkHistory      Work      `json:"work"`
}

// Education defines the structure of education data
type Education struct {
	Id          string `json:"id"`
	DegreeTitle string `json:"degreetitle"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Institution string `json:"institution"`
}

// Work defines the structure of user work data
type Work struct {
	Id           string `json:"id"`
	Position     string `json:"position"`
	Organization string `json:"organization"`
	Start        string `json:"start"`
	End          string `json:"end"`
}
