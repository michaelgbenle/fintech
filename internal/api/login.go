package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/fintech/internal/helpers"
	"github.com/michaelgbenle/fintech/internal/models"
)

func (u *HTTPHandler) LoginHandler(c *gin.Context) {
var loginRequest *models.LoginRequest
err := c.ShouldBindJSON(&loginRequest)
if err != nil {
	helpers.Response(c, "error", 400, nil, []string{"invalid request"})
	return
}


	helpers.Response(c, "login successful", 200, nil, nil)
}
