package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/fintech/internal/helpers"
	"github.com/michaelgbenle/fintech/internal/models"
)

func (u *HTTPHandler) SignUpHandler(c *gin.Context) {
	var user *models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		helpers.Response(c, "error", 400, nil, []string{"invalid request"})
		return
	}
//check for valid email
	if !user.ValidateEmail() {
		helpers.Response(c, "error", 400, nil, []string{"invalid email"})
		return
	}

	//hash password
	err = user.HashPassword()
	if err != nil {
		helpers.Response(c, "error", 500, nil, []string{"internal server error"})
		return
	}


	helpers.Response(c, "account created successfully", 201, nil, nil)
}
