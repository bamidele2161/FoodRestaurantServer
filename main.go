package main

import (
	"FoodServer/controllers"
	"FoodServer/db"
	"FoodServer/services"
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

func main() {
	db := db.NewDatabase()
	err := db.StartDb()

	authService := services.NewAuth(db)
	authController := controllers.NewAuthController(*authService)

	categoryService := services.NewCategory(db)
	categoryController := controllers.NewCategoryController(*categoryService)

	productService := services.NewProduct(db)
	productController := controllers.NewProductController(*productService)

	if err != nil {
		panic(err)
	}
	defer db.Db.Close() //ensuring the db is properly closed when not needed
	err = db.Db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("we are connected to postgres")


	router := mux.NewRouter()

	// authentication endpoints
	router.HandleFunc("/api/v1/auth/register", authController.Register).Methods("POST")
	
	router.HandleFunc("/api/v1/auth/login", authController.Login).Methods("POST")
	// router.HandleFunc("/api/v1/auth/profile", authController.Profile).Methods("GET")

	// product endpoints
	router.HandleFunc("/api/v1/allproducts", productController.GetProduct).Methods("GET")
	router.HandleFunc("/api/v1/product", productController.CreateProduct).Methods("POST")

	//category endpoints
	router.HandleFunc("/api/v1/category", categoryController.CreateCategory).Methods("POST")


	foodServer := http.Server{}
	foodServer.IdleTimeout = 30 * time.Second
	foodServer.Addr =  ":4000"
	foodServer.Handler = router
	foodServer.ListenAndServe()
}