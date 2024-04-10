package utilities

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(userPassword string, givenPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(givenPassword))
	return err == nil
}
