package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_String(t *testing.T) {
	tests := []struct {
		testname        string
		id          int
		name    string
		email    string
		expectedStr string
	}{
		{
			testname:        "test returns expected string with proper data",
			id:          1,
			name:    "testuser1",
			email:    "testuser1@vocationations.com",
			expectedStr: "ID: <1> Name: <testuser1> Email: <testuser1@vocationations.com>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usr := &User{Id: tt.id, Name: tt.name, Email: tt.email}

			assert.Equal(t, tt.expectedStr, fmt.Sprint(usr))
		})
	}
}
