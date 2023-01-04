package api

import (
	"fmt"
	"net/http"

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

//check if email exists
user, userErr := u.Repository.FindUserByEmail(loginRequest.Email)
if userErr != nil {
	fmt.Println(userErr)
	helpers.Response(c, "bad request", http.StatusBadRequest, nil, []string{"email does not exists"})
	return
}

//check if password is correct




	helpers.Response(c, "login successful", 200, nil, nil)
}
