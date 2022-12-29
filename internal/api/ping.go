package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/fintech/internal/helpers"
)

//PingHandler is for testing the connections
func (u *HTTPHandler) PingHandler(c *gin.Context) {
	data := "i'm ready"

	// healthcheck
	helpers.Response(c, "pong", 200, data, nil)
}
