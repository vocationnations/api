package model

import "fmt"

type UserType string

const (
	Admin     UserType = "Admin"
	JobSeeker UserType = "JobSeeker"
	Employer  UserType = "Employer"
)

type User struct {
	tableName struct{} `sql:"users"`
	Id        int      `json:"id"`
	Username  string   `json:"username"`
	UserType  UserType `json:"usertype"`
}

func (u User) String() string {
	return fmt.Sprintf("ID: <%d> Username: <%s> UserType: <%s>", u.Id, u.Username, u.UserType)
}
