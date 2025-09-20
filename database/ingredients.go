package database

import (
	"fmt"

	"github.com/Shizu-99/recipe-telegram/models"
)

func DBInsertIngredient(ingredient *models.Ingredient) error {
	_, err := db.NamedExec(`INSERT INTO ingredients (name, type, cost) VALUES (:name, :type, :cost)`, ingredient)
	if err != nil {
		return err
	}
	return nil
}

func DBGetIngredientByName(name string) (*models.Ingredient, error) {
	ingredient := &models.Ingredient{}
	err := db.Get(ingredient, `SELECT name, type, cost FROM ingredients WHERE name=$1`, name)
	if err != nil {
		return nil, err
	}
	return ingredient, nil
}

func DBGetAllIngredients() ([]models.Ingredient, error) {
	ingredients := []models.Ingredient{}
	err := db.Select(&ingredients, `SELECT name, type, cost FROM ingredients`)
	if err != nil {
		return nil, err
	}
	return ingredients, nil
}

func DBRemoveIngredientByName(name string) error {
	res, err := db.Exec(`DELETE FROM ingredients WHERE name=$1`, name)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row. affected %d", rows)
	}
	return nil
}
