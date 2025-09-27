package models

import "image"

type Recipe struct {
	Name        string      `db:"name" json:"name"`
	Ingredients []string    `db:"ingredients" json:"ingredients"`
	Method      string      `db:"method" json:"method"`
	Source      string      `db:"source" json:"source"`
	RecipeImage image.Image `db:"recipe_image" json:"recipe_image"`
}
