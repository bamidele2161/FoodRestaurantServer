package entities

type User struct {
	Id int `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Password string `json:"password"`
}

type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
