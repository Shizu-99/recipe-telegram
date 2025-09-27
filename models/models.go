package models

import "image"

type Ingredient struct {
	Name string  `db:"name" json:"name"`
	Type string  `db:"type" json:"type"`
	Cost float64 `db:"cost" json:"cost"`
}

type Recipe struct {
	Name        string       `db:"name" json:"name"`
	Ingredients []Ingredient `db:"ingredients" json:"ingredients"`
	Method      string       `db:"method" json:"method"`
	Source      string       `db:"source" json:"source"`
	RecipeImage image.Image  `db:"recipe_image" json:"recipe_image"`
}
