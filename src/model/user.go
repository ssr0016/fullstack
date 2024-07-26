package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           uint
	FirstName    string
	LastName     string
	Email        string
	Password     []byte
	IsAmbassador bool
}

func (user *User) SetPassword(password string) []byte {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	return hashedPassword

}
