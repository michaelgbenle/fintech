package repository

import (
	"log"
	"time"

	"github.com/michaelgbenle/fintech/internal/models"
)

// TokenInBlacklist checks if token is already in the blacklist collection
func (p *Postgres) TokenInBlacklist(token *string) bool {
	tok := &models.Blacklist{}
	if err := p.DB.Where("token = ?", token).First(&tok).Error; err != nil {
		return false
	}

	return true
}

// AddTokenToBlacklist adds used token to blacklist
func (p *Postgres) AddTokenToBlacklist(email string, token string) error {
	blacklisted := models.Blacklist{}
	blacklisted.Token = token
	blacklisted.Email = email
	blacklisted.CreatedAt = time.Now().String()

	err := p.DB.Create(&blacklisted).Error
	if err != nil {
		log.Println("error in ad token to blacklist")
		return err
	}
	log.Println("token added to blacklist")
	return nil

}

func (p *Postgres) FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	if err := p.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
