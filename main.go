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




type Login struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func main() {
	db := db.NewDatabase()
	err := db.StartDb()

	authService := services.NewAuth(db)
	authController := controllers.NewAuthController(*authService)

	if err != nil {
		panic(err)
	}
	defer db.Db.Close()
	err = db.Db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("we are connected to postgres")


	router := mux.NewRouter()

	// authentication endpoints
	router.HandleFunc("/api/v1/auth/register", authController.Register).Methods("POST")
	router.HandleFunc("/api/v1/auth/login", authController.Login).Methods("POST")

	// product endpoints
	// router.HandleFunc("/api/v1/admin", adminProduct).Methods("POST")
	// router.HandleFunc("/api/v1/allProducts", allProducts).Methods("GET")

	foodServer := http.Server{}
	foodServer.IdleTimeout = 30 * time.Second
	foodServer.Addr =  ":4000"
	foodServer.Handler = router
	foodServer.ListenAndServe()
}