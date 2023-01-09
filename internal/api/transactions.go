package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/fintech/internal/helpers"
)

func (u *HTTPHandler) TransactionsHandler(c *gin.Context) {
	debiter, err := u.GetUserFromContext(c)
	if err != nil {
		helpers.Response(c, "Unauthorized", http.StatusUnauthorized, nil, []string{"unauthorized"})
		return
	}
	transactions, err := u.Repository.GetTransactions()
	if err != nil {
		helpers.Response(c, "error", 500, nil, []string{"error getting transactions"})
		return
	}


	helpers.Response(c, "transactions", 201, transactions, nil)
}
