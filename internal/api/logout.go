package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/fintech/internal/helpers"
)

func (u *HTTPHandler) LogoutHandler(c *gin.Context) {
	tokenstr, err := u.GetTokenFromContext(c)
	if err != nil {
		helpers.Response(c, "error getting access token", http.StatusInternalServerError, nil, []string{"internal error"})
		return
	}

	foodBeneficiary, err := u.GetUserFromContext(c)
	if err != nil {
		helpers.JSON(c, "error getting access token", http.StatusBadRequest, nil, []string{"bad request"})
		return
	}


	helpers.Response(c, "transactions", 201, nil, nil)

}