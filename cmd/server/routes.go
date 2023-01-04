package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/fintech/internal/api"
	"github.com/michaelgbenle/fintech/internal/middleware"
	"github.com/michaelgbenle/fintech/internal/ports"
)

//SetupRouter is where router endpoints are called
func SetupRouter(handler *api.HTTPHandler, repository ports.Repository) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r := router.Group("/api/v1")
	{
		r.GET("/ping", handler.PingHandler)
		r.POST("/register", handler.SignUpHandler)
		r.POST("/login", handler.LoginHandler)
	}

	// authorizeUser authorizes all authorized users handlers
	authorizeUser := r.Group("/user")
	authorizeUser.Use(middleware.AuthorizeUser(repository.FindUserByEmail, repository.TokenInBlacklist))
	{
		authorizeUser.PATCH("credit", handler.CreditHandler)
		authorizeUser.PATCH("debit", handler.DebitHandler)
		authorizeUser.GET("transactions", handler.TransactionsHandler)

	}

	return router
}
