package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCultureCategory_String(t *testing.T) {
	tests := []struct {
		name          string
		id            int
		categoryTitle string
		expectedStr   string
	}{
		{
			name:          "test returns expected string with proper data",
			id:            1,
			categoryTitle: "testcategory",
			expectedStr:   "ID: <1> Category Title: <testcategory>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctgry := &CultureCategory{Id: tt.id, CategoryTitle: tt.categoryTitle}

			assert.Equal(t, tt.expectedStr, fmt.Sprint(ctgry))
		})
	}
}
