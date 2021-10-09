package model

import (
	"fmt"
	"time"
)

type UserEntry struct {
	tableName struct{} `sql:"user_entries"`
	Id        int      `json:"id"`
	Date      *time.Time
	UserId    int `json:"user_id"`
}

func (ue UserEntry) String() string {
	return fmt.Sprintf("ID: <%d> Date: <%s> UserId: <%d>", ue.Id, ue.Date, ue.UserId)
}
