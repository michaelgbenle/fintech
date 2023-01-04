package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
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


type Blacklist struct {
	Email     string
	Token     string
	CreatedAt string
}

type Transaction struct {}
type Money struct {}