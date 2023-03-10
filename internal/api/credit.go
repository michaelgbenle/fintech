package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/fintech/internal/helpers"
	"github.com/michaelgbenle/fintech/internal/models"
)

func (u *HTTPHandler) CreditHandler(c *gin.Context) {
	creditor, err := u.GetUserFromContext(c)
	if err != nil {
		helpers.Response(c, "Unauthorized", http.StatusUnauthorized, nil, []string{"unauthorized"})
		return
	}

	credit := &models.Money{}
	err = c.ShouldBindJSON(&credit)
	if err != nil {
		helpers.Response(c, "error", 400, nil, []string{"invalid request"})
		return
	}
	//validate amount
	if credit.Amount <= 0 {
		helpers.Response(c, "error", 400, nil, []string{"invalid amount"})
		return
	}

	//validate account number
	if !helpers.ValidateAccountNumber(credit.AccountNos) {
		helpers.Response(c, "error", 400, nil, []string{"invalid account number"})
		return
	}

	//check if account number exists
	_, err = u.Repository.FindUserByAccountNos(credit.AccountNos)
	if err != nil {
		helpers.Response(c, "error", 400, nil, []string{"account number does not exist"})
		return
	}

	//check for insufficient balance
	if creditor.Balance < credit.Amount {
		helpers.Response(c, "insufficient balance", 400, nil, []string{"insufficient balance"})
		return
	}

	//credit user
	transaction, CreditErr := u.Repository.Creditwallet(credit, creditor)
	if CreditErr != nil {
		helpers.Response(c, "unable to credit user", 500, nil, []string{"credit error"})
		return
	}

	helpers.Response(c, "account credited successfully", 200, transaction, nil)
}
