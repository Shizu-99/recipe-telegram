package models

type Ingredient struct {
	Name string  `db:"name" json:"name"`
	Type string  `db:"type" json:"type"`
	Cost float64 `db:"cost" json:"cost"`
}
