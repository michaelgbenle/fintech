package models

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required" gorm:"unique"`
	Password  string `json:"password" binding:"required"`
	Pin       string `json:"pin" binding:"required"`
	Wallet    string `json:"wallet"`
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}
func (user *User) HashPin() error {
	hashedPin, err := bcrypt.GenerateFromPassword([]byte(user.Pin), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Pin = string(hashedPin)
	return nil
}

func (user *User) ValidateEmail() bool {
	emailRegexp := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{5,50}$`)
	return emailRegexp.MatchString(user.Email)

}