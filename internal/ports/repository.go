package ports

import "github.com/michaelgbenle/fintech/internal/models"

type Repository interface {
	FindUserByEmail(email string) (*models.User, error)
	TokenInBlacklist(token *string) bool
	FindUserById(Id string) (*models.User, error)
	FindUserByAccountNos(account string) (*models.User, error)
	CreateUser(user *models.User) error
	Creditwallet(money *models.Money, creditor *models.User) (*models.Transaction, error)
}
