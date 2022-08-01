package helpers

import "golang.org/x/crypto/bcrypt"


func HashPassword(password string, minCost int) ([]byte, error) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), minCost) 
	return hashPassword, nil
}
