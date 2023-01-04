package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Model struct {
	Id 	  		string		`json:"id"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
	DeletedAt 	time.Time	`json:"deleted_at"`
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
	FirstName string	`json:"firstname"`
	LastName  string	`json:"lastname"`
	Email     string	`json:"email"`
	Password  string	`json:"password"`
	Pin 	 string		`json:"pin"`
	Wallet    string	`json:"wallet"`

}

type Blacklist struct {
	Email     string
	Token     string
	CreatedAt string
}
