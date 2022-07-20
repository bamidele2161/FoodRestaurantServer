package controllers

import (
	"FoodServer/entities"
	"FoodServer/services"
	"encoding/json"
	"io/ioutil"
	"net/http"

)

type authController struct {
	authservice services.Auth
}

func NewAuthController(auth services.Auth) *authController{
	return &authController{authservice: auth}
}

func(a authController) Register(w http.ResponseWriter, r *http.Request) {
		var singleUser entities.User

		reqBody, _ := ioutil.ReadAll(r.Body) // read all the data received
		err := json.Unmarshal(reqBody, &singleUser) //decode from json to struct

		if err != nil {
			w.Write([]byte("Please supply user details in Json format"))
		}

		createdUser, err := a.authservice.CreateUser(singleUser) //passing the struct into the database for processing 

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		} 
		userJson, err := json.Marshal(createdUser) // encode the struct into the json format
		
		w.WriteHeader(http.StatusCreated)
		w.Write(userJson)
}



func(a authController) Login(w http.ResponseWriter, r *http.Request) {
	var userDetail entities.Login
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &userDetail)
	loginUser, err := a.authservice.LoginUser(userDetail)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		} 
		loginJson, err := json.Marshal(loginUser)

		w.WriteHeader(http.StatusOK)
		w.Write(loginJson)
}