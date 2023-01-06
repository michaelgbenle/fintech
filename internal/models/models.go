package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	Id        uuid.UUID    `gorm:"primary_key; unique; type:uuid; column:id; default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = uuid.New()
	// if m.Id == nil {
	// 	err = errors.New("can't save invalid data")
	// }
	return
}

type Blacklist struct {
	Email     string `json:"email"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
}

type Transaction struct {
	Model
	CustomerId string `json:"customer_id"`
	AccountNos string `json:"account_nos"`
	Type       string `json:"type"`
	Success    bool   `json:"success"`
}
type Money struct {
	AccountNos string  `json:"account_nos" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
}
