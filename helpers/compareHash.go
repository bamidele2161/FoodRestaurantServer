package helpers

import "golang.org/x/crypto/bcrypt"

func CheckHash(hashpass string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(password)) 
	if err != nil {
		return err
	}
	return nil
}