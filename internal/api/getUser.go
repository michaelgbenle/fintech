package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/fintech/internal/helpers"
)

func (u *HTTPHandler) GetUserByAccountNumber(c *gin.Context) {
	
	user, err := u.Repository.FindUserByAccountNos(c.Query("accountNumber"))
	if err != nil {
		helpers.Response(c, "error", 500, nil, []string{"error getting user"})
		return
	}

	helpers.Response(c, "transactions", 201, user, nil)
}