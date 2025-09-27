package database

import (
	"fmt"

	"github.com/Shizu-99/recipe-telegram/models"
)

func DBInsertRecipe(recipe *models.Recipe) error {
	_, err := db.NamedExec(`INSERT INTO recipes (name, ingredients, method, source, recipe_image) VALUES (:name, :ingredients, :method, :source, :recipe_image)`, recipe)
	if err != nil {
		return err
	}
	return nil
}

func DBGetRecipeByName(name string) (*models.Recipe, error) {
	recipe := &models.Recipe{}
	err := db.Get(recipe, `SELECT name, ingredients, method, source, recipe_image FROM recipes WHERE name=$1`, name)
	if err != nil {
		return nil, err
	}
	return recipe, nil
}

func DBGetAllRecipes() ([]models.Recipe, error) {
	recipes := []models.Recipe{}
	err := db.Select(&recipes, `SELECT name, ingredients, method, source, recipe_image FROM recipes`)
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func DBDeleteRecipeByName(name string) error {
	res, err := db.Exec(`DELETE FROM recipes WHERE name=$1`, name)
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
