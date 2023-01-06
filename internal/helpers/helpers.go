package helpers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const min = 1111111111
const max = 9999999999

//Response is customized to help return all responses need
func Response(c *gin.Context, message string, status int, data interface{}, errs []string) {
	responsedata := gin.H{
		"message":   message,
		"data":      data,
		"errors":    errs,
		"status":    http.StatusText(status),
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(status, responsedata)
}

func CreateWallet() int {
	// set seed
	rand.Seed(time.Now().UnixNano())
	// generate random number
	return rand.Intn(max-min) + min
}

func CompareCode(a, b int) bool {
	return a == b
}

func ValidateAccountNumber(accountNos string) bool {
	return len(accountNos) == 10
}
