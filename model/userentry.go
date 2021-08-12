package model

import (
	"fmt"
	"time"
)

type UserEntry struct {
	tableName struct{} `sql:"user_entries"`
	Id        int      `json:"id"`
	Date      *time.Time
}

func (ue UserEntry) String() string {
	return fmt.Sprintf("ID: <%d> Date: <%s>", ue.Id, ue.Date)
}
