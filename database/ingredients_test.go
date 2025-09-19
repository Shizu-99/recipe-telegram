package database

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Shizu-99/recipe-telegram/models"
)

func TestDBInsertIngredient(t *testing.T) {
	ingredients := []*models.Ingredient{
		{
			Name: "STHL Vodka",
			Type: "Vodka",
			Cost: 45.34,
		},
		{
			Name: "Fakerson Vodka",
			Type: "Vodka",
			Cost: 32.0,
		},
	}
	tests := []struct {
		name                string
		ingredientsToInsert []*models.Ingredient
		expectedDBEntries   int
		expectedErrMsg      string
	}{
		{
			name:                "Successful Insertion",
			ingredientsToInsert: ingredients,
			expectedDBEntries:   2,
			expectedErrMsg:      "",
		},
		{
			name:                "Duplicate Name",
			ingredientsToInsert: ingredients,
			expectedDBEntries:   2,
			expectedErrMsg:      "UNIQUE constraint failed: ingredients.name",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := OpenDatabase(":memory:"); err != nil {
				panic(err)
			}
			defer CloseDatabase()

			for _, ingredient := range test.ingredientsToInsert {
				err := DBInsertIngredient(ingredient)
				if err != nil {
					t.Errorf("Did not expect error when inserting %+v. Error: %s", ingredient, err)
				}
			}

			if test.name != "Successful Insertion" {
				err := DBInsertIngredient(test.ingredientsToInsert[0])
				if test.expectedErrMsg != "" {
					assert.Equal(t, test.expectedErrMsg, err.Error())
				}
			}

			entries, err := DBGetAllIngredients()
			if err != nil {
				panic(err)
			}

			assert.Equal(t, test.expectedDBEntries, len(entries))
		})
	}
}
