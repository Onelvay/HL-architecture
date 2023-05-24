package http

import (
	"errors"
	"fmt"
	"github.com/Onelvay/HL-architecture/pkg/parser/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Middleware(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var key []byte
	username, err := token.Parse(headerParts[1], key)
	fmt.Println(username)
	c.Set("x-testing-abay", username)
	if err == errors.New("invalid auth token") {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	} else if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	return
}
