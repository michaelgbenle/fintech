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

	helpers.Response(c, "account debited successfully", 201, nil, nil)
}
