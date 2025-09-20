package database

import (
	"database/sql"
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

func TestGetIngredientByName(t *testing.T) {
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
		{
			Name: "Apple",
			Type: "Fruit",
			Cost: 1.2,
		},
		{
			Name: "Banana",
			Type: "Fruit",
			Cost: 2.45,
		},
	}
	tests := []struct {
		name               string
		nameToGet          string
		expectedIngredient *models.Ingredient
		expectedErrMsg     string
	}{
		{
			name:               "Successful Retrieval",
			nameToGet:          "STHL Vodka",
			expectedIngredient: ingredients[0],
			expectedErrMsg:     "",
		},
		{
			name:               "No Ingredient with Name",
			nameToGet:          "Rambutan",
			expectedIngredient: nil,
			expectedErrMsg:     sql.ErrNoRows.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := OpenDatabase(":memory:"); err != nil {
				panic(err)
			}
			defer CloseDatabase()

			for _, ingredient := range ingredients {
				err := DBInsertIngredient(ingredient)
				if err != nil {
					t.Errorf("Did not expect error when inserting %+v. Error: %s", ingredient, err)
				}
			}

			actualIngredient, err := DBGetIngredientByName(test.nameToGet)

			if test.expectedErrMsg != "" {
				assert.Equal(t, test.expectedErrMsg, err.Error())
			} else {
				assert.Equal(t, test.expectedIngredient, actualIngredient)
			}
		})
	}
}

func TestDBGetAllIngredients(t *testing.T) {
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
		{
			Name: "Apple",
			Type: "Fruit",
			Cost: 1.2,
		},
		{
			Name: "Banana",
			Type: "Fruit",
			Cost: 2.45,
		},
	}
	tests := []struct {
		name                   string
		expectedIngredientsNum int
	}{
		{
			name:                   "Successful Retrieval",
			expectedIngredientsNum: 4,
		},
		{
			name:                   "No rows",
			expectedIngredientsNum: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := OpenDatabase(":memory:"); err != nil {
				panic(err)
			}
			defer CloseDatabase()

			if test.name == "Successful Retrieval" {
				for _, ingredient := range ingredients {
					err := DBInsertIngredient(ingredient)
					if err != nil {
						t.Errorf("Did not expect error when inserting %+v. Error: %s", ingredient, err)
					}
				}
			}

			actualIngredients, err := DBGetAllIngredients()
			if err != nil {
				t.Errorf("Did not expect an error getting all ingredients. Error: %s", err)
			}

			assert.Equal(t, test.expectedIngredientsNum, len(actualIngredients))

		})
	}
}

func TestDBRemoveIngredientByName(t *testing.T) {
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
		{
			Name: "Apple",
			Type: "Fruit",
			Cost: 1.2,
		},
		{
			Name: "Banana",
			Type: "Fruit",
			Cost: 2.45,
		},
	}
	tests := []struct {
		name                 string
		nameToDelete         string
		remainingIngredients []models.Ingredient
		expectedErrMsg       string
	}{
		{
			name:         "Successful Deleteion",
			nameToDelete: "Apple",
			remainingIngredients: []models.Ingredient{
				*ingredients[0],
				*ingredients[1],
				*ingredients[3],
			},
			expectedErrMsg: "",
		},
		{
			name:         "No Ingredient with Name",
			nameToDelete: "Rambutan",
			remainingIngredients: []models.Ingredient{
				*ingredients[0],
				*ingredients[1],
				*ingredients[2],
				*ingredients[3],
			},
			expectedErrMsg: "expected to affect 1 row. affected 0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := OpenDatabase(":memory:"); err != nil {
				panic(err)
			}
			defer CloseDatabase()

			for _, ingredient := range ingredients {
				err := DBInsertIngredient(ingredient)
				if err != nil {
					t.Errorf("Did not expect error when inserting %+v. Error: %s", ingredient, err)
				}
			}

			err := DBRemoveIngredientByName(test.nameToDelete)

			if test.expectedErrMsg != "" {
				assert.Equal(t, test.expectedErrMsg, err.Error())
			}

			actualIngredients, err := DBGetAllIngredients()
			if err != nil {
				panic(err)
			}

			assert.Equal(t, test.remainingIngredients, actualIngredients)
		})
	}
}
