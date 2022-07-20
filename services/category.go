package services

import (
	"FoodServer/db"
	"FoodServer/entities"
	"errors"
	"fmt"
)

type Category struct {
	foodServerDb *db.Database
}

func NewCategory(foodServerDb *db.Database) *Category {
	return &Category{foodServerDb: foodServerDb}
}

func(c Category) CreateCategory(category entities.Category) (bool, error) {
	rows, _ := c.foodServerDb.Db.Exec("SELECT * FROM category WHERE category_name = $1", category.Name)

	nRows, _ := rows.RowsAffected()
	if nRows == 0 {
		_, err := c.foodServerDb.Db.Exec(`Insert into category (category_name) values ($1)`, category.Name)
		if err != nil {
			fmt.Println(err.Error())
			return false, errors.New("An error occurred while inserting category")
		}
		
		return true, nil
	}
	return false, errors.New("Category already exists")
}