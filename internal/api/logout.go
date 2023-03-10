package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/michaelgbenle/fintech/internal/helpers"
	"github.com/michaelgbenle/fintech/internal/middleware"
)

func (u *HTTPHandler) LogoutHandler(c *gin.Context) {
	tokenstr, err := u.GetTokenFromContext(c)
	if err != nil {
		helpers.Response(c, "error getting access token", http.StatusInternalServerError, nil, []string{"internal error"})
		return
	}

	user, err := u.GetUserFromContext(c)
	if err != nil {
		helpers.Response(c, "error getting access token", http.StatusInternalServerError, nil, []string{"internal error"})
		return
	}
	token, _ := jwt.ParseWithClaims(tokenstr, &middleware.Claims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("invalid signing algorithm")
		}
		return os.Getenv("JWT_SECRET"), nil
	})

	if claims, ok := token.Claims.(*middleware.Claims); !ok && !token.Valid {
		helpers.Response(c, "error inserting claims", http.StatusInternalServerError, nil, []string{"invalid claims"})
		return
	} else {
		claims.StandardClaims.ExpiresAt = time.Now().Add(-time.Hour).Unix()
	}

	err = u.Repository.AddTokenToBlacklist(user.Email, tokenstr)
	if err != nil {
		helpers.Response(c, "error inserting token into database", http.StatusInternalServerError, nil, []string{"Claims not valid type"})
		return
	}
	
	helpers.Response(c, "log out successful", 201, nil, nil)
}