package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/michaelgbenle/fintech/internal/helpers"
	"github.com/michaelgbenle/fintech/internal/middleware"
	"github.com/michaelgbenle/fintech/internal/models"
	"golang.org/x/crypto/bcrypt"
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
if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
	helpers.Response(c, "invalid Password", http.StatusBadRequest, nil, []string{"Bad Request"})
	return
}

	// Generates access claims and refresh claims
	accessClaims, refreshClaims := middleware.GenerateClaims(user.Email)

	secret := os.Getenv("JWT_SECRET")
	accToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)
	if err != nil {
		log.Printf("token generation error err: %v\n", err)
		helpers.Response(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}

	refreshToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, refreshClaims, &secret)
	if err != nil {
		log.Printf("token generation error err: %v\n", err)
		helpers.Response(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}
	c.Header("refresh_token", *refreshToken)
	c.Header("access_token", *accToken)

	helpers.Response(c, "login successful", http.StatusOK, gin.H{
		"user":          student,
		"access_token":  *accToken,
		"refresh_token": *refreshToken,
	}, nil)
}
