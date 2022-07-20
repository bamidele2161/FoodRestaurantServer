package controllers

import (
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	"net/http"
)


func admin(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Welcome to my admin page")
	// if r.Header.Get("content-type") == "application/json" {
	// 	var singleProduct Product
	// 	Products := make([]Product, 0)
	// 	reqBody, _ := ioutil.ReadAll(r.Body) // read all the data received
	// 	json.Unmarshal(reqBody, &singleProduct) //decode from json to struct
	// 	Products = append(Products, singleProduct) // append the new user to the list of users
	// 	json.NewEncoder(w).Encode(singleProduct) //convert struct to json 
	// 	w.WriteHeader(http.StatusCreated)
	// 	w.Write([]byte("Product Added Successfully"))
	// } else {
	// 	w.WriteHeader(http.StatusUnprocessableEntity)
	// 	w.Write([]byte("Please supply user details in Json format"))
	// }
}

func products(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Welcome to products")
	// kv := r.URL.Query() //this returns the key value pair 
    // for key, value := range kv {
    //     fmt.Println(key, value)
    // }
    // // returns all the courses in JSON
    // json.NewEncoder(w).Encode(Products)
}
