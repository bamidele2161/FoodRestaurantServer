package controllers

import (
	"FoodServer/entities"
	"FoodServer/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type productController struct {
	productService services.Product
}

func NewProductController(product services.Product) *productController {
	return &productController{productService: product}
}

func(p productController) CreateProduct(w http.ResponseWriter, r *http.Request) {

		var singleProduct entities.Product

		reqBody, _ := ioutil.ReadAll(r.Body) // read all the data received
		err := json.Unmarshal(reqBody, &singleProduct) //decode from json to struct
		
		if err != nil {
			w.Write([]byte("Please supply user details in Json format"))
		}
		createdProduct, err := p.productService.CreateProduct(singleProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		productJson, err := json.Marshal(createdProduct)
		w.WriteHeader(http.StatusCreated)
		w.Write(productJson)

}

func(p productController) GetProduct(w http.ResponseWriter, r *http.Request) {
	var product entities.Product
	AllProducts, err := p.productService.GetProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	getProducts, err := json.Marshal(AllProducts)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(getProducts)
}
