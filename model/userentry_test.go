package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserEntry_String(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		date        time.Time
		userId		int
		expectedStr string
	}{
		{
			name:        "test returns expected string with proper data",
			id:          1,
			date:        time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
			userId:		3,
			expectedStr: "ID: <1> Date: <2009-11-17 20:34:58.651387237 +0000 UTC> UserId: <3>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usr := &UserEntry{Id: tt.id, Date: &tt.date}

			assert.Equal(t, tt.expectedStr, fmt.Sprint(usr))
		})
	}
}
