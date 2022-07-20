package services

import (
	"FoodServer/db"
	"FoodServer/entities"
	"database/sql"
	"errors"
	"fmt"
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
	rows, _ := a.foodServerDb.Db.Exec("SELECT * from customers where user_email = $1 or user_phone = $2", user.Email, user.Phone);
	
	// rows affected 
	nRows, err := rows.RowsAffected()
	if nRows == 0{
		var userId int
		 err := a.foodServerDb.Db.QueryRow(`Insert into customers (user_firstname, user_lastname, user_email, user_password, user_phone) values ($1, $2, $3, $4, $5) returning user_id`, user.FirstName, user.LastName, user.Email, user.Password, user.Phone).Scan(&userId)
		if err != nil {
			fmt.Println(err)
			return entities.User{}, errors.New("An error occurred while creating user")
		}

		if err != nil {

			return entities.User{}, errors.New("An error occurred while creating user")
		}
		createdUser := entities.User{}

		row := a.foodServerDb.Db.QueryRow("SELECT * FROM customers WHERE user_id = $1", userId) 
		row.Scan(&createdUser.Id, &createdUser.FirstName, &createdUser.LastName, &createdUser.Email, &createdUser.Phone, &createdUser.Password)

		return createdUser, nil
	}
	fmt.Println(err)
	return entities.User{}, errors.New("User already exists")
}

func(a Auth) LoginUser(login entities.Login) (entities.User, error) {
	rows := a.foodServerDb.Db.QueryRow("SELECT * FROM customers WHERE user_email = $1 or user_phone = $2", login.Email, login.Password);

	if rows.Scan() == sql.ErrNoRows{
			return entities.User{}, errors.New("incorrect email...")
	} else {
		loginDetails := entities.User{}
		rows.Scan(&loginDetails.Id, &loginDetails.FirstName, &loginDetails.LastName, &loginDetails.Phone, &loginDetails.Password)
		return loginDetails, nil
	}
}