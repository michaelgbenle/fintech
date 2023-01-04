package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type Model struct {
	Id 	  		string		`sql:"type:uuid; default:uuid_generate_v4();size:100; not null"`
	CreatedAt 	time.Time	`json:"created_at,omitempty"`
	UpdatedAt 	time.Time	`json:"updated_at,omitempty"`
	DeletedAt 	time.Time	`json:"deleted_at,omitempty"`
}
func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = uuid.New().String()
	if m.Id == "" {
		err = errors.New("can't save invalid data")
	}
	return
}

type User struct {
	Model
	FirstName 	string	`json:"first_name" binding:"required"`
	LastName  	string	`json:"lastname" binding:"required"`
	Email     	string	`json:"email" binding:"required" gorm:"unique"`
	Password  	string	`json:"password" binding:"required"`
	Pin 	 	string	`json:"pin" binding:"required"`
	Wallet   	string	`json:"wallet"`

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
	hashedPin, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Pin = string(hashedPin)
	return nil
}

type Blacklist struct {
	Email     string
	Token     string
	CreatedAt string
}
