package models

type Ingredient struct {
	Name string  `db:"name" json:"name"`
	Cost float64 `db:"cost" json:"cost"`
}
