package model


import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSkill_String(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		skill_name  string
		value        int
		user_entry_id   int
		expectedStr string
	}{
		{
			name:        "test returns expected string with proper data",
			id:          1,
			skill_name:    "java",
			value:    50,
			user_entry_id: 1,
			expectedStr: "ID: <1> Skill_name: <java> Value: <50> User_entry_id: <1>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skill := &Skill{Id: tt.id, Skill_name: tt.skill_name, Value: tt.value, User_entry_id: tt.user_entry_id}

			assert.Equal(t, tt.expectedStr, fmt.Sprint(skill))
		})
	}
}

