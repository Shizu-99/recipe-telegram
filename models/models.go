package models

type Ingredient struct {
	ID   int     `db:"ingredient_id" json:"ingredient_id"`
	Name string  `db:"name" json:"name"`
	Cost float64 `db:"cost" json:"cost"`
}
