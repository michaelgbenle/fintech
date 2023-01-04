package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/fintech/internal/helpers"
	"github.com/michaelgbenle/fintech/internal/models"
)

func (u *HTTPHandler) LoginHandler(c *gin.Context) {
var loginRequest *models.LoginRequest


	helpers.Response(c, "login successful", 200, nil, nil)
}
