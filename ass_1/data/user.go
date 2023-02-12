package data

import "golang.org/x/crypto/bcrypt"

type User struct {
	Email      string
	Password   string
	GiveRating bool
}

type AuthUser struct {
	Email        string
	PasswordHash string
	GiveRating   bool
}

func getPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(hash), err
}
