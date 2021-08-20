package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryStatement_String(t *testing.T) {
	tests := []struct {
		name        string
		statement   string
		id          int
		domain      CultureDomain
		questionId  int
		expectedStr string
	}{
		{
			name:        "test returns expected string with proper data",
			id:          1,
			statement:   "test statement",
			domain:      "Hierarchy",
			questionId:  1,
			expectedStr: "ID: <1> Domain: <Hierarchy> Statement: <test statement> QuestionId: <1>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt := &CategoryStatement{Statement: tt.statement, Domain: tt.domain, Id: tt.id, QuestionId: tt.questionId}

			assert.Equal(t, tt.expectedStr, fmt.Sprint(stmt))
		})
	}
}
