package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/fintech/internal/helpers"
	"github.com/michaelgbenle/fintech/internal/models"
)

func (u *HTTPHandler) CreditHandler(c *gin.Context) {
	_, err := u.GetUserFromContext(c)
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
	//validate account number
	if !helpers.ValidateAccountNumber(credit.AccountNos) {
		helpers.Response(c, "error", 400, nil, []string{"invalid account number"})
		return
	}
	//validate amount
	if credit.Amount <= 0 {
		helpers.Response(c, "error", 400, nil, []string{"invalid amount"})
		return
	}

	//credit user
	transaction, CreditErr := h.DB.Creditwallet(credit)
	if CreditErr != nil {
		helpers.Response(c, "unable to credit user", 500, nil, []string{"credit error"})
		return
	}

	helpers.Response(c, "account created successfully", 201, nil, nil)
}
