package controllers

import (
	"FoodServer/entities"
	"FoodServer/helpers"
	"FoodServer/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type authController struct {
	authservice services.Auth
}

func NewAuthController(auth services.Auth) *authController{
	return &authController{authservice: auth}
	
}

func(a authController) Register(w http.ResponseWriter, r *http.Request) {
		var singleUser entities.User
		emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
		reqBody, _ := ioutil.ReadAll(r.Body) // read all the data received
		err := json.Unmarshal(reqBody, &singleUser) //decode from json to struct

		if err != nil {
			w.Write([]byte("Please supply user details in Json format"))
		}

		if !emailRegex.MatchString(singleUser.Email) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Please provide a valid email address"))
		}else if len(singleUser.Password) < 6 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" Password must be at least 6 characters"))
		} else {
			password, err := helpers.HashPassword(singleUser.Password, 6)
			if err != nil {
				fmt.Println(err)
			}
			singleUser.Password = string(password)

			createdUser, err := a.authservice.CreateUser(singleUser)  //passing the struct into the database for processing 
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			} 
			userJson, err := json.Marshal(createdUser) // encode the struct into the json format
			
			w.WriteHeader(http.StatusCreated)
			w.Write(userJson)
		}
		
}



func(a authController) Login(w http.ResponseWriter, r *http.Request) {
	var userDetail entities.Login
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &userDetail)
	if err != nil {
		w.Write([]byte("Please supply user details in Json format"))
	}
	loginUser, err := a.authservice.LoginUser(userDetail, []byte(userDetail.Password))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		} 
		
		loginJson, err := json.Marshal(loginUser)

		w.WriteHeader(http.StatusOK)
		w.Write(loginJson)
}
