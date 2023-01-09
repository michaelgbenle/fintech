package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/fintech/internal/helpers"
	"github.com/michaelgbenle/fintech/internal/models"
)

func (u *HTTPHandler) DebitHandler(c *gin.Context) {
	debiter, err := u.GetUserFromContext(c)
	if err != nil {
		helpers.Response(c, "Unauthorized", http.StatusUnauthorized, nil, []string{"unauthorized"})
		return
	}

	debit := &models.Money{}
	err = c.ShouldBindJSON(&debit)
	if err != nil {
		helpers.Response(c, "error", 400, nil, []string{"invalid request"})
		return
	}

//validate amount
if debit.Amount <= 0 {
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


	helpers.Response(c, "account debited successfully", 201, nil, nil)
}
