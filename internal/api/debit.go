package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/fintech/internal/helpers"
)

func (u *HTTPHandler) DebitHandler(c *gin.Context) {

	helpers.Response(c, "account created successfully", 201, nil, nil)
}