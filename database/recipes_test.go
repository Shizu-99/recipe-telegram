package database

import (
	"image"
	"testing"

	"github.com/Shizu-99/recipe-telegram/models"
	"github.com/stretchr/testify/assert"
)

// Doesn't work
func TestDBInsertRecipe(t *testing.T) {
	tests := []struct {
		name               string
		recipesToInsert    []*models.Recipe
		expectedNumRecipes int
		expectErr          bool
	}{
		{
			name: "Successful Insertion",
			recipesToInsert: []*models.Recipe{
				{
					Name: "Mocktail",
					Ingredients: []string{
						"Lemonade",
						"Raspberry Cordial",
					},
					Method:      "Pour ingredients into glass",
					Source:      "Brain",
					RecipeImage: image.White,
				},
			},
			expectedNumRecipes: 1,
			expectErr:          false,
		},
		{
			name: "Duplicate Name",
			recipesToInsert: []*models.Recipe{
				{
					Name: "Mocktail",
					Ingredients: []string{
						"Lemonade",
						"Raspberry Cordial",
					},
					Method:      "Pour ingredients into glass",
					Source:      "Brain",
					RecipeImage: image.White,
				},
				{
					Name: "Mocktail",
					Ingredients: []string{
						"Lemonade",
						"Raspberry Cordial",
					},
					Method:      "Pour ingredients into glass",
					Source:      "Brain",
					RecipeImage: image.White,
				},
			},
			expectedNumRecipes: 1,
			expectErr:          true,
		},
		{
			name: "No Ingredients",
			recipesToInsert: []*models.Recipe{
				{
					Name:        "Mocktail",
					Ingredients: []string{},
					Method:      "Pour ingredients into glass",
					Source:      "",
					RecipeImage: nil,
				},
			},
			expectedNumRecipes: 0,
			expectErr:          true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := OpenDatabase(":memory:"); err != nil {
				panic(err)
			}
			defer CloseDatabase()

			var err error

			err = DBInsertRecipe(test.recipesToInsert[0])
			if test.name == "Duplicate Name" {
				err = DBInsertRecipe(test.recipesToInsert[1])
			}

			if test.expectErr {
				assert.Error(t, err)
			}

			recipes, err := DBGetAllRecipes()
			if err != nil {
				t.Errorf("did not expect getting all recipes to error. Error: %s", err.Error())
			}
			assert.Equal(t, test.expectedNumRecipes, len(recipes))
		})
	}
}
