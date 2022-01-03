package model

import "fmt"

type OccupationTasks struct {
	Code	string  `json:"code"`
	Start      int  `json:"start"`
	End        int  `json:"end"`
	Total      int  `json:"total"`
	Tasks	[]Task  `json:"task"`
}

type Task struct {
	Id 		   int  `json:"id"`
	Green     bool  `json:"green"`
	Name 	string  `json:"name"`
}


func (t OccupationTasks) String() string {
	return fmt.Sprintf("Code: <%s> Start: <%d> End: <%d> Total: <%d> Tasks: <%s>", t.Code, t.Start, t.End, t.Total, t.Tasks)
}

func (o Task) String() string {
	return fmt.Sprintf("Id: <%d> Green: <%t> Name: <%s>", o.Id, o.Green, o.Name)
}