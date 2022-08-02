package services

import (
	"FoodServer/db"
	"FoodServer/entities"
	"errors"
	"fmt"
)

type Product struct {
	foodServerDb *db.Database
}

func NewProduct(foodServerDb *db.Database) *Product {
	return &Product{foodServerDb: foodServerDb}
}

func(p Product) CreateProduct(product entities.Product) (entities.Product, error) { 
	rows, _ := p.foodServerDb.Db.Exec("select * from product where product_name = $1", product.ProductName)

	nRows, err := rows.RowsAffected()
	if nRows == 0 {
		_, err := p.foodServerDb.Db.Exec(`INSERT INTO product (category_id, product_name, product_price, product_rating, product_description, product_image) VALUES ($1, $2, $3, $4, $5, $6)`, product.CategoryID, product.ProductName, product.ProductPrice, product.Rating, product.Description, product.Image)
		if err != nil {
			fmt.Println(err)
			return entities.Product{}, errors.New("An error occurred while creating product")
		}
		createdProduct := entities.Product{}
		row := p.foodServerDb.Db.QueryRow("SELECT * FROM product WHERE product_name = $1", product.ProductName)
		row.Scan(&createdProduct.Id, &createdProduct.CategoryID, &createdProduct.ProductName, &createdProduct.ProductPrice, &createdProduct.Rating, &createdProduct.Description, &createdProduct.Image)
		return createdProduct, nil
	}
	fmt.Println(err)
return entities.Product{}, errors.New("Product already exists")
}

func(p Product) GetProduct(product entities.Product) ([]entities.Product, error) {
	rows, err := p.foodServerDb.Db.Query("SELECT * FROM product")
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("An error occurred while retrieving product")
	}
	
	defer rows.Close()
	lists := []entities.Product{}
	for rows.Next() {
		getProduct := entities.Product{}
		err = rows.Scan(&getProduct.Id, &getProduct.CategoryID, &getProduct.ProductName, &getProduct.ProductPrice, &getProduct.Rating, &getProduct.Description, &getProduct.Image)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("An error occurred while retrieving product")
		}
		lists = append(lists, getProduct)
		fmt.Println(lists)

	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	  }

	return lists, nil
}