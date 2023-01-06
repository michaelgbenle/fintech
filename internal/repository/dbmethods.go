package repository

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
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

func (p *Postgres) FindUserById(Id string) (*models.User, error) {
	user := &models.User{}
	if err := p.DB.Where("id = ?", Id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (p *Postgres) FindUserByAccountNos(account string) (*models.User, error) {
	user := &models.User{}
	if err := p.DB.Where("account_nos = ?", account).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (p *Postgres) CreateUser(user *models.User)  error{
	err := p.DB.Create(&user).Error
	if err != nil {
		log.Println("error in creating user")
		return err
	}
return nil
}
func (p *Postgres) Creditwallet(money *models.Money, creditor *models.User)  (*models.Transaction,error){
	accountNos, amount := money.AccountNos, money.Amount
	user, _ := p.FindUserByAccountNos(accountNos)
	transaction := models.Transaction{
		CustomerId:user.Id ,
		AccountNos: money.AccountNos,
		Type: "credit",
		Success: true,
	}
	p.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&creditor).Update("balance", creditor.Balance + amount).Error; err != nil {
			return err
		}
		if err := tx.Model(&user).Update("balance", user.Balance - amount).Error; err != nil {
			return err
		}
		return nil
	})
	err := p.DB.Create(&transaction).Error
	if err != nil {
		log.Println("error in creating user")
		return nil,err
	}
	return &transaction,nil
}