package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/fintech/internal/helpers"
)

func (u *HTTPHandler) LoginHandler(c *gin.Context) {


	helpers.Response(c, "login successful", 200, nil, nil)
}
