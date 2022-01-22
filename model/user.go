package model

import "fmt"

type UserType string

type User struct {
	tableName struct{} `sql:"users"`
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
}

func (u User) String() string {
	return fmt.Sprintf("ID: <%d> Name: <%s> Email: <%s>", u.Id, u.Name, u.Email)
}
