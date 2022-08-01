package services

import (
	"FoodServer/db"
	"FoodServer/entities"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	foodServerDb *db.Database
}

func NewAuth(foodServerDb *db.Database) *Auth {
	return &Auth{foodServerDb: foodServerDb}
}

// CreateUser creates a new user in the database.
// Returns an error if the user already exists
func(a Auth) CreateUser(user entities.User) (entities.User, error) {
	rows, _ := a.foodServerDb.Db.Exec("SELECT * from users where user_email = $1 or user_phone = $2", user.Email, user.Phone);
	
	// rows affected 
	nRows, err := rows.RowsAffected()
	if nRows == 0{
		 _, err := a.foodServerDb.Db.Exec(`Insert into users (user_firstname, user_lastname, user_email,  user_phone, user_password) values ($1, $2, $3, $4, $5) returning user_id`, user.FirstName, user.LastName, user.Email, user.Phone, user.Password)
		if err != nil {
			fmt.Println(err)
			return entities.User{}, errors.New("An error occurred while creating user")
		}

		if err != nil {
			return entities.User{}, errors.New("An error occurred while creating user")
		}
		createdUser := entities.User{}

		row := a.foodServerDb.Db.QueryRow("SELECT * FROM users WHERE user_email = $1", user.Email) 
		row.Scan(&createdUser.Id, &createdUser.FirstName, &createdUser.LastName, &createdUser.Email, &createdUser.Phone, &createdUser.Password)

		return createdUser, nil
	}
	fmt.Println(err)
	return entities.User{}, errors.New("User already exists")
}

func(a Auth) LoginUser(login entities.Login, password []byte) (entities.User, error) {
	rows := a.foodServerDb.Db.QueryRow("SELECT user_id, user_firstname, user_lastname, user_email, user_phone, user_password FROM users WHERE user_email = $1", login.Email);

	loginDetails := entities.User{}
	err := rows.Scan(&loginDetails.Id, &loginDetails.FirstName, &loginDetails.LastName, &loginDetails.Email, &loginDetails.Phone, &loginDetails.Password) 
	if err != nil{
		fmt.Println(err)
		return entities.User{}, errors.New("incorrect email...")
	} 

	if err = bcrypt.CompareHashAndPassword([]byte(loginDetails.Password), []byte(password)); 
	err != nil {
		return entities.User{}, errors.New("Incorrect password.")
	}
	
	return loginDetails, nil
	
}