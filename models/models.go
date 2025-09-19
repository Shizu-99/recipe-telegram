package models

type Ingredient struct {
	ID   int     `db:"ingredient_id" json:"ingredient_id"`
	Name string  `db:"name" json:"name"`
	Type string  `db:"type" json:"type"`
	Cost float64 `db:"cost" json:"cost"`
}
