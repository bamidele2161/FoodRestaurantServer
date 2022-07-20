package controllers

import (
	"FoodServer/entities"
	"FoodServer/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type categoryController struct { 
	categoryService services.Category
}

func NewCategoryController(category services.Category) *categoryController { 
	return &categoryController{categoryService: category}
}

func(c categoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var singleCategory entities.Category

	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &singleCategory)
	if err != nil {
		w.Write([]byte("Please supply user details in Json format"))
	}

		_, err = c.categoryService.CreateCategory(singleCategory) //passing the struct into the database for processing 

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		} // encode the struct into the json format
		
		w.WriteHeader(http.StatusCreated)
}


