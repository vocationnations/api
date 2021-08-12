package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_String(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		username    string
		usertype    UserType
		expectedStr string
	}{
		{
			name:        "test returns expected string with proper data",
			id:          1,
			username:    "testuser1",
			usertype:    "Admin",
			expectedStr: "ID: <1> Username: <testuser1> UserType: <Admin>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usr := &User{Id: tt.id, Username: tt.username, UserType: tt.usertype}

			assert.Equal(t, tt.expectedStr, fmt.Sprint(usr))
		})
	}
}
