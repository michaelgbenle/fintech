package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/fintech/internal/helpers"
)

func (u *HTTPHandler) DebitHandler(c *gin.Context) {
	debiter, err := u.GetUserFromContext(c)
	if err != nil {
		helpers.Response(c, "Unauthorized", http.StatusUnauthorized, nil, []string{"unauthorized"})
		return
	}

	debit := &models.Money{}
	err = c.ShouldBindJSON(&credit)
	if err != nil {
		helpers.Response(c, "error", 400, nil, []string{"invalid request"})
		return
	}

	helpers.Response(c, "account debited successfully", 201, nil, nil)
}
